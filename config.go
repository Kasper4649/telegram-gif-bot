package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"runtime"
)

var (
	AccessToken string
	WebhookHost string
	Port        string
)

func init() {
	AccessToken = os.Getenv("ACCESS_TOKEN")
	WebhookHost = os.Getenv("WEBHOOK_HOST")
	Port = os.Getenv("PORT")
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	Formatter := &log.TextFormatter{
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return fmt.Sprintf("[%s()]", f.Function), ""
		},
	}
	log.SetFormatter(Formatter)
}

func InitBot() *tb.Bot {
	webhook := &tb.Webhook{
		Listen: ":" + Port,
		Endpoint: &tb.WebhookEndpoint{
			PublicURL: WebhookHost,
		},
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  AccessToken,
		Poller: webhook,
	})
	if err != nil {
		log.Fatal(err)
	}

	return b
}
