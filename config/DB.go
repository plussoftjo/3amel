// Package config ...
package config

import (
	"server/models"

	"github.com/jinzhu/gorm"
	// Connect mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// models
)

// SetupDB ...

// DB ..
var DB *gorm.DB

// SetupDB ..
func SetupDB() {
	database, err := gorm.Open("mysql", "root:00962s00962S!@tcp(127.0.0.1:3306)/services?charset=utf8mb4&parseTime=True&loc=Local")

	// If Error in Connect
	if err != nil {
		panic(err)
	}
	// User Models Setup
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.AuthClients{})
	database.AutoMigrate(&models.AuthTokens{})
	database.AutoMigrate(&models.Roles{})

	// Services
	database.AutoMigrate(&models.Services{})
	database.AutoMigrate(&models.SubServices{})
	database.AutoMigrate(&models.ServicesOptions{})

	// Orders
	database.AutoMigrate(&models.Orders{})

	DB = database
}