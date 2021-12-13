// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// SupplierInfo ..
type SupplierInfo struct {
	UserID     uint       `json:"userID"`
	Status     int        `json:"status" gorm:"default:0"`
	Latitude   float64    `json:"latitude" gorm:"default:0"`
	Longitude  float64    `json:"longitude" gorm:"default:0"`
	ServiceID  uint       `json:"serviceID" gorm:"default:0"`
	CategoryID uint       `json:"categoryID"`
	Service    Services   `json:"service" gorm:"foreignKey:ServiceID;references:ID"`
	Category   Categories `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	gorm.Model
}
