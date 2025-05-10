package model

type UserLoginRequest struct {
	StudentId string `json:"student_id" binding:"required"`
}

type UserLoginResponse struct {
	JwtToken string `json:"jwt_token" binding:"required"`
}

type UserTOTPInitBeginResponse struct {
	QrCode     string `json:"qr_code" binding:"required"`
	SecretKey  string `json:"secret_key" binding:"required"`
}

type UserTOTPInitFinishRequest struct {
	Code string `json:"code" binding:"required"`
}

type UserTOTPRequest struct {
	Code string `json:"code" binding:"required"`
}

type UserTOTPResponse struct {
	JwtToken string `json:"jwt_token" binding:"required"`
}
