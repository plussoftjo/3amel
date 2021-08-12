// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// OrdersWithDetails ..
func OrdersWithDetails(db *gorm.DB) *gorm.DB {
	return db.Preload("Service").Preload("SubService").Preload("ServiceOptions")
}

// Services ..
type Orders struct {
	ServiceID         uint              `json:"serviceID"`
	SubServiceID      uint              `json:"subServiceID"`
	ServiceOptionsIDs pq.Int64Array     `json:"serviceOptionsIDs" gorm:"type:integer[]"`
	Date              string            `json:"date"`
	Time              string            `json:"time"`
	Location          string            `json:"location"`
	GeoLocation       string            `json:"geoLocation"`
	UserID            uint              `json:"userID"`
	UserRate          bool              `json:"userRate"`
	Cost              float64           `json:"cost"`
	Status            int64             `json:"status" gorm:"default:0"`
	SupplierID        uint              `json:"supplierID" gorm:"default:0"`
	SupplierRate      bool              `json:"supplierRate" gorm:"default:false"`
	Service           Services          `json:"service"`
	SubService        SubServices       `json:"subService"`
	ServiceOptions    []ServicesOptions `json:"serviceOptions" gorm:"many2many:order_services_option;"`
	gorm.Model
}
