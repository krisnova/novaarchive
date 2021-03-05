# Bot

Libraries for bots on the interwebs

# Auth

Copypasta for auth config

```bash
export TWITTERTOKEN=""
export TWITTERTOKENSECRET=""
export TWITTERCONSUMERKEY=""
export TWITTERCONSUMERKEYSECRET=""
```

# Example Bot 

```go 
package main

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/kris-nova/logger"
	"github.com/kris-nova/novaarchive/bot"
)

func main() {
	logger.BitwiseLevel = logger.LogEverything
	logger.Info("Starting bot...")
	robot := bot.NewTwitterBot(bot.NewTwitterBotCredentialsFromEnvironmentalVariables())
	robot.AddCommand("/lubbi")
	logger.Info("Setting command /lubbi...")
	robot.SetBufferSizeGBytes(1)
	logger.Info("Setting buffer 1Gb...")
	robot.SetSendTweet(func(api *anaconda.TwitterApi, tweet anaconda.Tweet) error {
		logger.Always("Found tweet: %s", tweet.IdStr)
		//
		// Your logic here
		//
		return nil
	})
	logger.Info("Setting SendTweet...")
	err := robot.Run()
	if err != nil {
		logger.Critical(err.Error())
		os.Exit(1)
	}
	logger.Info("Running bot...")
	for {
		err := robot.NextError()
		logger.Critical(err.Error())
	}
}

```


