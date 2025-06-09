package service

import (
	"singo/model"
	"singo/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService Manage user login service
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=40"`
}

// setSession Set session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login User login function
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("Username or password error", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("Username or password error", nil)
	}

	// Set session
	service.setSession(c, user)

	return serializer.BuildUserResponse(user)
}
