package server

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/muhammedarifp/tg-movie-bot/config"
)

func InitBot(cfg *config.Config) (bot *tgbotapi.BotAPI, update *tgbotapi.UpdatesChannel) {
	if cfg.BotToken == "" {
		panic("Your bot token is empty, please recheck that .!")
	}
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Fatalf("connection error : %s", err.Error())
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	update = &updates

	return
}
