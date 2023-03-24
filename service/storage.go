package service

import (
	"context"
	"encoding/json"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"log"
	"main/model/db"
	"mime/multipart"
	"os"
)

var conf storageConfig

func Init() {
	db.OpenDB()
	file, err := os.Open("./conf/qn.json")
	if err != nil {
		conf = storageConfig{
			AccessKey: os.Getenv("access_key"),
			SecretKey: os.Getenv("secret_key"),
			Bucket:    os.Getenv("bucket_name"),
			Domain:    os.Getenv("domain_name"),
		}
		if conf.SecretKey == "" {
			log.Fatal("Failed to connect to cloud storage. Check the env settings")
		}
		return
	} else {
		tmp, _ := io.ReadAll(file)
		err = json.Unmarshal(tmp, &conf)
	}
	if err != nil {
		log.Fatal("Failed to connect to cloud storage. Error:" + err.Error())
	}
}

func UploadProfilePhoto(file *multipart.File, size int64) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope:        conf.Bucket,
		SaveKey:      "${year}_${mon}_${day}_${hour}_${min}_${sec}.jpg",
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
