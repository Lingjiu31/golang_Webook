package middleware

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
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
			// k8s 环境
			if strings.Contains(origin, "webook.test") {
				return true
			}
			// 实际环境
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	})
}
