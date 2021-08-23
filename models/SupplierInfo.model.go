// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// SupplierInfo ..
type SupplierInfo struct {
	UserID    uint    `json:"userID"`
	Status    int     `json:"status" gorm:"default:0"`
	Latitude  float64 `json:"latitude" gorm:"default:0"`
	Longitude float64 `json:"longitude" gorm:"default:0"`
	User      User    `json:"user"`
	gorm.Model
}
