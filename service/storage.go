package service

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

var conf storageConfig

func UploadProfilePhoto(file *multipart.File, size int64) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope:        conf.Bucket,
		SaveKey:      "melting/${year}_${mon}_${day}_${hour}_${min}_${sec}.jpg",
		ForceSaveKey: true,
	}

	mac := qbox.NewMac(conf.AccessKey, conf.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	formUploader := storage.NewFormUploader(&storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	})

	putExtra := new(storage.PutExtra)
	ret := new(storage.PutRet)

	if err := formUploader.Put(context.Background(), ret,
		upToken, "", *file, size, putExtra); err != nil {
		return "", err
	}

	return conf.Domain + "/" + ret.Key, nil
}

type storageConfig struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket_name"`
	Domain    string `json:"domain_name"`
}
