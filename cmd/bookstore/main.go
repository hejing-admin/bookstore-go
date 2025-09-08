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

//	@title			BookStore Api Server Swagger API
//	@version		1.0
//	@description	BookStore API Server

//	@host		localhost:19101
//	@BasePath	/api/v1

// @securityDefinitions.basic	BasicAuth
func main() {
	// init log printf util
	mlog.Init(mlog.DefaultOptions())

	// Note: swagger接入：https://theguodong.com/articles/Gin/Gin%E4%BD%BF%E7%94%A8swagger%E7%94%9F%E6%88%90%E6%96%87%E6%A1%A3/

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
