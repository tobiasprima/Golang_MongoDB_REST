package main

import (
	"fmt"
	"golang-mongodb/internal/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	databaseURI := os.Getenv("DATABASE_URI")

	err := database.Init((databaseURI), "development")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to MongoDB!")

	defer func(){
		err := database.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
}
