package handler

import "github.com/younocode/go-vue-starter/server/internal/service"

type Handler struct {
	*UserHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(service),
	}
}
