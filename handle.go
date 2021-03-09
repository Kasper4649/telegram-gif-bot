package main

import tb "gopkg.in/tucnak/telebot.v2"

func StartHandler(m *tb.Message) {
	if !m.Private() {
		return
	}
	_ = b.Notify(m.Chat, tb.Typing)
	_, _ = b.Send(m.Chat, "不準開始。")
}

func HelpHandler(m *tb.Message) {
	_ = b.Notify(m.Chat, tb.Typing)
	_, _ = b.Send(m.Chat, "禁止幫助⛔。")
}

func OnTextHandler(m *tb.Message) {
	if !m.Private() {
		return
	}
	_ = b.Notify(m.Chat, tb.Typing)
	_, _ = b.Send(m.Chat, "這位先生，本小姐不陪聊哦。")
}

func OnAnimationHandler(m *tb.Message) {
	if !m.Private() {
		return
	}
	_ = b.Notify(m.Chat, tb.UploadingPhoto)
	_, _ = b.Reply(m, "這是一個動圖")
}
