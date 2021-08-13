// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// IndexNewSupplierJoinRequest ..
func IndexNewSupplierJoinRequest(c *gin.Context) {
	var users []models.User

	config.DB.Where("status = ?", 1).Where("user_type = ?", 2).Find(&users)

	c.JSON(200, users)
}

// ApproveSupplier ..
func ApproveSupplier(c *gin.Context) {
	ID := c.Param("id")

	var user models.User
	err := config.DB.Where("id = ?", ID).First(&user).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	user.Status = 2

	config.DB.Save(&user)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

// BlockSupplier
func BlockSupplier(c *gin.Context) {
	ID := c.Param("id")

	var user models.User
	err := config.DB.Where("id = ?", ID).First(&user).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	user.Status = 3

	config.DB.Save(&user)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

// IndexActiveSupplier ..
func IndexActiveSupplier(c *gin.Context) {
	var users []models.User

	config.DB.Where("status = ?", 2).Where("user_type = ?", 2).Find(&users)

	c.JSON(200, users)
}

// IndexBlockListSupplier ..
func IndexBlockListSupplier(c *gin.Context) {
	var users []models.User

	config.DB.Where("status = ?", 3).Where("user_type = ?", 2).Find(&users)

	c.JSON(200, users)
}
