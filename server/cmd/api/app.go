package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/younocode/go-vue-starter/server/config"
	"github.com/younocode/go-vue-starter/server/internal/cache"
	"github.com/younocode/go-vue-starter/server/internal/core"
	"github.com/younocode/go-vue-starter/server/internal/database"
	"github.com/younocode/go-vue-starter/server/internal/handler"
	"github.com/younocode/go-vue-starter/server/internal/service"
	"github.com/younocode/go-vue-starter/server/pkg/emailSender"
	"github.com/younocode/go-vue-starter/server/pkg/jwt"
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

func InitApp(file string) (*core.App, error) {
	var err error
	cfg, err := config.NewConfig(file)
	if err != nil {
		return nil, err
	}
	// 初始化数据库
	db, err := database.NewDB(cfg.Database)
	if err != nil {
		return nil, err
	}

	// 初始化 redis
	redisCache, err := cache.NewRedisCache(cfg.Redis)
	if err != nil {
		return nil, err
	}

	// 初始化 jwt
	jr := jwt.NewJWT(cfg.JWT)

	// 初始化 emailSender
	es := emailSender.NewEmailSend(cfg.Email)

	// 初始化 echo
	e := NewEcho()
	e.Server = &http.Server{
		Handler:      e,                // Echo 本身实现了 http.h
		ReadTimeout:  10 * time.Second, // 读取超时
		WriteTimeout: 30 * time.Second, // 写入超时
		IdleTimeout:  1 * time.Minute,  // 空闲连接超时
	}

	s := service.NewService(db, redisCache, jr, es)
	h := handler.NewHandler(s)

	app := &core.App{
		Cfg:         cfg,
		E:           e,
		DB:          db,
		RedisCache:  redisCache,
		Jwt:         jr,
		Service:     s,
		Handler:     h,
		EmailSender: es,
	}

	r := NewRouter()
	r.InitRouter(e.Group(""))

	return app, err
}

func Start() {
	var err error
	app, err := InitApp("./config/config.yaml")
	if err != nil {
		panic(fmt.Sprintf("init app error: %s", err.Error()))
	}

	// start
	if err = app.E.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(app.E.Server, done)

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")
}
