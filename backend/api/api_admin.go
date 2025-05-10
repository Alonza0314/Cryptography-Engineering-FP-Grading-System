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
