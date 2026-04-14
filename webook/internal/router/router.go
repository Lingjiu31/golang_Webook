package router

import (
	"Project-WeBook/webook/internal/handler"
	"Project-WeBook/webook/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(u *handler.UserHandler) *gin.Engine {
	server := gin.Default()

	server.Use(middleware.CORSMiddleware())

	ug := server.Group("/users")

	//注册
	ug.POST("/signup", u.SignUp)

	//登录
	ug.POST("/login", u.Login)

	//编辑
	ug.POST("/edit", u.Edit)

	//用户信息
	ug.GET("/profile", u.ProFile)

	return server
}
