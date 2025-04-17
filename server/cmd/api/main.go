package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/younocode/go-vue-starter/server/config"
	"github.com/younocode/go-vue-starter/server/internal/cache"
	"github.com/younocode/go-vue-starter/server/internal/database"
	"log"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

type App struct {
	e           *echo.Echo
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
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}

	// 注册路由

	app := &App{
		e:           e,
		db:          db,
		redisClient: redisCache,
	}

	return app, err
}

func main() {
	var err error
	app, err := InitApp("./config/config.yaml")
	if err != nil {
		panic(fmt.Sprintf("init app error: %s", err.Error()))
	}

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(app.e.Server, done)

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")
}
