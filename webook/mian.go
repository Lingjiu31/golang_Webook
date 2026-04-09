package main

import (
	"Project-WeBook/webook/internal/web"
)

func main() {
	server := web.RegisterRoutes()
	err := server.Run(":8080")
	if err != nil {
		return
	}
}
