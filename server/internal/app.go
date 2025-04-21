package internal

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/younocode/go-vue-starter/server/config"
	"github.com/younocode/go-vue-starter/server/internal/cache"
	"github.com/younocode/go-vue-starter/server/internal/database"
	"log/slog"
	"net/http"
	"time"
)

type App struct {
	E           *echo.Echo
	db          *sql.DB
	redisClient *redis.Client
}

func InitApp(file string) (*App, error) {
	var err error
	cfg, err := config.NewConfig(file)
	if err != nil {
		return nil, err
	}
	// 初始化数据库
	db, err := database.NewPgSql(cfg.Database)
	if err != nil {
		return nil, err
	}

	// 初始化 redis
	redisCache, err := cache.NewRedisClient(cfg.Redis)
	if err != nil {
		return nil, err
	}

	// 初始化 echo
	e := NewEcho()
	e.Server = &http.Server{
		Handler:      e,                // Echo 本身实现了 http.Handler
		ReadTimeout:  10 * time.Second, // 读取超时
		WriteTimeout: 30 * time.Second, // 写入超时
		IdleTimeout:  1 * time.Minute,  // 空闲连接超时
	}

	app := &App{
		E:           e,
		db:          db,
		redisClient: redisCache,
	}

	// 注册路由
	app.initRouter()
	slog.Info("app InitRouter")
	return app, err
}
