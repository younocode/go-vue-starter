package main

import (
	"github.com/labstack/echo/v4"
	"github.com/younocode/go-vue-starter/server/internal/router"
	"log/slog"
	"net/http"
)

type Router struct {
	UserRouter router.UserRouter
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) InitRouter(g *echo.Group) {
	g.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello GET World")
	})
	g.POST("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello POST World")
	})
	r.UserRouter.InitRouter(g)
	slog.Info("Router InitRouter")
}
