package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	demo "../controllers/demo"
	edu "../controllers/edu"
)

func MongoConnect() {

	mg_url := "mongodb://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?authSource=" + os.Getenv("DB_AUTH") + "&ssl=" + os.Getenv("DB_SSL")
	fmt.Println(mg_url)

	// Database Config
	clientOptions := options.Client().ApplyURI(mg_url)
	client, err := mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	db := client.Database("vnedu")

	demo.TodoCollection(db)
	edu.StudentCollection(db)

	return
}
