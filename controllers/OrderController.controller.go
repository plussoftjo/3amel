// Package controllers ...
package controllers

import (
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
	ID := c.Param("id")

	var order models.Orders
	err := config.DB.Where("id = ?", ID).First(&order).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}

	order.UserRate = true

	config.DB.Save(&order)

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
