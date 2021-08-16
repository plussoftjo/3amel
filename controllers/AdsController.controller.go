// Package controllers ...
package controllers

import (
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// ------------- Services Options -------------//

// StoreAds ..
func StoreAds(c *gin.Context) {
	var data models.Ads
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Where("id = ?", data.ID).Preload("Service").Preload("SubService").Find(&data)
	c.JSON(200, data)
}

// IndexAds ..
func IndexAds(c *gin.Context) {
	var data []models.Ads
	config.DB.Preload("Service").Preload("SubService").Find(&data)
	c.JSON(200, data)
}

// DestroyAds ..
func DestroyAds(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.Ads{}, ID)
	var data []models.Ads
	config.DB.Preload("Service").Preload("SubService").Find(&data)
	c.JSON(200, data)
}

// UpdateAds ..
func UpdateAds(c *gin.Context) {
	var ad models.Ads
	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&ad).Update(&ad).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var data []models.Ads
	config.DB.Preload("Service").Preload("SubService").Find(&data)
	c.JSON(200, gin.H{
		"ad":  ad,
		"ads": data,
	})
}

// ------------- End Sub Services ------------ //

// IndexAdsForMain ...
func IndexMainServiceAds(c *gin.Context) {
	ID := c.Param("id")
	var ads []models.Ads
	config.DB.Where("service_id = ?", ID).Find(&ads)

	c.JSON(200, ads)
}

// IndexSubServiceAds ..
func IndexSubServiceAds(c *gin.Context) {
	ID := c.Param("id")
	var ads []models.Ads
	config.DB.Where("sub_service_id = ?", ID).Find(&ads)

	c.JSON(200, ads)
}

// -------------End Ads-------------//
// ------------- Services Options -------------//

// StoreAppIntro ..
func StoreAppIntro(c *gin.Context) {
	var data models.AppIntro
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Where("id = ?", data.ID).Find(&data)
	c.JSON(200, data)
}

// IndexAppIntro ..
func IndexAppIntro(c *gin.Context) {
	var data []models.AppIntro
	config.DB.Find(&data)
	c.JSON(200, data)
}

// DestroyAppIntro ..
func DestroyAppIntro(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.AppIntro{}, ID)
	var data []models.AppIntro
	config.DB.Find(&data)
	c.JSON(200, data)
}

// UpdateAppIntro ..
func UpdateAppIntro(c *gin.Context) {
	var appintro models.AppIntro
	if err := c.ShouldBindJSON(&appintro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&appintro).Update(&appintro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var data []models.AppIntro
	config.DB.Find(&data)
	c.JSON(200, gin.H{
		"appIntro":  appintro,
		"AppIntros": data,
	})
}

// ------------- End Sub Services ------------ //

// ------------- Services Options -------------//

// StoreNotification ..
func StoreNotification(c *gin.Context) {
	var data models.Notifications
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Where("id = ?", data.ID).Preload("Service").Preload("SubService").Find(&data)
	c.JSON(200, data)
}

// IndexNotification ..
func IndexNotification(c *gin.Context) {
	var data []models.Notifications
	config.DB.Order("id desc").Preload("Service").Preload("SubService").Find(&data)
	c.JSON(200, data)
}

// DestroyNotification ..
func DestroyNotification(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.Notifications{}, ID)
	var data []models.Notifications
	config.DB.Preload("Service").Preload("SubService").Find(&data)
	c.JSON(200, data)
}

// UpdateNotification ..
func UpdateNotification(c *gin.Context) {
	var notification models.Notifications
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Preload("Service").Preload("SubService").Model(&notification).Update(&notification).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var data []models.Notifications
	config.DB.Preload("Service").Preload("SubService").Find(&data)
	c.JSON(200, gin.H{
		"notification":  notification,
		"notifications": data,
	})
}

// ------------- End Sub Services ------------ //
