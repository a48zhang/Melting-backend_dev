package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"main/model/db"
	"os"
	"time"
)

func InitService() {
	db.OpenDB()
	initQN()
	loadCert()
	loadAddr()
}

var TLSCert = "conf/chain.crt"
var TLSKey = "conf/key.key"
var ServerAddr = ":65000"

func loadAddr() {
	if addr := os.Getenv("MELT_ADDR"); addr != "" {
		ServerAddr = addr
	}
}

func loadCert() {
	if cert := os.Getenv("MELT_CERT"); cert != "" {
		TLSCert = cert
	}
	if key := os.Getenv("MELT_KEY"); key != "" {
		TLSKey = key
	}
}

func initQN() {
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

func Logger() {
	y, m, d := time.Now().Date()
	target := fmt.Sprintf("./log/%v_%v_%v_%v.log", y, m, d, time.Now().Nanosecond())
	f, _ := os.Create(target)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
