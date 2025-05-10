package model

type GroupGradeComment struct {
	BigGroup              string                  `bson:"big_group"`
	GroupId               int                     `bson:"group_id"`
	GroupGradeCommentList []GroupGradeCommentItem `bson:"group_grade_comment_list"`
}

type GroupGradeCommentItem struct {
	StudentId string `bson:"student_id"`
	Grade     int    `bson:"grade"`
	Comment   string `bson:"comment"`
}
