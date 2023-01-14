package model

import "main/model/db"

type QNconfig struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket_name"`
	Domain    string `json:"domain_name"`
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
