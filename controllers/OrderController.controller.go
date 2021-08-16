// Package controllers ...
package controllers

import (
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// StoreOrder ..
func StoreOrder(c *gin.Context) {
	var order models.Orders

	c.ShouldBindJSON(&order)
	err := config.DB.Create(&order).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
	}

	for _, serviceOption := range order.ServiceOptionsIDs {
		config.DB.Exec("INSERT INTO order_services_option (orders_id,services_options_id) VALUES (?,?)", order.ID, serviceOption)
	}

	var orderData models.Orders
	config.DB.Scopes(models.OrdersWithDetails).Where("id = ?", order.ID).First(&orderData)

	c.JSON(200, orderData)
}

// FinishOrderFromUser ..
func FinishOrderFromUser(c *gin.Context) {
	var userRate models.UserRate
	c.ShouldBindJSON(&userRate)
	var order models.Orders
	err := config.DB.Where("id = ?", userRate.OrderID).First(&order).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}

	order.UserRate = true
	config.DB.Save(&order)

	config.DB.Create(&userRate)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func ShowOrder(c *gin.Context) {
	var order models.Orders
	config.DB.Scopes(models.OrdersWithDetails).Where("id = ?", 3).First(&order)

	c.JSON(200, gin.H{
		"order": order,
	})
}

// IndexNewOrders ..
func IndexNewOrders(c *gin.Context) {

	var orders []models.Orders
	config.DB.Scopes(models.OrdersWithDetails).Order("id desc").Where("status = ?", 0).Find(&orders)

	c.JSON(200, orders)
}

// IndexInWorkOrders ..
func IndexInWorkOrders(c *gin.Context) {

	var orders []models.Orders
	config.DB.Scopes(models.OrdersWithDetails).Order("id desc").Where("status = ?", 1).Find(&orders)

	c.JSON(200, orders)
}

// IndexEndingOrders ..
func IndexEndingOrders(c *gin.Context) {

	var orders []models.Orders
	config.DB.Scopes(models.OrdersWithDetails).Order("id desc").Where("status = ?", 2).Find(&orders)

	c.JSON(200, orders)
}

// ShowOrder ..
func ViewOrder(c *gin.Context) {
	ID := c.Param("id")
	var order models.Orders
	config.DB.Where("id = ?", ID).Preload("User").Scopes(models.OrdersWithDetails).First(&order)

	c.JSON(200, order)
}

// OrderApproveFromController ..
func OrderApproveFromController(c *gin.Context) {
	var data models.Orders
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var order models.Orders
	config.DB.Where("id = ?", data.ID).First(&order)

	order.Status = data.Status
	order.SupplierID = data.SupplierID
	order.Cost = data.Cost

	config.DB.Save(&order)

	c.JSON(200, gin.H{
		"message": "Success",
	})
}
