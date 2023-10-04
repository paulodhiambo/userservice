package controller

import (
	"github.com/gin-gonic/gin"
	"structure/app/service"
)

type UserController interface {
	GetAllUserData(c *gin.Context)
	AddUserData(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
	RefreshToken(c *gin.Context)
	VerifyEmail(c *gin.Context)
}

type UserControllerImpl struct {
	svc service.UserService
}

func (u UserControllerImpl) GetAllUserData(c *gin.Context) {
	u.svc.GetAllUser(c)
}

func (u UserControllerImpl) AddUserData(c *gin.Context) {
	u.svc.AddUserData(c)
}

func (u UserControllerImpl) GetUserByEmail(c *gin.Context) {
	u.svc.GetUserById(c)
}

func (u UserControllerImpl) UpdateUserData(c *gin.Context) {
	u.svc.UpdateUserData(c)
}

func (u UserControllerImpl) DeleteUser(c *gin.Context) {
	u.svc.DeleteUser(c)
}

func (u UserControllerImpl) RefreshToken(c *gin.Context) {
	u.svc.DeleteUser(c)
}
func (u UserControllerImpl) VerifyEmail(c *gin.Context) {
	u.svc.DeleteUser(c)
}

func UserControllerInit(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		svc: userService,
	}
}
