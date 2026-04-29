package middleware

import (
	"Project-WeBook/webook/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc {
	//store := cookie.NewStore([]byte("secret"))
	store, err := redis.NewStore(16, "tcp",
		config.Config.Redis.Addr, "", "",
		[]byte("h9X2&kL#vR6wP*zY4nM1jT8fQ5sC7aE9"),
		[]byte("7sB#2kQ!9zR$5mG&1dF@8cV%3jH^6*N"))
	if err != nil {
		panic(err)
	}
	return sessions.Sessions("mysession", store)
}
