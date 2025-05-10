package model

type Group struct {
	BigGroup        string   `bson:"big_group" json:"big_group" binding:"required"`
	GroupId         int      `bson:"group_id" json:"group_id" binding:"required"`
	GroupName       string   `bson:"group_name" json:"group_name" binding:"required"`
	LeaderName      string   `bson:"leader_name" json:"leader_name" binding:"required"`
	LeaderStudentId string   `bson:"leader_student_id" json:"leader_student_id" binding:"required"`
	Members         []Member `bson:"members" json:"members" binding:"required"`
	GradedStudentId map[string]bool `bson:"graded_student_id" json:"graded_student_id"`
}

type Member struct {
	MemberName      string `bson:"member_name" json:"member_name" binding:"required"`
	MemberStudentId string `bson:"member_student_id" json:"member_student_id" binding:"required"`
}
