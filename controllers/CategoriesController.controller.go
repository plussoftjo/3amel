// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// ------------- Categories -------------//

// StoreCategory ..
func StoreCategory(c *gin.Context) {
	var category models.Categories
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, category)
}

// IndexCategories ..
func IndexCategories(c *gin.Context) {
	var categories []models.Categories
	config.DB.Find(&categories)
	c.JSON(200, categories)
}

// DeleteCategory ..
func DeleteCategory(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.Categories{}, ID)
	var categories []models.Categories
	config.DB.Find(&categories)
	c.JSON(200, categories)
}

// UpdateCategory ..
func UpdateCategory(c *gin.Context) {
	var category models.Categories
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&category).Update(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var categories []models.Categories
	config.DB.Find(&categories)
	c.JSON(200, gin.H{
		"category":   category,
		"categories": categories,
	})
}

// ------------- End Services ------------ //
