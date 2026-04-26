package router

import (
	"Project-WeBook/webook/internal/handler"
	"Project-WeBook/webook/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(u *handler.UserHandler) *gin.Engine {
	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	//server.Use(middleware.SessionMiddleware())
	//server.Use(middleware.NewLoginMiddlewareBuilder().
	//	IgnorePaths("/users/signup").
	//	IgnorePaths("/users/login").
	//	Build())
	server.Use(middleware.NewLoginJWTMiddlewareBuilder().
		IgnorePaths("/users/signup").
		IgnorePaths("/users/login").
		Build())

	ug := server.Group("/users")

	//注册
	ug.POST("/signup", u.SignUp)

	//登录
	ug.POST("/login", u.LoginJWT)

	//登出
	ug.POST("/logout", u.Logout)

	//编辑
	ug.POST("/edit", u.Edit)

	//用户信息
	ug.GET("/profile", u.ProFile)

	//测试用

	ug.GET("/test", u.Test)

	return server
}
