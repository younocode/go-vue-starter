package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/younocode/go-vue-starter/server/internal/cache"
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

type UserServicer interface {
	Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error)
	IsEmailAvaliable(ctx context.Context, email string) error
	Register(ctx context.Context, req model.RegisterRequest) (*model.LoginResponse, error)
	GenerateEmailCode(len int) string
	SendEmailCode(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, req model.ForgetPasswordRequest) (*model.LoginResponse, error)
}

type UserService struct {
	queries     *repo.Queries
	db          *sql.DB
	redisCache  *cache.RedisCache
	jwt         *jwt.JWT
	emailSender *emailSender.EmailSender
}

func NewUserService(db *sql.DB, redisCache *cache.RedisCache, jwt *jwt.JWT, emailSender *emailSender.EmailSender) *UserService {
	return &UserService{
		queries:     repo.New(db),
		db:          db,
		redisCache:  redisCache,
		jwt:         jwt,
		emailSender: emailSender,
	}
}

func (s *UserService) Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error) {
	user, err := s.queries.GetUserByEmail(ctx, req.Email)
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

	return &model.LoginResponse{
		AccessToken: accessToken,
		Email:       user.Email,
		UserID:      user.ID,
	}, nil
}

func (s *UserService) IsEmailAvaliable(ctx context.Context, email string) error {
	available, err := s.queries.IsEmailAvailable(ctx, email)
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

	user, err := s.queries.CreateUser(ctx, repo.CreateUserParams{
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

	return &model.LoginResponse{
		AccessToken: accessToken,
		Email:       user.Email,
		UserID:      user.ID,
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

	user, err := s.queries.UpdatePasswordByEmail(ctx, repo.UpdatePasswordByEmailParams{
		Email:    req.Email,
		Password: hash,
	})

	accessToken, err := s.jwt.Generate(user.Email, user.ID)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		AccessToken: accessToken,
		Email:       user.Email,
		UserID:      user.ID,
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

func (s *UserService) GenerateEmailCode(len int) string {
	result := make([]byte, len)
	const nums = "0123456789"
	for i := range result {
		result[i] = nums[rand.Intn(10)]
	}

	return string(result)
}
