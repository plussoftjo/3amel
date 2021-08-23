// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// NotificationsToken ..
type NotificationsToken struct {
	gorm.Model
	UserID uint   `json:"userID"`
	Token  string `json:"token"`
	Active bool   `json:"active" gorm:"default:false"`
}
