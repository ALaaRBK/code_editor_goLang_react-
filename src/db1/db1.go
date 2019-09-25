package db1

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Init(opt map[string]interface{}) error {
	url := opt["MONGO_URL"]
	dbName := opt["DB_NAME"]

	clientOptions := options.Client().ApplyURI(url.(string))
	client, err := mongo.NewClient(clientOptions)
	err = client.Connect(context.Background())

	if err != nil {
		fmt.Println("can't connect to mongo server", err)
		return err
	}

	// err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		fmt.Println("can't ping to mongo server", err)
		return err
	}

	DB = client.Database(dbName.(string))

	return nil
}
