package model

type AdminLoginRequest struct {
	TaId       string `json:"ta_id" binding:"required"`
	TaPassword string `json:"ta_password" binding:"required"`
}

type AdminLoginResponse struct {
	JwtToken string `json:"jwt_token" binding:"required"`
}

type AdminGetBigGroupsResponse struct {
	BigGroups []string `json:"big_groups" binding:"required"`
}

type AdminAddBigGroupRequest struct {
	BigGroup string `json:"big_group" binding:"required"`
}

type AdminDeleteBigGroupRequest struct {
	BigGroup string `json:"big_group" binding:"required"`
}

type AdminGetGroupsRequest struct {
	BigGroup string `json:"big_group" binding:"required"`
}

type AdminGetGroupsResponse struct {
	Groups []Group `json:"groups" binding:"required"`
}

type AdminAddGroupRequest struct {
	Group
}

type AdminDeleteGroupRequest struct {
	BigGroup string `json:"big_group" binding:"required"`
	GroupId  int    `json:"group_id" binding:"required"`
}
