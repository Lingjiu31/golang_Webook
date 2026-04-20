package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginMiddlewareBuilder struct{}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (b LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/users/login" || ctx.Request.URL.Path == "/users/signup" {
			return
		}
		session := sessions.Default(ctx)
		id := session.Get("userId")
		if id == nil {
			//未登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
