package api

import (
	"ce/backend/factory"
	"ce/backend/model"
	"ce/backend/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ApiAdminLogin(c *gin.Context) {
	Log.Info("API ADMIN", "login start")

	var request model.AdminLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handleAdminLogin(c, request)

	Log.Info("API ADMIN", "login end")
}

func ApiAdminGetBigGroups(c *gin.Context) {
	Log.Info("API ADMIN", "get big groups start")

	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "missing authorization header",
		})
		return
	}
	valid, _, err := util.VerifyJwtToken(jwtToken)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	handleAdminGetBigGroups(c)

	Log.Info("API ADMIN", "get big groups end")
}

func ApiAdminAddBigGroup(c *gin.Context) {
	Log.Info("API ADMIN", "add big group start")

	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "missing authorization header",
		})
		return
	}
	valid, _, err := util.VerifyJwtToken(jwtToken)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	var request model.AdminAddBigGroupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handleAdminAddBigGroup(c, request)

	Log.Info("API ADMIN", "add big group end")
}

func ApiAdminDeleteBigGroup(c *gin.Context) {
	Log.Info("API ADMIN", "delete big group start")

	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "missing authorization header",
		})
		return
	}
	valid, _, err := util.VerifyJwtToken(jwtToken)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	var request model.AdminDeleteBigGroupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handleAdminDeleteBigGroup(c, request)

	Log.Info("API ADMIN", "delete big group end")
}

func ApiAdminGetGroups(c *gin.Context) {
	Log.Info("API ADMIN", "get groups start")

	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "missing authorization header",
		})
		return
	}
	valid, _, err := util.VerifyJwtToken(jwtToken)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	var request model.AdminGetGroupsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handleAdminGetGroups(c, request)

	Log.Info("API ADMIN", "get groups end")
}

func ApiAdminAddGroup(c *gin.Context) {
	Log.Info("API ADMIN", "add group start")

	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "missing authorization header",
		})
		return
	}
	valid, _, err := util.VerifyJwtToken(jwtToken)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	var request model.AdminAddGroupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handleAdminAddGroup(c, request)

	Log.Info("API ADMIN", "add group end")
}

func ApiAdminDeleteGroup(c *gin.Context) {
	Log.Info("API ADMIN", "delete group start")

	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "missing authorization header",
		})
		return
	}
	valid, _, err := util.VerifyJwtToken(jwtToken)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	var request model.AdminDeleteGroupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handleAdminDeleteGroup(c, request)

	Log.Info("API ADMIN", "delete group end")
}

func ApiAdminGetGroupGrades(c *gin.Context) {
	Log.Info("API ADMIN", "get group start")

	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "missing authorization header",
		})
		return
	}
	valid, _, err := util.VerifyJwtToken(jwtToken)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	var request model.AdminGetGroupGradesRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handleAdminGetGroupGrades(c, request)

	Log.Info("API ADMIN", "get group end")
}

func handleAdminLogin(c *gin.Context, request model.AdminLoginRequest) {
	inTaList := false
	taIndex := -1
	for i, ta := range factory.CeConfig.Admin.Ta {
		if ta.TaId == request.TaId {
			inTaList = true
			taIndex = i
			break
		}
	}

	if !inTaList {
		Log.Error("API ADMIN", "you are not in the TA list")
		c.JSON(http.StatusNotFound, gin.H{"error": "you are not in the TA list"})
		return
	}

	if factory.CeConfig.Admin.Ta[taIndex].TaPassword != request.TaPassword {
		Log.Error("API ADMIN", "wrong password")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}

	jwtToken, err := util.GenerateJwtToken(request.TaId, time.Now().Unix(), time.Now().Unix()+AUTHORIZED_JWT_VALID_TIME, true)
	if err != nil {
		Log.Error("API ADMIN", "failed to generate jwt token: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := model.AdminLoginResponse{
		JwtToken: jwtToken,
	}

	c.JSON(http.StatusOK, response)
}

func handleAdminGetBigGroups(c *gin.Context) {
	bigGroups, err := GetBigGroups()
	if err != nil {
		Log.Error("API ADMIN", "failed to get big groups: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bigGroupNames := make([]string, len(bigGroups))
	for i, bigGroup := range bigGroups {
		bigGroupNames[i] = bigGroup.BigGroup
	}
	response := model.AdminGetBigGroupsResponse{
		BigGroups: bigGroupNames,
	}

	c.JSON(http.StatusOK, response)
}

func handleAdminAddBigGroup(c *gin.Context, request model.AdminAddBigGroupRequest) {
	bigGroup, err := GetBigGroup(request.BigGroup)
	if err != nil {
		Log.Error("API ADMIN", "failed to get big group: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if bigGroup != nil {
		Log.Error("API ADMIN", "big group already exists")
		c.JSON(http.StatusBadRequest, gin.H{"error": "big group already exists"})
		return
	}

	if err = InsertBigGroup(request.BigGroup); err != nil {
		Log.Error("API ADMIN", "failed to insert big group: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "big group added successfully"})
}

func handleAdminDeleteBigGroup(c *gin.Context, request model.AdminDeleteBigGroupRequest) {
	bigGroup, err := GetBigGroup(request.BigGroup)
	if err != nil {
		Log.Error("API ADMIN", "failed to get big group: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if bigGroup == nil {
		Log.Error("API ADMIN", "big group not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "big group not found"})
		return
	}

	if err = DeleteBigGroup(request.BigGroup); err != nil {
		Log.Error("API ADMIN", "failed to delete big group: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "big group deleted successfully"})
}

func handleAdminGetGroups(c *gin.Context, request model.AdminGetGroupsRequest) {
	bigGroup := request.BigGroup
	groups, err := GetGroups(bigGroup)
	if err != nil {
		Log.Error("API ADMIN", "failed to get groups: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := model.AdminGetGroupsResponse{
		Groups: make([]model.Group, len(groups)),
	}
	for i, group := range groups {
		response.Groups[i] = *group
	}

	c.JSON(http.StatusOK, response)
}

func handleAdminAddGroup(c *gin.Context, request model.AdminAddGroupRequest) {
	group, err := GetGroup(request.BigGroup, request.GroupId)
	if err != nil {
		Log.Error("API ADMIN", "failed to get group: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if group != nil {
		Log.Error("API ADMIN", "group already exists")
		c.JSON(http.StatusBadRequest, gin.H{"error": "group already exists"})
		return
	}

	if err = InsertGroup(&model.Group{
		BigGroup:        request.BigGroup,
		GroupId:         request.GroupId,
		GroupName:       request.GroupName,
		LeaderName:      request.LeaderName,
		LeaderStudentId: request.LeaderStudentId,
		Members:         request.Members,
		GradedStudentId: make(map[string]bool),
	}); err != nil {
		Log.Error("API ADMIN", "failed to insert group: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = InsertGroupGradeComment(&model.GroupGradeComment{
		BigGroup:              request.BigGroup,
		GroupId:               request.GroupId,
		GroupGradeCommentList: make([]model.GroupGradeCommentItem, 0),
	}); err != nil {
		Log.Error("API ADMIN", "failed to insert group grade comment: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "group added successfully"})
}

func handleAdminDeleteGroup(c *gin.Context, request model.AdminDeleteGroupRequest) {
	group, err := GetGroup(request.BigGroup, request.GroupId)
	if err != nil {
		Log.Error("API ADMIN", "failed to get group: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if group == nil {
		Log.Error("API ADMIN", "group not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
		return
	}

	if err = DeleteGroup(request.BigGroup, request.GroupId); err != nil {
		Log.Error("API ADMIN", "failed to delete group: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = DeleteGroupGradeComment(request.BigGroup, request.GroupId); err != nil {
		Log.Error("API ADMIN", "failed to delete group grade comment: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "group deleted successfully"})
}

func handleAdminGetGroupGrades(c *gin.Context, request model.AdminGetGroupGradesRequest) {
	groupGradeComment, err := GetGroupGradeComment(request.BigGroup, request.GroupId)
	if err != nil {
		Log.Error("API ADMIN", "failed to get group grade comment: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := model.AdminGetGroupGradesResponse{
		GroupGradeCommentList: groupGradeComment.GroupGradeCommentList,
	}

	c.JSON(http.StatusOK, response)
}
