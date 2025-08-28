package router

import (
	"bookstore-go/web/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	userHandler *handler.UserHandler
}

func NewRouter(userHandler *handler.UserHandler) *Router {
	return &Router{
		userHandler: userHandler,
	}
}

func (router *Router) InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/register", router.userHandler.UserRegister)
			user.POST("/login", router.userHandler.UserLogin)
		}
	}

	return r
}
