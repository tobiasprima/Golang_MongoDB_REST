package main

import (
	"fmt"
	"golang-mongodb/internal/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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
	
	r := gin.Default()

	r.GET("/products")
	r.GET("/products/:id")

	r.POST("/products")
	r.PATCH("/products/:id/stock")
	r.PATCH("/products/:id/price")
	r.DELETE("/products/:id")

	r.Run(":8080")
}
