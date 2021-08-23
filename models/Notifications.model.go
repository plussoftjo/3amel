// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Notifications ..
type Notifications struct {
	gorm.Model
	Title        string      `json:"title"`
	Body         string      `json:"body"`
	Data         string      `json:"data"`
	ServiceID    int         `json:"serviceID"`
	SubServiceID int         `json:"subServiceID" gorm:"default:0"`
	Image        string      `json:"image"`
	Service      Services    `json:"service"`
	SubService   SubServices `json:"subService"`
}
