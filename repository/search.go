package repository

import (
	"context"

	"github.com/muhammedarifp/tg-movie-bot/consts"
	"github.com/muhammedarifp/tg-movie-bot/db"
	"github.com/muhammedarifp/tg-movie-bot/models"
	"go.mongodb.org/mongo-driver/bson"
)

func SearchMovie(inp string) ([]models.MovieUpload, int) {
	db := db.GetDB()
	collection := db.Database(consts.DB_NAME).Collection(consts.COLLECTION_NAME)
	filter := bson.M{
		"$text": bson.M{
			"$search": inp,
		},
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return []models.MovieUpload{}, 0
	}
	defer cursor.Close(context.Background())
	var files []models.MovieUpload
	cursor.All(context.Background(), &files)
	if len(files) < 1 {
		return files, 0
	}

	return files, len(files)
}
