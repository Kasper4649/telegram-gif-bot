package main

import tb "gopkg.in/tucnak/telebot.v2"

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
	fileURL := m.Animation.FileURL
	_, _ = bot.Reply(m, fileURL)
	_, _ = bot.Send(m.Chat, m.Animation.File)
}
