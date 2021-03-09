package main

import (
	"cloud.google.com/go/storage"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	bot    *tb.Bot
	bucket *storage.BucketHandle
)

func main() {
	bot, err := InitBot()
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err = InitBucket()
	if err != nil {
		log.Fatalln(err)
	}

	bot.Handle("/start", StartHandler)
	bot.Handle("/help", HelpHandler)
	bot.Handle(tb.OnText, OnTextHandler)
	bot.Handle(tb.OnAnimation, OnAnimationHandler)

	bot.Start()
}
