package main

import (
	"fmt"

	"github.com/jon4hz/cryptoK9/internal/config"
	"github.com/jon4hz/cryptoK9/internal/telegram"
)

func main() {
	fmt.Println("started...")

	cfg := config.Get()

	bot := telegram.NewBot(cfg.Telegram.Token)
	bot.Start()
}
