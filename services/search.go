package services

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SearchMovie(bot *tgbotapi.BotAPI, updates *tgbotapi.UpdatesChannel) {

	for update := range *updates {
		if update.Message == nil { // Ignore any non-message updates
			continue
		}

		// Check if the message is a private message (DM)
		if update.Message.Chat.Type == "private" {
			// Check if the message text matches any movie name in the Movies map
			movieName := strings.TrimSpace(update.Message.Text)
			if fileID, ok := Movies[movieName]; ok {
				// Send the file to the user
				//msg := tgbotapi.NewDocumentShare(update.Message.Chat.ID, fileID)
				fileCfg := tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FileID(fileID))
				_, err := bot.Send(fileCfg)
				if err != nil {
					log.Println("Error sending file:", err)
				}
			} else {
				// If the movie name doesn't exist, notify the user
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Movie not found.")
				_, err := bot.Send(msg)
				if err != nil {
					log.Println("Error sending message:", err)
				}
			}
		}
	}
}
