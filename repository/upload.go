package repository

import (
	"context"
	"errors"

	"github.com/muhammedarifp/tg-movie-bot/consts"
	"github.com/muhammedarifp/tg-movie-bot/db"
	"github.com/muhammedarifp/tg-movie-bot/models"
)

func UploadMovie(data *models.MovieUpload) error {
	db := db.GetDB()

	if data.FileID == "" || data.Name == "" {
		return errors.New("invalid input provided")
	}

	if _, err := db.Database(consts.DB_NAME).Collection(consts.COLLECTION_NAME).InsertOne(context.Background(), data); err != nil {
		return err
	}

	return nil
}
