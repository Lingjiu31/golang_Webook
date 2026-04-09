package web

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册线路,路由
func RegisterRoutes() *gin.Engine {
	server := gin.Default()

	// CORS 跨域配置
	server.Use(cors.New(cors.Config{
		// 由于使用下面那个,所以不用这个了
		//AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods:  []string{"POST", "GET"},
		AllowHeaders:  []string{"Content-Type", "Authorization"},
		ExposeHeaders: []string{"x-jwt-token"},
		//是否允许你带 cookie 之类的
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// origin 有这个前缀就能传输(开发环境)
			if strings.Contains(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	u := NewUserHandler()

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
