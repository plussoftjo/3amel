// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// ServicesOptions ..
type ServicesOptions struct {
	Title        string      `json:"title"`
	ServiceID    uint        `json:"serviceID"`
	SubServiceID uint        `json:"subServiceID"`
	Service      Services    `json:"service"`
	SubService   SubServices `json:"subService"`
	gorm.Model
}
