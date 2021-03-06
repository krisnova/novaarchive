package bot

import (
	"encoding/binary"
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/ChimeraCoder/anaconda"
)

// DefaultBufferSizeBytes is the default buffer size to use
// 1024 * 1024 * 1024 * 2 = 2Gb (<-- This is huge)
const DefaultBufferSizeBytes int64 = 1024 * 1024 * 1024 * 2

// TwitterBot is used to listen for strings on Twitter
// and respond by calling an associated SendTweet function
type TwitterBot struct {
	api           *anaconda.TwitterApi
	searchKeys    []string
	sendFunc      SendTweet
	fSet          bool
	errs          chan error
	errBuffer     []error
	maxBufferSize int64
	tweetMutex    sync.Mutex
}

// TwitterBotCredentials is just an unexported wrapper for
// Twitter auth material
type TwitterBotCredentials struct {
	accessToken    string
	accessSecret   string
	consumerKey    string
	consumerSecret string
}

// NewTwitterBotCredentialsFromEnvironmentalVariables will read from the
// associated environmental variables regardless if they are set.
func NewTwitterBotCredentialsFromEnvironmentalVariables() *TwitterBotCredentials {
	return NewTwitterBotCredentials(os.Getenv("TWITTERTOKEN"),
		os.Getenv("TWITTERTOKENSECRET"),
		os.Getenv("TWITTERCONSUMERKEY"),
		os.Getenv("TWITTERCONSUMERKEYSECRET"))
}

// NewTwitterBotCredentials returns a package safe auth struct
func NewTwitterBotCredentials(accessToken, accessSecret, consumerKey, consumerSecret string) *TwitterBotCredentials {
	return &TwitterBotCredentials{
		accessToken:    accessToken,
		accessSecret:   accessSecret,
		consumerKey:    consumerKey,
		consumerSecret: consumerSecret,
	}
}

// NewTwitterBot is used to build a new TwitterBot
func NewTwitterBot(c *TwitterBotCredentials) *TwitterBot {
	api := anaconda.NewTwitterApiWithCredentials(c.accessToken, c.accessSecret, c.consumerKey, c.consumerSecret)
	return &TwitterBot{
		maxBufferSize: DefaultBufferSizeBytes,
		api:           api,
	}
}

// SetBufferSizeGBytes will set the buffer size in B
func (t *TwitterBot) SetBufferSizeBytes(i int64) {
	t.maxBufferSize = i
}

// SetBufferSizeGBytes will set the buffer size in Gb
func (t *TwitterBot) SetBufferSizeGBytes(i int64) {
	t.maxBufferSize = i * 1024 * 1024 * 1024
}

// AddSlashKey is used to take a string "meeps" and search Twitter for "/meeps"
// add a command string without the slash.
func (t *TwitterBot) AddSlashKey(key string) {
	if strings.HasPrefix(key, "/") {
		t.AddKey(key)
		return
	}
	t.AddKey(fmt.Sprintf("/%s", key))
}

// AddCommand is used
func (t *TwitterBot) AddKey(key string) {
	t.searchKeys = append(t.searchKeys, key)
}

// SendTweet is the function type that will be executed
// for each TwitterBot instance
type SendTweet func(api *anaconda.TwitterApi, tweet anaconda.Tweet) error

// SetSendTweet will set a bot's SendTweet function
func (t *TwitterBot) SetSendTweet(s SendTweet) {
	t.fSet = true
	t.sendFunc = s
}

func (t *TwitterBot) Login() (*anaconda.User, error) {
	values := url.Values{}
	user, err := t.api.GetSelf(values)
	if err != nil {
		return nil, fmt.Errorf("unable to login: error message: %v", err)
	}
	return &user, nil
}

// Run will start the bot concurrently, and return an error if the bot cannot start
func (t *TwitterBot) Run() error {
	ch := make(chan error)

	if len(t.searchKeys) < 1 {
		return fmt.Errorf("unable to start bot, empty search keys, please use AddCommand() to add a search key")
	}
	if !t.fSet {
		return fmt.Errorf("unable to start bot, missing SendTweet function, please use SetSendTweet() to set a function")
	}

	// GetStream()
	values := url.Values{}
	// Documentation: https://developer.twitter.com/en/docs/twitter-api/v1/tweets/filter-realtime/guies/basic-stream-parameters
	values.Set("track", strings.Join(t.searchKeys, ","))
	values.Set("stall_warnings", "true")
	stream := t.api.PublicStreamFilter(values)

	// ------- [ Concurrent Listener ] ------
	go func() {
		for tweetInterface := range stream.C {
			switch v := tweetInterface.(type) {
			case anaconda.Tweet:
				t.tweetMutex.Lock()
				e := t.sendFunc(t.api, v)
				if e != nil {
					ch <- e
				}
				t.tweetMutex.Unlock()
			case error:
				ch <- v
			default:
				ch <- fmt.Errorf("unable to parse type (%v) from anaconda {%+v}", v, tweetInterface)
			}
		}
	}()
	t.errs = ch

	// ------ [ Concurrent Buffer Loader ] ------
	go func() {
		for {
			bufferSize := int64(binary.Size(t.errBuffer))
			if bufferSize > t.maxBufferSize {
				// Droping errors here!
				continue
			}
			err := <-t.errs
			t.errBuffer = append(t.errBuffer, err)
		}
	}()
	return nil
}

// NextError is just like a Next() function and will just pop
// the next error of the queue
//
// Note: Not calling this message is dangerous as eventually the
// buffer will fill up, and messages will be dropped
func (t *TwitterBot) NextError() error {
	for {
		if len(t.errBuffer) > 0 {
			break
		}
	}
	var err error
	// Pop
	// err -> { err, err, err } -> {err}
	//                              Pop
	err, t.errBuffer = t.errBuffer[len(t.errBuffer)-1], t.errBuffer[:len(t.errBuffer)-1]
	return err
}
