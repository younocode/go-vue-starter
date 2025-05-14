package model

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Email        string `json:"email"`
	UserID       int32  `json:"user_id"`
}

type RegisterRequest struct {
	LoginRequest
	EmailCode string `json:"email_code" validate:"required,len=6"`
}

type ForgetPasswordRequest struct {
	LoginRequest
	EmailCode string `json:"email_code" validate:"required,len=6"`
}

type SendEmailCodeRequest struct {
	Email string `json:"email" validate:"required,email"`
}
