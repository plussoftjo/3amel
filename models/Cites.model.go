// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Cites ..
type Cites struct {
	gorm.Model
	Title     string    `json:"title"`
	CountryID uint      `json:"countryID"`
	Country   Countries `json:"country" gorm:"foreignKey:CountryID;references:ID"`
}
