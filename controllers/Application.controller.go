// Package controllers ...
package controllers

import (
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// IndexMain ..
func IndexMain(c *gin.Context) {
	var services []models.Services
	err := config.DB.Preload("SubServices", func(db *gorm.DB) *gorm.DB {
		return db.Preload("ServiceOptions")
	}).Preload("ServiceOptions").Find(&services).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"services": services,
	})
}

// StoreNotificationsToken ..
func StoreNotificationsToken(c *gin.Context) {
	var data models.NotificationsToken
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var notificationToken models.NotificationsToken
	err := config.DB.Where("user_id = ?", data.UserID).First(&notificationToken).Error

	if err != nil {
		config.DB.Create(&data)
		c.JSON(200, gin.H{
			"message": "Create New One",
			"code":    100,
		})
		return
	}

	var notificationTokenForCheckTokenItsTheSame models.NotificationsToken
	notSameTokenError := config.DB.Where("user_id = ?", data.UserID).Where("token = ?", data.Token).First(&notificationTokenForCheckTokenItsTheSame).Error
	if notSameTokenError != nil {
		notificationToken.Token = data.Token
		config.DB.Save(&notificationToken)
		c.JSON(200, gin.H{
			"message": "Update the token",
			"code":    101,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Up to date",
		"code":    200,
	})

}
