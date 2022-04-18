package controllers

import (
	"net/http"

	"github.com/carlosm27/go_projects/apiwithgorm/model"
	"github.com/gin-gonic/gin"
)

type NewGrocery struct {
	Name     string `json:"name" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type GroceryUpdate struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func GetGroceries(c *gin.Context) {

	var groceries []model.Grocery

	model.Db.Find(&groceries)
	c.JSON(http.StatusOK, groceries)

}

func GetGrocery(c *gin.Context) {

	var grocery model.Grocery

	if err := model.Db.Where("id= ?", c.Param("id")).First(&grocery).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Grocery not found"})
		return
	}

	c.JSON(http.StatusOK, grocery)

}

func PostGrocery(c *gin.Context) {

	var grocery NewGrocery

	if err := c.ShouldBindJSON(&grocery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newGrocery := model.Grocery{Name: grocery.Name, Quantity: grocery.Quantity}
	model.Db.Create(&newGrocery)

	c.JSON(http.StatusOK, newGrocery)
}

func UpdateGrocery(c *gin.Context) {

	var grocery model.Grocery

	if err := model.Db.Where("id = ?", c.Param("id")).First(&grocery).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Grocery not found!"})
		return
	}

	var updateGrocery GroceryUpdate

	if err := c.ShouldBindJSON(&updateGrocery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model.Db.Model(&grocery).Updates(model.Grocery{Name: updateGrocery.Name, Quantity: updateGrocery.Quantity})
	c.JSON(http.StatusOK, grocery)

}

func DeleteGrocery(c *gin.Context) {

	var grocery model.Grocery

	if err := model.Db.Where("id = ?", c.Param("id")).First(&grocery).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Grocery not found!"})
		return
	}

	model.Db.Delete(&grocery)
	c.JSON(http.StatusOK, gin.H{"message": "Grocery deleted"})

}
