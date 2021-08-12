// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// SubServices ..
type SubServices struct {
	Title          string            `json:"title"`
	Image          string            `json:"image"`
	ServiceID      uint              `json:"serviceID"`
	Service        Services          `json:"service"`
	ServiceOptions []ServicesOptions `json:"serviceOptions" gorm:"foreignKey:SubServiceID;references:ID"`
	gorm.Model
}
