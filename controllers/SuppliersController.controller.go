// Package controllers ...
package controllers

import (
	"fmt"
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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
	var users []models.User

	config.DB.Preload("SupplierInfo", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Service")
	}).Find(&users)

	c.JSON(200, users)
}

// IndexSupplierInfoWithServiceID ..
func IndexSupplierInfoWithServiceID(c *gin.Context) {
	ID := c.Param("id")

	var supplierInfosIDs []string
	config.DB.Model(&models.SupplierInfo{}).Where("service_id = ?", ID).Pluck("id", &supplierInfosIDs)

	var users []models.User
	err := config.DB.Where("id IN (?)", supplierInfosIDs).Preload("SupplierInfo", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Service")
	}).Find(&users).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
			"codes": 1001,
		})
		return
	}

	c.JSON(200, users)
}

// ShowSupplier ..
func ShowSupplier(c *gin.Context) {
	ID := c.Param("id")
	// Get Supplier With Info
	var supplier models.User
	err := config.DB.Where("id = ?", ID).Preload("SupplierInfo", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Service")
	}).First(&supplier).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
			"code":  500,
		})
		return
	}

	var orders []models.Orders
	config.DB.Where("supplier_id = ?", ID).Scopes(models.OrdersWithDetails).Find(&orders)

	c.JSON(200, gin.H{
		"supplier": supplier,
		"orders":   orders,
	})

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

// IndexAllSupplier ..
func IndexAllSupplier(c *gin.Context) {
	var suppliers []models.User

	config.DB.Where("user_type = ?", 2).Preload("SupplierInfo", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Service")
	}).Find(&suppliers)

	c.JSON(200, suppliers)
}

// StoreSupplierUser ..
func StoreSupplierUser(c *gin.Context) {
	type StoreSupplierUserType struct {
		User      models.User `json:"user"`
		ServiceID uint        `json:"serviceID"`
	}

	var data StoreSupplierUserType
	c.ShouldBindJSON(&data)

	// Store The User
	user := data.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	user.Password = string(hashedPassword)
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplierInfo := models.SupplierInfo{
		UserID:    user.ID,
		Status:    0,
		Latitude:  0,
		Longitude: 0,
		ServiceID: data.ServiceID,
	}

	supplierInfoStoreError := config.DB.Create(&supplierInfo).Error
	if supplierInfoStoreError != nil {
		c.JSON(500, gin.H{
			"err": supplierInfoStoreError.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})

}
