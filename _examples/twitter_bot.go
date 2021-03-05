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
	robot.AddKey("/lubbi")
	logger.Info("Listening for /lubbi...")
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
	user, err := robot.Login()
	if err != nil {
		logger.Critical(err.Error())
		os.Exit(1)
	}
	logger.Info("Running as user @%s (%s)", user.ScreenName, user.Name)
	err = robot.Run()
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
