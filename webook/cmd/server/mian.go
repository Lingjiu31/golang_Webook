package main

import (
	"Project-WeBook/webook/config"
	"Project-WeBook/webook/internal/handler"
	"Project-WeBook/webook/internal/repository"
	"Project-WeBook/webook/internal/repository/dao"
	"Project-WeBook/webook/internal/router"
	"Project-WeBook/webook/internal/service"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	if err != nil {
		// panic 相当于整个 goroutine 结束
		panic(err)
	}

	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}

	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := handler.NewUserHandler(svc)
	server := router.RegisterRoutes(u)

	err = server.Run(":8082")
	if err != nil {
		return
	}
}
