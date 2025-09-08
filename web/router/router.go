package router

import (
	"bookstore-go/config"
	_ "bookstore-go/docs"
	"bookstore-go/pkg/utils"
	"bookstore-go/web/handler"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// health check
	router.GET("", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, nil)
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// user module
	users := v1.Group("/users")
	users.POST("/register", utils.Bind(handler.RegisterRequest{}), r.userHandler.UserRegister)
	users.POST("/login", utils.Bind(handler.LoginRequest{}), r.userHandler.UserLogin)

	if err := router.Run(fmt.Sprintf("%s:%d", r.config.Value().Server.Host, r.config.Value().Server.Port)); err != nil {
		return err
	}

	return nil
}
