package handler

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jon4hz/cryptoK9/internal/mnemonic"
)

var (
	re = regexp.MustCompile(`\s|[\W]`)
)

func MessageHandler(b *gotgbot.Bot, ctx *ext.Context) error {

	batch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		wg.Wait()
		close(batch)
	}()

	go batchMessageText(wg, ctx.EffectiveMessage.Text, 12, batch)
	go batchMessageText(wg, ctx.EffectiveMessage.Text, 24, batch)

	for v := range batch {
		if mnemonic.IsValid(v) {
			return deleteMessage(b, ctx)
		}
	}
	return nil
}

func batchMessageText(wg sync.WaitGroup, msg string, batchSize int, batch chan<- string) {
	defer wg.Done()
	x := re.Split(msg, -1)

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
		log.Println(err)
		return err
	}
	_, err = ctx.EffectiveChat.SendMessage(b, fmt.Sprintf("⚠️ mnemonic phrase detected! Deleted the message (id=%d)", ctx.Message.MessageId), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
