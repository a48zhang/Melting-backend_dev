package model

import "main/model/db"

type QNconfig struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Domain    string
}

type sth interface {
	db.User | db.Template | db.ProposalInfo | db.Tag | db.Question | db.Game
	TableName() string
	GetKey() (string, int)
}

type LoginRequest struct {
	NickName string `json:"nick_name"`
	QQ       string `json:"qq"`
	Auth     string `json:"auth"`
}
