// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// UserImages ..
type UserImages struct {
	UserID uint   `json:"userID"`
	Image  string `json:"image"`
	gorm.Model
}
