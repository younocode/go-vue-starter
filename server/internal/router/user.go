package router

import (
	"github.com/labstack/echo/v4"
	"github.com/younocode/go-vue-starter/server/internal/handler"
	"log/slog"
)

type UserRouter struct {
	UserHandler *handler.UserHandler
}

func NewUserRouter(handler *handler.Handler) *UserRouter {
	return &UserRouter{
		UserHandler: handler.UserHandler,
	}
}

func (r UserRouter) InitRouter(g *echo.Group) {
	g.POST("/login", r.UserHandler.Login)
	g.POST("/register", r.UserHandler.Register)
	g.POST("/send-email-code", r.UserHandler.SendEmailCode)
	g.POST("/forget-password", r.UserHandler.ForgetPassword)
	slog.Info("User InitRouter")
}
