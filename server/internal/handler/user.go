package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/younocode/go-vue-starter/server/internal/model"
	"github.com/younocode/go-vue-starter/server/internal/service"
	"net/http"
)

type UserHandler struct {
	service *service.Service
}

func NewUserHandler(service *service.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Login(c echo.Context) error {
	res, err := h.service.UserService.Login(c.Request().Context(), model.LoginRequest{
		Email:    c.FormValue("emailSender"),
		Password: c.FormValue("password"),
	})
	if err != nil {
		if errors.Is(err, model.ErrUserNameOrPasswordFailed) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Register(c echo.Context) error {
	res, err := h.service.UserService.Register(c.Request().Context(), model.RegisterRequest{
		LoginRequest: model.LoginRequest{
			Email:    c.FormValue("emailSender"),
			Password: c.FormValue("password"),
		},
		EmailCode: c.FormValue("code"),
	})
	if err != nil {
		if errors.Is(err, model.ErrEmailCodeNotEqual) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) SendEmailCode(c echo.Context) error {
	err := h.service.UserService.SendEmailCode(c.Request().Context(), c.FormValue("email"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "发送成功")
}

func (h *UserHandler) ForgetPassword(c echo.Context) error {
	var req model.ForgetPasswordRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.service.UserService.ResetPassword(c.Request().Context(), req)
	if err != nil {
		if errors.Is(err, model.ErrEmailCodeNotEqual) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)

}
