package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User User model
type User struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Nickname       string
	Status         string
	Avatar         string `gorm:"size:1000"`
}

const (
	// PassWordCost Password encryption difficulty
	PassWordCost = 12
	// Active Active user
	Active string = "active"
	// Inactive Inactive user
	Inactive string = "inactive"
	// Suspend Suspended user
	Suspend string = "suspend"
)

// GetUser Get user by ID
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword Set password
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword Verify password
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
