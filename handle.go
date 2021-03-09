package main

import (
	"bytes"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"io"
	"io/ioutil"
	"os"
)

func StartHandler(m *tb.Message) {
	_ = bot.Notify(m.Chat, tb.Typing)
	_, _ = bot.Send(m.Chat, "不準開始。")
}

func HelpHandler(m *tb.Message) {
	_ = bot.Notify(m.Chat, tb.Typing)
	_, _ = bot.Send(m.Chat, "禁止幫助⛔。")
}

func OnTextHandler(m *tb.Message) {
	_ = bot.Notify(m.Chat, tb.Typing)
	_, _ = bot.Send(m.Chat, "這位先生，本小姐不陪聊哦。")
}

func OnAnimationHandler(m *tb.Message) {
	_ = bot.Notify(m.Chat, tb.UploadingPhoto)
	fmt.Println("111111111111111111111111")
	out, _ := os.Create(m.Animation.FileName + ".mp4")
	fmt.Println("33333333333333333333333")
	_, _ = io.Copy(out, m.Animation.FileReader)
	fmt.Println("44444444444444444444444")
	newFile, _ := ioutil.ReadAll(out)
	_, _ = bot.Reply(m, &tb.File{
		FileReader: bytes.NewReader(newFile),
	})
}
