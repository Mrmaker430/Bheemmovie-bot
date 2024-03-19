package services

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/muhammedarifp/tg-movie-bot/repository"
)

func SearchMovie(bot *tgbotapi.BotAPI, updates *tgbotapi.UpdatesChannel) {

	for update := range *updates {
		if update.Message == nil { // Ignore any non-message updates
			continue
		}

		inp := update.Message.Text
		inp = strings.TrimSpace(inp)
		if inp == "" {
			return
		}
		inp = strings.ToLower(inp)
		_, count := repository.SearchMovie(inp)

		if count < 1 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, no records found.")
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes")
			bot.Send(msg)
		}
	}
}
