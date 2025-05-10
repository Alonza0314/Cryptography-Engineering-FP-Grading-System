package model

type GroupGradeComment struct {
	BigGroup              string                  `bson:"big_group"`
	GroupId               int                     `bson:"group_id"`
	GroupGradeCommentList GroupGradeCommentList   `bson:"group_grade_comment_list"`
}

type GroupGradeCommentList []GroupGradeCommentItem

type GroupGradeCommentItem struct {
	StudentId string `bson:"student_id" json:"student_id"`
	Grade     int    `bson:"grade" json:"grade"`
	Comment   string `bson:"comment" json:"comment"`
}
