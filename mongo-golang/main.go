package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main(){
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: %v", err)
	}

	databaseURI := os.Getenv("DATABASE_URI")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(databaseURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Database started successfully!")
}
