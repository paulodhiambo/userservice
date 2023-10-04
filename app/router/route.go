package router

import (
	"github.com/gin-gonic/gin"
	"structure/config"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		user := api.Group("/user")
		user.POST("/login", init.UserCtrl.GetAllUserData)
		user.POST("/register", init.UserCtrl.AddUserData)
		user.POST("/refreshToken", init.UserCtrl.RefreshToken)
		user.POST("/verifyEmail", init.UserCtrl.VerifyEmail)
		user.GET("/:userEmail", init.UserCtrl.GetUserByEmail)
		user.PUT("/:userEmail", init.UserCtrl.UpdateUserData)
		user.DELETE("/:userEmail", init.UserCtrl.DeleteUser)
	}
	return router
}
