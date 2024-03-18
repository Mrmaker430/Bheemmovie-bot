package services

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	Movies = make(map[string]string)
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
			name := doc.FileName
			id := doc.FileID
			Movies[name] = id
			msg := tgbotapi.NewMessage(chanId, fmt.Sprintf("%s saved !", name))
			msg.ParseMode = "markdown"
			bot.Send(msg)
		}
	}
}
