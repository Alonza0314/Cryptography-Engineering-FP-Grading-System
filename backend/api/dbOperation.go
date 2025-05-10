package api

import (
	"ce/backend/model"
	"ce/backend/mongo"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetUserAccount(studentId string) (*model.UserAccount, error) {
	filter := bson.M{"student_id": studentId}
	result, err := Db.GetOne(mongo.USER_ACCOUNT_COLLECTION, filter)
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
	return Db.InsertOne(mongo.USER_ACCOUNT_COLLECTION, userAccount)
}

func UpdateUserAccount(userAccount *model.UserAccount) error {
	return Db.UpdateOne(mongo.USER_ACCOUNT_COLLECTION, bson.M{"student_id": userAccount.StudentId}, bson.M{"$set": userAccount})
}

func GetBigGroups() ([]*model.BigGroup, error) {
	filter := bson.M{}
	cursor, err := Db.GetAll(mongo.BIG_GROUP_COLLECTION, filter)
	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			return nil, nil
		}
		return nil, err
	}

	var bigGroups []*model.BigGroup
	for _, result := range cursor {
		data, err := bson.Marshal(result)
		if err != nil {
			return nil, err
		}

		var bigGroup model.BigGroup
		err = bson.Unmarshal(data, &bigGroup)
		if err != nil {
			return nil, err
		}
		bigGroups = append(bigGroups, &bigGroup)
	}

	return bigGroups, nil
}

func GetBigGroup(bigGroupName string) (*model.BigGroup, error) {
	filter := bson.M{"big_group": bigGroupName}
	result, err := Db.GetOne(mongo.BIG_GROUP_COLLECTION, filter)
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

	var bigGroup model.BigGroup
	err = bson.Unmarshal(data, &bigGroup)
	if err != nil {
		return nil, err
	}

	return &bigGroup, nil
}

func InsertBigGroup(bigGroupName string) error {
	return Db.InsertOne(mongo.BIG_GROUP_COLLECTION, model.BigGroup{BigGroup: bigGroupName})
}

func DeleteBigGroup(bigGroupName string) error {
	return Db.DeleteOne(mongo.BIG_GROUP_COLLECTION, bson.M{"big_group": bigGroupName})
}

func GetGroups(bigGroupName string) ([]*model.Group, error) {
	filter := bson.M{"big_group": bigGroupName}
	cursor, err := Db.GetAll(mongo.GROUP_COLLECTION, filter)
	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			return nil, nil
		}
		return nil, err
	}

	var groups []*model.Group
	for _, result := range cursor {
		data, err := bson.Marshal(result)
		if err != nil {
			return nil, err
		}

		var group model.Group
		err = bson.Unmarshal(data, &group)
		if err != nil {
			return nil, err
		}
		groups = append(groups, &group)
	}

	return groups, nil
}

func GetGroup(bigGroupName string, groupId int) (*model.Group, error) {
	filter := bson.M{"big_group": bigGroupName, "group_id": groupId}
	result, err := Db.GetOne(mongo.GROUP_COLLECTION, filter)
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

	var group model.Group
	err = bson.Unmarshal(data, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func UpdateGroupGradedStudentId(bigGroupName string, groupId int, studentId string) error {
	return Db.UpdateOne(
		mongo.GROUP_COLLECTION,
		bson.M{"big_group": bigGroupName, "group_id": groupId},
		bson.M{"$set": bson.M{
			"graded_student_id." + studentId: true,
		}},
	)
}

func InsertGroup(group *model.Group) error {
	return Db.InsertOne(mongo.GROUP_COLLECTION, group)
}

func DeleteGroup(bigGroupName string, groupId int) error {
	return Db.DeleteOne(mongo.GROUP_COLLECTION, bson.M{"big_group": bigGroupName, "group_id": groupId})
}

func GetGroupGradeComment(bigGroupName string, groupId int) (*model.GroupGradeComment, error) {
	filter := bson.M{"big_group": bigGroupName, "group_id": groupId}
	result, err := Db.GetOne(mongo.GROUP_GRADE_COMMENT_COLLECTION, filter)
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

	var groupGradeComment model.GroupGradeComment
	err = bson.Unmarshal(data, &groupGradeComment)
	if err != nil {
		return nil, err
	}

	return &groupGradeComment, nil
}

func InsertGroupGradeComment(groupGradeComment *model.GroupGradeComment) error {
	return Db.InsertOne(mongo.GROUP_GRADE_COMMENT_COLLECTION, groupGradeComment)
}

func UpdateGroupGradeComment(bigGroupName string, groupId int, commentItem *model.GroupGradeCommentItem) error {
	return Db.UpdateOne(
		mongo.GROUP_GRADE_COMMENT_COLLECTION,
		bson.M{"big_group": bigGroupName, "group_id": groupId},
		bson.M{"$push": bson.M{"group_grade_comment_list": commentItem}},
	)
}

func DeleteGroupGradeComment(bigGroupName string, groupId int) error {
	return Db.DeleteOne(mongo.GROUP_GRADE_COMMENT_COLLECTION, bson.M{"big_group": bigGroupName, "group_id": groupId})
}
