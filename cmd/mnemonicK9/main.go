package main

import (
	"github.com/jon4hz/mnemonicK9/internal/config"
	"github.com/jon4hz/mnemonicK9/internal/telegram"
)

func main() {
	cfg := config.Get()

	bot := telegram.NewBot(cfg.Telegram.Token)
	bot.Start()
}
