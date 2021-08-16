// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// AppIntro ..
type AppIntro struct {
	gorm.Model
	Title   string `json:"title"`
	Color   string `json:"color"`
	Image   string `json:"image"`
	Content string `json:"content"`
}
