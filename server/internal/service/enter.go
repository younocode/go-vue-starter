// Package service 只关注业务规则实现
package service

import (
	"database/sql"
	"github.com/younocode/go-vue-starter/server/internal/cache"
	"github.com/younocode/go-vue-starter/server/pkg/emailSender"
	"github.com/younocode/go-vue-starter/server/pkg/jwt"
)

type Service struct {
	*UserService
}

func NewService(db *sql.DB, redisCache *cache.RedisCache, jwt *jwt.JWT, emailSender *emailSender.EmailSender) *Service {
	userServicer := NewUserService(db, redisCache, jwt, emailSender)
	return &Service{
		UserService: userServicer,
	}
}
