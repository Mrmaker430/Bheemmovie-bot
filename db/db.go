package db

import (
	"context"
	"log"
	"time"

	"github.com/muhammedarifp/tg-movie-bot/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Client
)

func InitDatabase(cfg *config.Config) (client *mongo.Client) {
	var err error
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(cfg.DbClient))
	if err != nil {
		log.Println("Database Connection Error")
		return nil
	}

	// keys := mongo.IndexModel{
	// 	Keys:    bson.D{{"name", 1}},
	// 	Options: options.Index().SetUnique(true),
	// }

	// if _, err := client.Database(consts.DB_NAME).Collection(consts.COLLECTION_NAME).Indexes().CreateOne(context.TODO(), keys); err != nil {
	// 	panic(err.Error())
	// }

	db = client
	return
}

func GetDB() *mongo.Client {
	return db
}
