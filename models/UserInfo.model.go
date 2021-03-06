// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// UserInfo ..
type UserInfo struct {
	UserID    uint      `json:"userID"`
	CountryID uint      `json:"countryID"`
	CityID    uint      `json:"cityID"`
	Country   Countries `json:"country"`
	City      Cites     `json:"city"`
	gorm.Model
}
