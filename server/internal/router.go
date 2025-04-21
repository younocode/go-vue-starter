package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/younocode/go-vue-starter/server/internal/server"
	"log/slog"
	"net/http"
)

func (app *App) initRouter() {
	root := app.E.Group("")
	root.GET("/hi", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "hi")
	})
	s := server.NewServer()
	s.InitRouter(root)
	slog.Info("server InitRouter")
}
