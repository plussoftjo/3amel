// Package controllers ...
package controllers

import (
	"fmt"
	"server/config"
	"server/models"
	"server/vendors"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginController ...
func LoginController(c *gin.Context) {

	var user models.User
	var login models.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if have user
	if err := config.DB.Preload("Roles").Where("phone = ?", login.Phone).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Check Password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	token, err := vendors.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"user":  user,
		"token": token,
	})
}

// RegisterController ...
func RegisterController(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	user.Password = string(hashedPassword)
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := vendors.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var users []models.User
	config.DB.Preload("Roles").Find(&users)
	config.DB.Preload("Roles").Where("id = ?", user.ID).First(&user)

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token, "users": users})
}

// AppRegisterController ...
func AppRegisterController(c *gin.Context) {
	type RegisterData struct {
		User     models.User     `json:"user"`
		UserInfo models.UserInfo `json:"userInfo"`
	}

	var data RegisterData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.User.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	data.User.Password = string(hashedPassword)
	if err := config.DB.Create(&data.User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": 101})
		return
	}

	token, err := vendors.CreateToken(data.User.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.UserInfo.UserID = data.User.ID
	config.DB.Create(&data.UserInfo)

	var User models.User
	config.DB.Preload("Roles").Preload("UserInfo").Where("id = ?", data.User.ID).First(&User)

	c.JSON(http.StatusOK, gin.H{"user": User, "token": token})
}

// Auth ..
func Auth(c *gin.Context) {
	user, err := AuthWithReturnUser(c.Request.Header["Authorization"][0])
	if err != nil {
		c.JSON(401, gin.H{
			"error": "UnAuthorized",
		})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

// AppAuth ..
func AppAuth(c *gin.Context) {
	user, err := AuthWithReturnUser(c.Request.Header["Authorization"][0])
	if err != nil {
		c.JSON(401, gin.H{
			"error": "UnAuthorized",
		})
		return
	}

	var User models.User
	config.DB.Where("id = ?", user.ID).First(&User)
	c.JSON(200, gin.H{
		"user": User,
	})
}

// AuthWithReturnUser ..
func AuthWithReturnUser(tok string) (*models.User, error) {
	// Auth
	token := strings.Split(tok, " ")[1]

	userID, err := vendors.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	var user models.User
	// Check if have user
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// AppUpdateUser ...
func AppUpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Password == "" {
		config.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(models.User{
			Name:  user.Name,
			Phone: user.Phone,
		})
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		config.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(models.User{
			Name:     user.Name,
			Phone:    user.Phone,
			Password: string(hashedPassword),
		})
	}

	var userData models.User
	config.DB.Where("id = ?", user.ID).First(&user)

	c.JSON(200, gin.H{
		"user": userData,
	})
}

// UpdateUser ...
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Password == "" {
		config.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(models.User{
			Name:    user.Name,
			Phone:   user.Phone,
			RolesID: user.RolesID,
		})
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		config.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(models.User{
			Name:     user.Name,
			Phone:    user.Phone,
			RolesID:  user.RolesID,
			Password: string(hashedPassword),
		})
	}

	var users []models.User
	config.DB.Preload("Roles").Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

// UsersListIndex ..
func UsersListIndex(c *gin.Context) {
	var users []models.User
	config.DB.Preload("Roles").Where("roles_id != ?", 0).Find(&users)
	c.JSON(200, gin.H{
		"users": users,
	})
}

// DeleteUser ...
func DeleteUser(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.User{}, ID)
	var users []models.User
	config.DB.Preload("Roles").Find(&users)
	c.JSON(200, gin.H{
		"users": users,
	})
}

// AuthAppUser ..
func AuthAppUser(c *gin.Context) {
	user, err := AuthWithReturnUser(c.Request.Header["Authorization"][0])
	if err != nil {
		c.JSON(401, gin.H{
			"error": "UnAuthorized",
		})
		return
	}

	var User models.User
	config.DB.Where("id = ?", user.ID).First(&User)

	var Orders []models.Orders
	config.DB.Scopes(models.OrdersWithDetails).Order("id desc").Where("user_id = ?", user.ID).Find(&Orders)

	c.JSON(200, gin.H{
		"user":   User,
		"orders": Orders,
	})
}

// AuthWorker ..
func AuthWorker(c *gin.Context) {
	user, err := AuthWithReturnUser(c.Request.Header["Authorization"][0])
	if err != nil {
		c.JSON(401, gin.H{
			"error": "UnAuthorized",
		})
		return
	}

	var User models.User
	config.DB.Where("id = ?", user.ID).First(&User)

	// if user type ..
	if User.UserType == 1 {
		var Orders []models.Orders
		config.DB.Scopes(models.OrdersWithDetails).Order("id desc").Where("supplier_id = ?", user.ID).Find(&Orders)

		c.JSON(200, gin.H{
			"user":   User,
			"orders": Orders,
		})
		return
	} else {
		c.JSON(401, gin.H{
			"error": "UnAuthorized",
		})
		return
	}

}

// AppLoginController ...
func AppLoginController(c *gin.Context) {

	var user models.User
	var login models.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if have user
	if err := config.DB.Preload("Roles").Where("phone = ?", login.Phone).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    100,
		})
		return
	}
	// Check Password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    101,
		})
		return
	}
	token, err := vendors.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var User models.User
	config.DB.Where("id = ?", user.ID).First(&User)

	var Orders []models.Orders
	config.DB.Scopes(models.OrdersWithDetails).Order("id desc").Where("user_id = ?", user.ID).Find(&Orders)

	c.JSON(200, gin.H{
		"user":   User,
		"token":  token,
		"orders": Orders,
	})
}

// WorkerLoginController ...
func WorkerLoginController(c *gin.Context) {

	var user models.User
	var login models.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if have user
	if err := config.DB.Preload("Roles").Where("phone = ?", login.Phone).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    100,
		})
		return
	}
	// Check Password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    101,
		})
		return
	}

	token, err := vendors.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var User models.User
	config.DB.Where("id = ?", user.ID).First(&User)

	// if the worker type is 1
	if User.UserType == 1 {
		var Orders []models.Orders
		config.DB.Scopes(models.OrdersWithDetails).Order("id desc").Where("supplier_id = ?", user.ID).Find(&Orders)

		c.JSON(200, gin.H{
			"user":   User,
			"token":  token,
			"orders": Orders,
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    102,
		})
		return
	}

}

// CheckIfHasPhone ..
func CheckIfHasPhone(c *gin.Context) {

	type PhoneData struct {
		Phone string `json:"phone"`
	}

	var data PhoneData
	c.ShouldBindJSON(&data)

	var user models.User
	// Check if have user
	if err := config.DB.Where("phone = ?", data.Phone).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "HasPhone",
		"user":    user,
	})

}

// ResetPassword ..
func ResetPassword(c *gin.Context) {
	type ResetType struct {
		UserID   uint   `json:"userID"`
		Password string `json:"password"`
	}

	var data ResetType
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := config.DB.Where("id = ?", data.UserID).First(&user).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	user.Password = string(hashedPassword)

	config.DB.Save(&user)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

// ChangePassword ..
func ChangePassword(c *gin.Context) {
	type ChangePasswordType struct {
		UserID      uint   `json:"userID"`
		Password    string `json:"password"`
		OldPassword string `json:"oldPassword"`
	}
	var data ChangePasswordType
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	config.DB.Where("id = ?", data.UserID).First(&user)

	// Check Password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.OldPassword))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    101,
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	user.Password = string(hashedPassword)
	config.DB.Save(&user)

	c.JSON(200, gin.H{
		"user": user,
		"code": 100,
	})
}
