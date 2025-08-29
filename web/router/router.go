package router

import (
	"bookstore-go/config"
	"bookstore-go/utils"
	"bookstore-go/web/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	ApiV1 = "/api/v1"
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

	v1 := router.Group(ApiV1)

	// user module
	users := v1.Group("/users")
	users.POST("/register", utils.Bind(handler.RegisterRequest{}), r.userHandler.UserRegister)
	users.POST("/login", utils.Bind(handler.LoginRequest{}), r.userHandler.UserLogin)

	if err := router.Run(fmt.Sprintf("%s:%d", config.AppConfig.Server.Host, config.AppConfig.Server.Port)); err != nil {
		return err
	}

	return nil
}
