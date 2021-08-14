package handler

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jon4hz/mnemonicK9/internal/mnemonic"
)

func MessageHandler(b *gotgbot.Bot, ctx *ext.Context) error {

	batch := make(chan string)

	go batchMessageText(ctx.EffectiveMessage.Text, batch)

	for v := range batch {
		fmt.Println(v)
		if mnemonic.IsValid(v) {
			return deleteMessage(b, ctx)
		}
	}
	return nil
}

func batchMessageText(msg string, batch chan<- string) {
	x := strings.Fields(msg)

	for i := 0; i < len(x); i++ {
		var phrase = make([]string, 12)
		for j := 0; j < 12; j++ {
			if i+j < len(x) {
				phrase[j] = x[j+i]
			} else {
				close(batch)
				return
			}
		}
		batch <- strings.Join(phrase, " ")
	}
}

func deleteMessage(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Delete(b)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = ctx.EffectiveChat.SendMessage(b, fmt.Sprintf("⚠️ mnemonic phrase detected! Deleted the message (id=%d)", ctx.Message.MessageId), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
