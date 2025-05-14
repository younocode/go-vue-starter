package cache

import (
	"context"
	"time"
)

const emailPrefix = "email_"

func (r *RedisCache) GetEmailCode(ctx context.Context, email string) (string, error) {
	code := r.Client.Get(ctx, emailPrefix+email).Val()
	return code, nil
}

func (r *RedisCache) SetEmailCode(ctx context.Context, email string, code string) error {
	return r.Client.Set(ctx, emailPrefix+email, code, 5*time.Minute).Err()
}
