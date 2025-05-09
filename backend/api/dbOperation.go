package api

import (
	"ce/backend/model"
	"ce/backend/mongo"
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetUserAccount(studentId string) (*model.UserAccount, error) {
	filter := bson.M{"student_id": studentId}
	result, err := mongo.GetOne(mongo.USER_ACCOUNT_COLLECTION, filter)

	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			return nil, nil
		}
		return nil, err
	}

	userAccount, ok := result.(*model.UserAccount)
	if !ok {
		return nil, errors.New("failed to convert result to UserAccount")
	}

	return userAccount, nil
}
