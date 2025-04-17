package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "health")
}

func (s *Server) hiSqlcHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "hiSqlc")
}
