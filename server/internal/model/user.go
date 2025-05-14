package model

type LoginRequest struct {
	Email    string `json:"emailSender" validate:"required,emailSender"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	Email       string `json:"emailSender"`
	UserID      int32  `json:"user_id"`
}

type RegisterRequest struct {
	LoginRequest
	EmailCode string `json:"email_code" validate:"required,len=6"`
}

type ForgetPasswordRequest struct {
	LoginRequest
	EmailCode string `json:"email_code" validate:"required,len=6"`
}

type SendCodeRequest struct {
	Email string `validate:"required,emailSender"`
}
