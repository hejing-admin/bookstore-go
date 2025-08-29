package handler

import (
	"bookstore-go/pkg/mlog"
	"bookstore-go/service"
	"bookstore-go/web/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	userService *service.UserService
}

type RegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) UserRegister(c *gin.Context) {
	// step 1. param verify
	req := c.MustGet(gin.BindKey).(*RegisterRequest)

	// step 2.  todo 验证码的校验

	// step 3. 验证两次密码是否一致
	if req.Password != req.ConfirmPassword {
		response.BadRequest(c, "passwords do not match", nil)
		return
	}

	// step 3. call service to register user
	if err := u.userService.UserRegister(req.Username, req.Password, req.Email, req.Phone); err != nil {
		response.InternalError(c, "user register fail", err)
		return
	}

	response.Success(c, "user register success", nil)
}

func (u *UserHandler) UserLogin(c *gin.Context) {

}
