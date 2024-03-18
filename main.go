package main

import (
	"github.com/muhammedarifp/tg-movie-bot/config"
	"github.com/muhammedarifp/tg-movie-bot/server"
	"github.com/muhammedarifp/tg-movie-bot/services"
)

func init() {}

func main() {
	cfg := config.InitConfig()
	bot, update := server.InitBot(cfg)
	go services.MovieUpload(bot, update)
	services.SearchMovie(bot, update)
}
