package router

import (
	"bookstore-go/config"
	"bookstore-go/web/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

type HttpRouter struct {
	userHandler *handler.UserHandler
}

func InitHttpRouter(userHandler *handler.UserHandler) error {
	httpRouter := &HttpRouter{
		userHandler: userHandler,
	}

	return httpRouter.init()
}

// 接口规则: restful
func (r *HttpRouter) init() error {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	// user module
	users := v1.Group("/users")
	users.POST("/register", r.userHandler.UserRegister)
	users.POST("/login", r.userHandler.UserLogin)

	if err := router.Run(fmt.Sprintf("%s:%d", config.AppConfig.Server.Host, config.AppConfig.Server.Port)); err != nil {
		return err
	}

	return nil
}
