package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/younocode/go-vue-starter/server/internal/cache"
	"github.com/younocode/go-vue-starter/server/internal/database"
	"github.com/younocode/go-vue-starter/server/internal/model"
	"github.com/younocode/go-vue-starter/server/internal/repo"
	"github.com/younocode/go-vue-starter/server/pkg/emailSender"
	"github.com/younocode/go-vue-starter/server/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) bool
}

type RandGenerator interface {
	GenerateEmailCode(len int) string
	GenerateRefreshToken(len int) string
}

type UserServicer interface {
	Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error)
	IsEmailAvailable(ctx context.Context, email string) error
	Register(ctx context.Context, req model.RegisterRequest) (*model.LoginResponse, error)
	SendEmailCode(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, req model.ForgetPasswordRequest) (*model.LoginResponse, error)
}

type UserService struct {
	db          *database.Database
	redisCache  *cache.RedisCache
	jwt         *jwt.JWT
	emailSender *emailSender.EmailSender
}

func NewUserService(db *database.Database, redisCache *cache.RedisCache, jwt *jwt.JWT, emailSender *emailSender.EmailSender) *UserService {
	return &UserService{
		db:          db,
		redisCache:  redisCache,
		jwt:         jwt,
		emailSender: emailSender,
	}
}

func (s *UserService) Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error) {
	user, err := s.db.Query.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed get user by emailSender: %w", err)
	}

	if !s.ComparePassword(user.Password, req.Password) {
		return nil, model.ErrUserNameOrPasswordFailed
	}

	accessToken, err := s.jwt.Generate(user.Email, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed generate access token: %w", err)
	}

	refreshToken, err := s.jwt.GenerateRefreshToken(user.Email, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed generate refresh token: %w", err)
	}

	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Email:        user.Email,
		UserID:       user.ID,
	}, nil
}

func (s *UserService) IsEmailAvailable(ctx context.Context, email string) error {
	available, err := s.db.Query.IsEmailAvailable(ctx, email)
	if err != nil {
		return err
	}
	if available {
		return model.ErrEmailAleadyExist
	}
	return nil
}

func (s *UserService) Register(ctx context.Context, req model.RegisterRequest) (*model.LoginResponse, error) {
	hash, err := s.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.db.Query.CreateUser(ctx, repo.CreateUserParams{
		Email:    req.Email,
		Password: hash,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	accessToken, err := s.jwt.Generate(user.Email, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed generate access token: %w", err)
	}

	refreshToken, err := s.jwt.GenerateRefreshToken(user.Email, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed generate refresh token: %w", err)
	}

	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Email:        user.Email,
		UserID:       user.ID,
	}, nil
}

func (s *UserService) SendEmailCode(ctx context.Context, email string) error {
	code := s.GenerateEmailCode(6)

	if err := s.redisCache.SetEmailCode(ctx, email, code); err != nil {
		return fmt.Errorf("failed to set emailcode in cache: %w", err)
	}

	if err := s.emailSender.Send(email, code); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (s *UserService) ResetPassword(ctx context.Context, req model.ForgetPasswordRequest) (*model.LoginResponse, error) {
	code, err := s.redisCache.GetEmailCode(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to get emailcode in cache: %w", err)
	}

	if code != req.EmailCode {
		return nil, model.ErrEmailCodeNotEqual
	}

	hash, err := s.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// 清除 RefreshToken
	user, err := s.db.Query.UpdatePasswordByEmail(ctx, repo.UpdatePasswordByEmailParams{
		Email:    req.Email,
		Password: hash,
	})

	return &model.LoginResponse{
		Email:  user.Email,
		UserID: user.ID,
	}, nil
}

func (s *UserService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *UserService) ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (s *UserService) GenerateEmailCode(n int) string {
	result := make([]byte, n)
	const letters = "0123456789"
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func (s *UserService) GenerateRefreshToken(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return ""
	}

	// 使用URL安全的Base64编码（去掉填充字符）
	token := base64.URLEncoding.EncodeToString(b)
	return token[:n]
}
