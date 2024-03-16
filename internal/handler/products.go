package handler

import (
	"golang-mongodb/internal/database"
	"golang-mongodb/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProducts(c *gin.Context){
	cursor, err := database.Products.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "unable to fetch products"})
		return
	}

	var products []model.Product
	if err = cursor.All(c, &products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func AddProduct(c *gin.Context){
	var body model.CreateProductRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	res, err := database.Products.InsertOne(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add product"})
		return
	}

	product := model.Product{
		ID:       res.InsertedID.(primitive.ObjectID),
		Name:     body.Name,
		Category: body.Category,
		Price:    body.Price,
		Stock:    body.Stock,
	}

	c.JSON(http.StatusCreated, product)
}

func GetProductById(c *gin.Context){

}

func UpdateProductStockById(c *gin.Context){

}

func UpdateProductPriceById(c *gin.Context){

}

func DeleteProductById(c *gin.Context){

}