// Package service 只关注业务规则实现
package service

import (
	"github.com/younocode/go-vue-starter/server/internal/cache"
	"github.com/younocode/go-vue-starter/server/internal/database"
	"github.com/younocode/go-vue-starter/server/pkg/emailSender"
	"github.com/younocode/go-vue-starter/server/pkg/jwt"
)

type Service struct {
	*UserService
}

func NewService(db *database.Database, redisCache *cache.RedisCache, jwt *jwt.JWT, emailSender *emailSender.EmailSender) *Service {
	userServicer := NewUserService(db, redisCache, jwt, emailSender)
	return &Service{
		UserService: userServicer,
	}
}
