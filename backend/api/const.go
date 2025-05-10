package api

import (
	"ce/backend/logger"
)

var Log *logger.Logger

const (
	UNAUTHORIZED_JWT_VALID_TIME = 60 * 5 // 5 minutes
	AUTHORIZED_JWT_VALID_TIME  = 60 * 60 * 5 // 5 hours
)
