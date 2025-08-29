package main

import (
	"bookstore-go/config"
	"bookstore-go/global"
	"bookstore-go/pkg/mlog"
	"bookstore-go/repository"
	"bookstore-go/service"
	"bookstore-go/web/handler"
	"bookstore-go/web/router"
	"log"
)

func main() {
	// init log printf util
	mlog.Init(mlog.DefaultOptions())

	// init config
	config, err := config.NewConfig("./conf/config.yaml")
	if err != nil {
		mlog.Fatalf("failed to init config: %v", err)
		return
	}

	// init mysql
	global.InitMysql(config.Value().Mysql)

	// init redis
	global.InitRedis(config.Value().Redis)

	// init dao
	userDao := repository.NewUserDAO(global.GetDB())

	// init service
	userService := service.NewUserService(userDao)

	// init handler
	userHandler := handler.NewUserHandler(userService)

	// init router
	if err := router.InitHttpRouter(config, userHandler); err != nil {
		log.Fatalf("init http router fail: %s", err)
		return
	}
}
