// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// User ..
type User struct {
	gorm.Model
	Name         string       `json:"name"`
	Phone        string       `json:"phone" gorm:"unique"`
	Password     string       `json:"password"`
	Email        string       `json:"email"`
	RolesID      uint         `json:"roles_id"`
	UserType     uint         `json:"user_type"`               // 01 -> User , 02 -> Supplier, 03 -> Controller
	Status       int64        `json:"status" gorm:"default:2"` // 01 -> notActive , 02 -> active, 03 -> blocked, 04 -> On Work
	Avatar       string       `json:"avatar"`
	Roles        Roles        `json:"roles" gorm:"foreignKey:RolesID;references:ID"`
	UserInfo     UserInfo     `json:"userInfo" gorm:"foreignKey:UserID;references:ID"`
	SupplierInfo SupplierInfo `json:"supplierInfo" gorm:"foreignKey:UserID;references:ID"`
	UserImages   []UserImages `json:"userImages" gorm:"foreignKey:UserID;references:ID"`
}

// Login ...
type Login struct {
	Phone    string `json:"phone" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
}
