package main

import (
	"cloud.google.com/go/storage"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	bot    *tb.Bot
	bucket *storage.BucketHandle
)

func main() {
	bot = InitBot()
	//bucket, _ = InitBucket()

	bot.Handle("/start", StartHandler)
	bot.Handle("/help", HelpHandler)
	bot.Handle(tb.OnText, OnTextHandler)
	bot.Handle(tb.OnAnimation, OnAnimationHandler)

	bot.Start()
}
