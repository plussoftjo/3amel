// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// UserRate ..
type UserRate struct {
	UserID  uint   `json:"userID"`
	RaterID uint   `json:"raterID"`
	Note    string `json:"note"`
	Value   string `json:"value"`
	OrderID uint   `json:"orderID"`
	Paid    string `json:"paid"`
	gorm.Model
}
