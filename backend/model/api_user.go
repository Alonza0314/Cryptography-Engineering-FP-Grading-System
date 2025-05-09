package model

type UserLoginRequest struct {
	StudentId string `json:"student_id" binding:"required"`
}

type UserLoginResponse struct {
	JwtToken string `json:"jwt_token"`
}