package api

import (
	"ce/backend/model"
	"ce/backend/mongo"
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

	data, err := bson.Marshal(result)
	if err != nil {
		return nil, err
	}

	var userAccount model.UserAccount
	err = bson.Unmarshal(data, &userAccount)
	if err != nil {
		return nil, err
	}

	return &userAccount, nil
}

func InsertUserAccount(userAccount *model.UserAccount) error {
	return mongo.InsertOne(mongo.USER_ACCOUNT_COLLECTION, userAccount)
}

func UpdateUserAccount(userAccount *model.UserAccount) error {
	return mongo.UpdateOne(mongo.USER_ACCOUNT_COLLECTION, bson.M{"student_id": userAccount.StudentId}, bson.M{"$set": userAccount})
}
