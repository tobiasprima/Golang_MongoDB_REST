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

func GetProductById(c *gin.Context){
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id provided"})
		return
	}

	result := database.Products.FindOne(c, primitive.M{"_id": _id})
	product := model.Product{}
	err = result.Decode(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to find product"})
		return
	}

	c.JSON(http.StatusOK, product)
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

func UpdateProductStockById(c *gin.Context){
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id provided"})
		return
	}

	var body struct {
		Stock int `json:"stock" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	_, err = database.Products.UpdateOne(c, bson.M{"_id": _id}, bson.M{"$set": bson.M{"stock": body.Stock}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update product stock"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "product stock updated"})
}

func UpdateProductPriceById(c *gin.Context){
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id provided"})
		return
	}

	var body struct {
		Price float32 `json:"price" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	_, err = database.Products.UpdateOne(c, bson.M{"_id": _id}, bson.M{"$set": bson.M{"stock": body.Price}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update product price"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "product price updated"})
}

func DeleteProductById(c *gin.Context){

}