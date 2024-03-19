package services

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/muhammedarifp/tg-movie-bot/models"
	"github.com/muhammedarifp/tg-movie-bot/repository"
)

func MovieUpload(bot *tgbotapi.BotAPI, updates *tgbotapi.UpdatesChannel) {
	bot.Debug = true // Optional for debugging

	chanId := int64(-1002066794810)

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	for update := range *updates {
		if update.ChannelPost == nil { // Ignore non-message updates
			continue
		}

		if update.ChannelPost.Chat.ID == chanId && update.ChannelPost.Document != nil {
			doc := update.ChannelPost.Document
			data := models.MovieUpload{
				Name:     doc.FileName,
				FileID:   doc.FileID,
				FileSize: int64(doc.FileSize),
			}
			if err := repository.UploadMovie(&data); err != nil {
				log.Println("-- error -- :: ", err.Error())
			}
		}
	}
}
