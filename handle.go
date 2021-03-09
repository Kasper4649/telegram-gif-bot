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
	file, err := ioutil.ReadAll(m.Animation.File.FileReader)
	fmt.Println("222222222222222222222222")
	if err != nil {
		_, _ = bot.Send(m.Chat, "文件上傳失敗，請重試;)")
		return
	}
	out, _ := os.Create(m.Animation.FileName + ".mp4")
	fmt.Println("33333333333333333333333")
	_, _ = io.Copy(out, bytes.NewReader(file))
	fmt.Println("44444444444444444444444")
	newFile, _ := ioutil.ReadAll(out)
	_, _ = bot.Reply(m, &tb.File{
		FileReader: bytes.NewReader(newFile),
	})
}
