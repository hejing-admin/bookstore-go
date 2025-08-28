package main

import (
	"bookstore-go/config"
	"bookstore-go/global"
	"bookstore-go/repository"
	"bookstore-go/service"
	"bookstore-go/web/handler"
	"bookstore-go/web/router"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// init config
	config.InitConfig("./conf/config.yaml")

	// init mysql
	global.InitMysql(config.AppConfig.Mysql)

	// init redis
	global.InitRedis(config.AppConfig.Redis)

	// init dao
	userDao := repository.NewUserDAO(global.GetDB())

	// init service
	userService := service.NewUserService(userDao)

	// init handler
	userHandler := handler.NewUserHandler(userService)

	// init router
	rRouter := router.NewRouter(userHandler)
	r := rRouter.InitRouter()
	addr := fmt.Sprintf("%s:%d", config.AppConfig.Server.Host, config.AppConfig.Server.Port)

	// start http server
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("server listen err:", err)
		os.Exit(1)
	}

}
