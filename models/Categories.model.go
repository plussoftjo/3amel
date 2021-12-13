// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Categories ..
type Categories struct {
	Title string `json:"title"`
	gorm.Model
}
