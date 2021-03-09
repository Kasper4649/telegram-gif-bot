package main

import tb "gopkg.in/tucnak/telebot.v2"

var b *tb.Bot

func main() {
	b = InitBot()

	b.Handle("/start", StartHandler)
	b.Handle("/help", HelpHandler)
	b.Handle(tb.OnText, OnTextHandler)
	b.Handle(tb.OnAnimation, OnAnimationHandler)

	b.Start()
}
