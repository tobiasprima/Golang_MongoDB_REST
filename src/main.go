package main

import (
	"fmt"
	"golang-mongodb/internal/database"
	"golang-mongodb/internal/handler"
	"os"

	"github.com/gin-gonic/gin"
)

func main(){

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

	r.GET("/products", handler.GetProducts)
	r.GET("/products/:id", handler.GetProductById)

	r.POST("/products", handler.AddProduct)
	r.PATCH("/products/:id/stock", handler.UpdateProductStockById)
	r.PATCH("/products/:id/price", handler.UpdateProductPriceById)
	r.DELETE("/products/:id", handler.DeleteProductById)

	r.Run(":8080")
}

