// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Services ..
type Services struct {
	Title          string            `json:"title"`
	Image          string            `json:"image"`
	SubServices    []SubServices     `json:"subServices" gorm:"foreignKey:ServiceID;references:ID"`
	ServiceOptions []ServicesOptions `json:"serviceOptions" gorm:"foreignKey:ServiceID;references:ID"`
	gorm.Model
}
