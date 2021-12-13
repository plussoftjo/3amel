// Package controllers ...
package controllers

import (
	"fmt"
	"net/http"
	"server/config"
	"server/models"
	"server/vendors"

	"github.com/gin-gonic/gin"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
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

// IndexCanceldOrders ..
func IndexCanceldOrders(c *gin.Context) {
	var orders []models.Orders
	config.DB.Scopes(models.OrdersWithDetails).Order("id desc").Where("status = ?", 3).Find(&orders)

	c.JSON(200, orders)
}

// ShowOrder ..
func ViewOrder(c *gin.Context) {
	ID := c.Param("id")
	var order models.Orders
	config.DB.Where("id = ?", ID).Preload("User").Preload("Supplier").Scopes(models.OrdersWithDetails).First(&order)

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

	// Send Notification to the user
	userID := order.UserID

	var notificationToken models.NotificationsToken
	config.DB.Where("user_id = ?", userID).First(&notificationToken)

	var exponentTokens []expo.ExponentPushToken
	pushToken, err := expo.NewExponentPushToken(notificationToken.Token)
	if err != nil {
		fmt.Println("Not Expo token")
	}
	exponentTokens = append(exponentTokens, pushToken)

	vendors.SendNotification(exponentTokens, vendors.NotificationMessage{
		Title: "تم تاكيد طلبك",
		Body:  "الرجاء متابعة الطلب ومتابعة السعر",
	}, vendors.NotificationData{
		ServiceID:    "0",
		SubServiceID: "0",
	})

	c.JSON(200, gin.H{
		"message": "Success",
	})
}

// CancelOrderFromController ..
func CancelOrderFromController(c *gin.Context) {
	ID := c.Param("id")

	var order models.Orders
	err := config.DB.Where("id = ?", ID).First(&order).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	order.Status = 3

	saveErr := config.DB.Save(&order).Error
	if saveErr != nil {
		c.JSON(500, gin.H{
			"error": saveErr.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Cancel success",
	})

}
