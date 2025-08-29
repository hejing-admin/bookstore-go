package main

import (
	"bookstore-go/config"
	"bookstore-go/global"
	"bookstore-go/repository"
	"bookstore-go/service"
	"bookstore-go/web/handler"
	"bookstore-go/web/router"
	"log"
)

func main() {
	// init config
	config.InitConfig("./conf/config.yaml")

	// init mysql
	// global.InitMysql(config.AppConfig.Mysql)

	// init redis
	// global.InitRedis(config.AppConfig.Redis)

	// init dao
	userDao := repository.NewUserDAO(global.GetDB())

	// init service
	userService := service.NewUserService(userDao)

	// init handler
	userHandler := handler.NewUserHandler(userService)

	// init router
	if err := router.InitHttpRouter(userHandler); err != nil {
		log.Fatalf("init http router fail: %s", err)
		return
	}
}
