package main

import (
	"fmt"

	"github.com/muhammedarifp/tg-movie-bot/config"
	"github.com/muhammedarifp/tg-movie-bot/db"
	"github.com/muhammedarifp/tg-movie-bot/server"
	"github.com/muhammedarifp/tg-movie-bot/services"
)

func init() {
}

func main() {
	cfg := config.InitConfig()
	db := db.InitDatabase(cfg)
	if db != nil {
		fmt.Println("connected")
	}
	bot, update := server.InitBot(cfg)
	go services.MovieUpload(bot, update)
	services.SearchMovie(bot, update)
}
