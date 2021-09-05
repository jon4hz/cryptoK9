package handler

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"log"
	"net/http"

	"image/jpeg"
	"image/png"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jon4hz/cryptoK9/internal/scam"
)

func ImageHandler(b *gotgbot.Bot, ctx *ext.Context) error {

	id := ctx.Message.Photo[len(ctx.Message.Photo)-1].FileId
	f, err := b.GetFile(id)
	if err != nil {
		log.Println(err)
		return err
	}
	res, err := http.Get("https://api.telegram.org/file/bot" + b.Token + "/" + f.FilePath)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()
	m, s, err := image.Decode(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	buf := new(bytes.Buffer)

	switch s {
	case "jpeg":
		err := jpeg.Encode(buf, m, nil)
		if err != nil {
			log.Println(err)
			return err
		}
	case "png":
		err := png.Encode(buf, m)
		if err != nil {
			log.Println(err)
			return err
		}
	default:
		log.Println("Unsupported image format")
		return errors.New("unsupported image format")
	}
	img := buf.Bytes()

	scam, err := scam.IsScam(img)
	if err != nil {
		log.Println(err)
		return err
	}

	if !scam {
		return nil
	}

	return handleScam(b, ctx)
}

func handleScam(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Delete(b)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = b.BanChatMember(ctx.EffectiveMessage.Chat.Id, ctx.EffectiveMessage.From.Id, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = ctx.EffectiveChat.SendMessage(
		b,
		fmt.Sprintf("⚠️ Scam detected! \n\nBanned user: %s", ctx.EffectiveUser.Username),
		nil,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
