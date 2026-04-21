package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc {
	//store := cookie.NewStore([]byte("secret"))
	store, err := redis.NewStore(16, "tcp",
		"localhost:6379", "", "",
		[]byte("h9X2&kL#vR6wP*zY4nM1jT8fQ5sC7aE9"),
		[]byte("h9X2&kL#vR6wP*zY4nM1jT8fQ5sC7aE9"))
	if err != nil {
		panic(err)
	}
	return sessions.Sessions("mysession", store)
}
