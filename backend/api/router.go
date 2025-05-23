package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	addMiddleware(router)
	addRoutes(router)
	return router
}

func addMiddleware(router *gin.Engine) {
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
}

func addRoutes(router *gin.Engine) {
	group := router.Group("/api")
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			group.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			group.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			group.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			group.DELETE(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			group.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodOptions:
			group.OPTIONS(route.Pattern, route.HandlerFunc)
		}
	}
}

var routes = []route{
	{
		Name:        "Admin Login",
		Method:      http.MethodPost,
		Pattern:     "/admin/login",
		HandlerFunc: ApiAdminLogin,
	},
	{
		Name:        "Admin Get Big Groups",
		Method:      http.MethodGet,
		Pattern:     "/admin/group/big",
		HandlerFunc: ApiAdminGetBigGroups,
	},
	{
		Name:        "Admin Add Big Group",
		Method:      http.MethodPost,
		Pattern:     "/admin/group/big",
		HandlerFunc: ApiAdminAddBigGroup,
	},
	{
		Name:        "Admin Delete Big Group",
		Method:      http.MethodDelete,
		Pattern:     "/admin/group/big",
		HandlerFunc: ApiAdminDeleteBigGroup,
	},
	{
		Name:        "Admin Get Groups",
		Method:      http.MethodGet,
		Pattern:     "/admin/group",
		HandlerFunc: ApiAdminGetGroups,
	},
	{
		Name:        "Admin Add Group",
		Method:      http.MethodPost,
		Pattern:     "/admin/group",
		HandlerFunc: ApiAdminAddGroup,
	},
	{
		Name:        "Admin Delete Group",
		Method:      http.MethodDelete,
		Pattern:     "/admin/group",
		HandlerFunc: ApiAdminDeleteGroup,
	},
	{
		Name:        "Admin Get Group Grades",
		Method:      http.MethodGet,
		Pattern:     "/admin/group/grade",
		HandlerFunc: ApiAdminGetGroupGrades,
	},
	{
		Name:        "User Login",
		Method:      http.MethodPost,
		Pattern:     "/user/login",
		HandlerFunc: ApiUserLogin,
	},
	{
		Name:        "User TOTP Init Begin",
		Method:      http.MethodGet,
		Pattern:     "/user/totp/init/begin",
		HandlerFunc: ApiUserTOTPInitBegin,
	},
	{
		Name:        "User TOTP Init Finish",
		Method:      http.MethodPost,
		Pattern:     "/user/totp/init/finish",
		HandlerFunc: ApiUserTOTPInitFinish,
	},
	{
		Name:        "User TOTP",
		Method:      http.MethodPost,
		Pattern:     "/user/totp",
		HandlerFunc: ApiUserTOTP,
	},
	{
		Name:        "User Get Big Groups",
		Method:      http.MethodGet,
		Pattern:     "/user/group/big",
		HandlerFunc: ApiUserGetBigGroups,
	},
	{
		Name:        "User Get Groups",
		Method:      http.MethodGet,
		Pattern:     "/user/group",
		HandlerFunc: ApiUserGetGroups,
	},
	{
		Name:        "User Add Group Grade",
		Method:      http.MethodPost,
		Pattern:     "/user/group/grade",
		HandlerFunc: ApiUserAddGroupGrade,
	},
}
