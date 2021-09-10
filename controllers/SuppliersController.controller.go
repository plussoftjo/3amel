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

// IndexSupplierInfo ..
func IndexSupplierInfo(c *gin.Context) {
	var suppliersInfo []models.SupplierInfo

	config.DB.Preload("User").Find(&suppliersInfo)

	c.JSON(200, suppliersInfo)
}

// SetSupplierInfo ...
func SetSupplierInfo(c *gin.Context) {
	var data models.SupplierInfo

	c.ShouldBindJSON(&data)

	var supplierInfoCount int64
	config.DB.Model(&models.SupplierInfo{}).Where("user_id = ?", data.UserID).Count(&supplierInfoCount)

	if supplierInfoCount == 0 {
		config.DB.Create(&data)
	} else {
		var supplierInfo models.SupplierInfo
		if err := config.DB.Where("user_id = ?", data.UserID).First(&supplierInfo).Error; err != nil {
			c.JSON(500, gin.H{
				"code":  500,
				"error": err,
			})
		}

		supplierInfo.Latitude = data.Latitude
		supplierInfo.Longitude = data.Longitude

		config.DB.Model(&models.SupplierInfo{}).Save(&supplierInfo)
	}

	c.JSON(200, gin.H{
		"message": "Done",
	})
}

// IndexOrdersForSupplier ..
func IndexOrdersForSupplier(c *gin.Context) {

	var ID = c.Param("id")

	var newOrders []models.Orders
	var inWorkOrders []models.Orders
	var endingOrders []models.Orders

	config.DB.Where("supplier_id = ?", ID).Where("status = ?", 0).Order("id desc").Find(&newOrders)
	config.DB.Where("supplier_id = ?", ID).Where("status = ?", 1).Order("id desc").Find(&inWorkOrders)
	config.DB.Where("supplier_id = ?", ID).Where("status = ?", 2).Order("id desc").Find(&endingOrders)

	c.JSON(200, gin.H{
		"newOrders":    newOrders,
		"inWorkOrders": inWorkOrders,
		"endingOrders": endingOrders,
	})

}
