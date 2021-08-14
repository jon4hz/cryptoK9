package telegram

import (
	"log"
	"net/http"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"github.com/jon4hz/mnemonicK5/internal/telegram/handler"
)

var (
	TGBot Bot
)

type Config struct {
	Token string `yaml:"token"`
}

type Bot struct {
	Bot *gotgbot.Bot
}

func NewBot(token string) *Bot {
	var err error
	TGBot.Bot, err = gotgbot.NewBot(token, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		log.Fatal("failed to create new bot: " + err.Error())
	}

	return &TGBot
}

func (b *Bot) Start() {
	// pass value
	updater := ext.NewUpdater(nil)
	dispatcher := updater.Dispatcher

	setHandlers(dispatcher)

	// TODO replace with webhook
	err := updater.StartPolling(b.Bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: gotgbot.GetUpdatesOpts{
			AllowedUpdates: []string{
				"message",
			},
		},
	})
	if err != nil {
		log.Fatal("failed to start polling: " + err.Error())
	}
	log.Printf("%s has been started...", b.Bot.User.Username)

	updater.Idle()
}

func setHandlers(d *ext.Dispatcher) {

	// msgs
	msgHandler := handlers.NewMessage(message.Text, handler.MessageHandler)
	msgHandler.AllowChannel = true
	d.AddHandler(msgHandler)
}
