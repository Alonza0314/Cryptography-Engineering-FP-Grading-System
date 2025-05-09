package api

import (
	"ce/backend/model"
	"ce/backend/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ApiUserLogin(c *gin.Context) {
	Log.Info("API USER", "login start")

	var request model.UserLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handleUserLogin(c, request)

	Log.Info("API USER", "login end")
}

func ApiUserTOTPInitBegin(c *gin.Context) {
	Log.Info("API USER", "totp init begin start")

	Log.Info("API USER", "totp init begin end")
}

func ApiUserTOTPInitFinish(c *gin.Context) {
	Log.Info("API USER", "totp init finish start")

	Log.Info("API USER", "totp init finish end")
}

func ApiUserTOTP(c *gin.Context) {
	Log.Info("API USER", "totp start")

	Log.Info("API USER", "totp end")
}

func handleUserLogin(c *gin.Context, request model.UserLoginRequest) {
	userAccount, err := GetUserAccount(request.StudentId)
	if err != nil {
		Log.Error("API USER", "failed to get user account: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var jwtToken string
	var totp bool
	if userAccount == nil || !userAccount.Totp {
		totp = false
	} else {
		totp = true
	}
	jwtToken, err = util.GenerateJwtToken(request.StudentId, time.Now().Unix(), time.Now().Unix()+1000*60*5, totp)
	if err != nil {
		Log.Error("API USER", "failed to generate jwt token: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := model.UserLoginResponse{
		JwtToken: jwtToken,
	}

	c.JSON(http.StatusOK, response)
}
