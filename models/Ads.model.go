// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Ads ..
type Ads struct {
	gorm.Model
	Title        string      `json:"title"`
	Image        string      `json:"image"`
	ServiceID    uint        `json:"serviceID"`
	SubServiceID uint        `json:"subServiceID"`
	Service      Services    `json:"service"`
	SubService   SubServices `json:"subService"`
}
