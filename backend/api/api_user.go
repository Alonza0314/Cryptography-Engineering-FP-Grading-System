package api

import (
	"ce/backend/model"
	"ce/backend/util"
	"net/http"
	"time"

	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
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

	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "missing authorization header",
		})
		return
	}
	valid, studentId, err := util.VerifyJwtToken(jwtToken)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}
	handleUserTOTPInitBegin(c, studentId)

	Log.Info("API USER", "totp init begin end")
}

func ApiUserTOTPInitFinish(c *gin.Context) {
	Log.Info("API USER", "totp init finish start")

	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "missing authorization header",
		})
		return
	}
	valid, studentId, err := util.VerifyJwtToken(jwtToken)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	var request model.UserTOTPInitFinishRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handleUserTOTPInitFinish(c, studentId, request.Code)

	Log.Info("API USER", "totp init finish end")
}

func ApiUserTOTP(c *gin.Context) {
	Log.Info("API USER", "totp start")

	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "missing authorization header",
		})
		return
	}
	valid, studentId, err := util.VerifyJwtToken(jwtToken)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	var request model.UserTOTPRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handleUserTOTP(c, studentId, request.Code)

	Log.Info("API USER", "totp end")
}

func ApiUserGetBigGroups(c *gin.Context) {
	Log.Info("API USER", "get big groups start")

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

	handleUserGetBigGroups(c)

	Log.Info("API USER", "get big groups end")
}

func ApiUserGetGroups(c *gin.Context) {
	Log.Info("API USER", "get groups start")

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

	var request model.UserGetGroupsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handleUserGetGroups(c, request.BigGroup)

	Log.Info("API USER", "get groups end")
}

func ApiUserAddGroupGrade(c *gin.Context) {
	Log.Info("API USER", "add group grade start")

	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "missing authorization header",
		})
		return
	}
	valid, studentId, err := util.VerifyJwtToken(jwtToken)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	var request model.UserAddGroupGradeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handleUserAddGroupGrade(c, studentId, request)

	Log.Info("API USER", "add group grade end")
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
	jwtToken, err = util.GenerateJwtToken(request.StudentId, time.Now().Unix(), time.Now().Unix()+UNAUTHORIZED_JWT_VALID_TIME, totp)
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

func handleUserTOTPInitBegin(c *gin.Context, studentId string) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "CE Final Project",
		AccountName: studentId,
	})
	if err != nil {
		Log.Error("API USER", "failed to generate TOTP: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate TOTP",
		})
		return
	}

	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		Log.Error("API USER", "failed to generate QR code: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate QR code",
		})
		return
	}

	if err := png.Encode(&buf, img); err != nil {
		Log.Error("API USER", "failed to encode image: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to encode image",
		})
		return
	}

	userAccount := model.UserAccount{
		StudentId: studentId,
		Totp:      false,
		SecretKey: key.Secret(),
	}
	if tgtUserAccount, err := GetUserAccount(studentId); err != nil {
		Log.Error("API USER", "failed to get user account: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else if tgtUserAccount != nil {
		if err = UpdateUserAccount(&userAccount); err != nil {
			Log.Error("API USER", "failed to update user account: "+err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err = InsertUserAccount(&userAccount); err != nil {
			Log.Error("API USER", "failed to insert user account: "+err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// filePath := "./qr_code.png"
	// file, err := os.Create(filePath)
	// if err != nil {
	// 	Log.Error("API USER", "failed to create file: "+err.Error())
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "failed to create file",
	// 	})
	// 	return
	// }
	// defer file.Close()

	// if _, err := file.Write(buf.Bytes()); err != nil {
	// 	Log.Error("API USER", "failed to write file: "+err.Error())
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "failed to write file",
	// 	})
	// 	return
	// }

	qrCode := base64.StdEncoding.EncodeToString(buf.Bytes())

	response := model.UserTOTPInitBeginResponse{
		QrCode:    qrCode,
		SecretKey: key.Secret(),
	}

	c.JSON(http.StatusOK, response)
}

func handleUserTOTPInitFinish(c *gin.Context, studentId string, code string) {
	userAccount, err := GetUserAccount(studentId)
	if err != nil {
		Log.Error("API USER", "failed to get user account: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userAccount == nil {
		Log.Error("API USER", "user account not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "user account not found"})
		return
	}

	valid, err := totp.ValidateCustom(code, userAccount.SecretKey, time.Now(), totp.ValidateOpts{
		Period: 30,
		Digits: 6,
	})
	if err != nil {
		Log.Error("API USER", "failed to validate code: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !valid {
		Log.Error("API USER", "invalid code")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid code"})
		return
	}

	userAccount.Totp = true
	err = UpdateUserAccount(userAccount)
	if err != nil {
		Log.Error("API USER", "failed to update user account: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "TOTP initialized"})
}

func handleUserTOTP(c *gin.Context, studentId string, code string) {
	userAccount, err := GetUserAccount(studentId)
	if err != nil {
		Log.Error("API USER", "failed to get user account: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userAccount == nil {
		Log.Error("API USER", "user account not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "user account not found"})
		return
	}

	valid, err := totp.ValidateCustom(code, userAccount.SecretKey, time.Now(), totp.ValidateOpts{
		Period: 30,
		Digits: 6,
	})
	if err != nil {
		Log.Error("API USER", "failed to validate code: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !valid {
		Log.Error("API USER", "invalid code")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid code"})
		return
	}

	jwtToken, err := util.GenerateJwtToken(studentId, time.Now().Unix(), time.Now().Unix()+AUTHORIZED_JWT_VALID_TIME, true)
	if err != nil {
		Log.Error("API USER", "failed to generate jwt token: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := model.UserTOTPResponse{
		JwtToken: jwtToken,
	}

	c.JSON(http.StatusOK, response)
}

func handleUserGetBigGroups(c *gin.Context) {
	bigGroups, err := GetBigGroups()
	if err != nil {
		Log.Error("API USER", "failed to get big groups: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := model.UserGetBigGroupsResponse{
		BigGroups: make([]string, len(bigGroups)),
	}
	for i, bigGroup := range bigGroups {
		response.BigGroups[i] = bigGroup.BigGroup
	}

	c.JSON(http.StatusOK, response)
}

func handleUserGetGroups(c *gin.Context, bigGroup string) {
	groups, err := GetGroups(bigGroup)
	if err != nil {
		Log.Error("API USER", "failed to get groups: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := model.UserGetGroupsResponse{
		Groups: make([]model.Group, len(groups)),
	}
	for i, group := range groups {
		response.Groups[i] = *group
	}

	c.JSON(http.StatusOK, response)
}

func handleUserAddGroupGrade(c *gin.Context, studentId string, request model.UserAddGroupGradeRequest) {
	groupGradeCommentItem := &model.GroupGradeCommentItem{
		StudentId: studentId,
		Grade:     request.Grade,
		Comment:   request.Comment,
	}

	group, err := GetGroup(request.BigGroup, request.GroupId)
	if err != nil {
		Log.Error("API USER", "failed to get group: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if group.GradedStudentId[studentId] {
		Log.Error("API USER", "already graded")
		c.JSON(http.StatusBadRequest, gin.H{"error": "already graded"})
		return
	}

	if err := UpdateGroupGradeComment(request.BigGroup, request.GroupId, groupGradeCommentItem); err != nil {
		Log.Error("API USER", "failed to update group grade comment: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := UpdateGroupGradedStudentId(request.BigGroup, request.GroupId, studentId); err != nil {
		Log.Error("API USER", "failed to update group graded student id: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "grade added successfully"})
}
