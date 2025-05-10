package model

type AdminLoginRequest struct {
	TaId     string `json:"ta_id" binding:"required"`
	TaPassword string `json:"ta_password" binding:"required"`
}

type AdminLoginResponse struct {
	JwtToken string `json:"jwt_token" binding:"required"`
}
