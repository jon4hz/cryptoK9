package handler

import (
	"fmt"
	"strings"
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jon4hz/mnemonicK9/internal/mnemonic"
)

func MessageHandler(b *gotgbot.Bot, ctx *ext.Context) error {

	batch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		wg.Wait()
		close(batch)
	}()

	go batch12(wg, ctx.EffectiveMessage.Text, batch)
	go batch24(wg, ctx.EffectiveMessage.Text, batch)

	for v := range batch {
		if mnemonic.IsValid(v) {
			return deleteMessage(b, ctx)
		}
	}
	return nil
}

func batch12(wg sync.WaitGroup, msg string, batch chan<- string) {
	batchMessageText(wg, msg, 12, batch)
}

func batch24(wg sync.WaitGroup, msg string, batch chan<- string) {
	batchMessageText(wg, msg, 24, batch)
}

func batchMessageText(wg sync.WaitGroup, msg string, batchSize int, batch chan<- string) {
	defer wg.Done()
	x := strings.Fields(msg)

	for i := 0; i < len(x); i++ {
		var phrase = make([]string, batchSize)
		for j := 0; j < batchSize; j++ {
			if i+j < len(x) {
				phrase[j] = x[j+i]
			} else {
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
