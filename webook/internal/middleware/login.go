package middleware

import (
	"encoding/gob"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		// 无需登录地址
		for _, path := range l.paths {
			if ctx.Request.RequestURI == path {
				return
			}
		}

		session := sessions.Default(ctx)
		id := session.Get("userId")
		if id == nil {
			// 未登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		updateTime := session.Get("updateTime")
		session.Set("userId", id)
		now := time.Now()
		// 如果是空说明没有刷新过, 刚刚登录
		if updateTime == nil {
			session.Set("updateTime", now)
			session.Options(sessions.Options{
				MaxAge: 60,
			})
			err := session.Save()
			if err != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			return
		}

		// 已经刷新过, 再次刷新
		updateTimeVal, ok := updateTime.(time.Time)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if now.Sub(updateTimeVal) > time.Minute {
			session.Set("updateTime", now)
			err := session.Save()
			if err != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}
	}
}
