// Package controllers ...
package controllers

import (
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
