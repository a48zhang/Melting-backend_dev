package model

type QNconfig struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket_name"`
	Domain    string `json:"domain_name"`
}

type LoginRequest struct {
	NickName string `json:"nick_name"`
	QQ       string `json:"qq"`
	Auth     string `json:"auth"`
}
