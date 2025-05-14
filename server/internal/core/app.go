package core

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/younocode/go-vue-starter/server/config"
	"github.com/younocode/go-vue-starter/server/internal/cache"
	"github.com/younocode/go-vue-starter/server/internal/handler"
	"github.com/younocode/go-vue-starter/server/internal/service"
	"github.com/younocode/go-vue-starter/server/pkg/emailSender"
	"github.com/younocode/go-vue-starter/server/pkg/jwt"
)

type App struct {
	DB          *sql.DB
	RedisCache  *cache.RedisCache
	E           *echo.Echo
	Service     *service.Service
	Handler     *handler.Handler
	Cfg         *config.Config
	Jwt         *jwt.JWT
	EmailSender *emailSender.EmailSender
}
