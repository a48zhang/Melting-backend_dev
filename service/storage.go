package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/model"
	"main/model/db"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var conf model.QNconfig

func Init() {
	db.OpenDB()
	file, err := os.Open("./conf/qn.json")
	if err != nil {
		conf = model.QNconfig{
			AccessKey: os.Getenv("access_key"),
			SecretKey: os.Getenv("secret_key"),
			Bucket:    os.Getenv("bucket_name"),
			Domain:    os.Getenv("domain_name"),
		}
	} else {
		tmp, _ := io.ReadAll(file)
		json.Unmarshal(tmp, &conf)
	}
	if conf.AccessKey == "" {
		log.Fatal("Failed to connect to cloud storage")
	}
}

func UploadProfilePhoto(id int, file *multipart.File, size int64) (string, error) {
	keyToOverwrite := strconv.Itoa(id) + ".jpg"
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", conf.Bucket, keyToOverwrite),
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
		upToken, keyToOverwrite, *file, size, putExtra); err != nil {
		return "", err
	}
	return conf.Domain + "/" + ret.Key, nil
}
