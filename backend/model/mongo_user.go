package model

type UserAccount struct {
	StudentId string `bson:"student_id"`
	Totp      bool   `bson:"totp"`
	SecretKey string `bson:"secret_key"`
}
