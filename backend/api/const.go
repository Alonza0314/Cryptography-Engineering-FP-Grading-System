package api

import (
	"ce/backend/logger"
	"ce/backend/mongo"
)

var (
	Log *logger.Logger
	Db  *mongo.Mongo
)

const (
	UNAUTHORIZED_JWT_VALID_TIME = 60 * 5      // 5 minutes
	AUTHORIZED_JWT_VALID_TIME   = 60 * 60 * 5 // 5 hours
)
