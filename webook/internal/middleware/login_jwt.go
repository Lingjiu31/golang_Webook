package middleware

import (
	"Project-WeBook/webook/internal/handler"
	"encoding/gob"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginJWTMiddlewareBuilder struct {
	paths []string
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{}
}
func (l *LoginJWTMiddlewareBuilder) IgnorePaths(path string) *LoginJWTMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		// 无需登录的 地址
		for _, path := range l.paths {
			if ctx.Request.RequestURI == path {
				return
			}
		}
		// 使用jwt校验
		tokenHeader := ctx.GetHeader("Authorization")
		if tokenHeader == "" {
			// 未登录(Authorization为空)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 拿到的 Authorization 有格式中的空格, 切掉
		segs := strings.Split(tokenHeader, " ")
		if len(segs) != 2 {
			// 正常一定是2段
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 第二段为token
		tokenStr := segs[1]
		claims := &handler.UserClaims{}
		// 这里是因为要给id赋值, 不是指针就只是复制一份, 所以传指针
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("9sK7$pR2!zG5&qB8@tN3#mC6%vH1*dJ4"), nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token == nil || !token.Valid || claims.UserId == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if claims.UserAgent != ctx.Request.UserAgent() {
			// 出现安全问题
			// 需要加一个监控
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 还剩7天就刷新
		now := time.Now()
		if claims.ExpiresAt.Sub(now) < time.Hour*24*7 {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30))
			tokenStr, err = token.SignedString([]byte("9sK7$pR2!zG5&qB8@tN3#mC6%vH1*dJ4"))
			if err != nil {
				// 记录日志
				log.Println("登录状态刷新失败", err)
			}
			ctx.Header("x-jwt-token", tokenStr)
		}

		ctx.Set("userid", claims.UserId)
	}
}
