package router

import (
	"bookstore-go/config"
	"bookstore-go/pkg/utils"
	"bookstore-go/web/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	ApiV1 = "/api/v1"
)

type HttpRouter struct {
	config      config.Config
	userHandler *handler.UserHandler
}

func InitHttpRouter(config config.Config, userHandler *handler.UserHandler) error {
	httpRouter := &HttpRouter{
		config:      config,
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

	if err := router.Run(fmt.Sprintf("%s:%d", r.config.Value().Server.Host, r.config.Value().Server.Port)); err != nil {
		return err
	}

	return nil
}
