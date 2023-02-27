package service

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

func B64Encode(data string) []byte {
	ret := make([]byte, 0)
	base64.StdEncoding.Encode(ret, []byte(data))
	return ret
}
func B64Decode(data string) ([]byte, error) {
	ret, err := base64.StdEncoding.DecodeString(data)
	return ret, err
}

func Encrypter(Pwd string) (encryptPwd string, err error) {
	bytePwd := []byte(Pwd)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.DefaultCost)
	if err != nil {
		return
	}
	encryptPwd = string(hash)
	return
}

func Validator(encryptPwd string, plainPwd string) bool {
	hashByte := []byte(encryptPwd)
	plainByte := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(hashByte, plainByte)
	if err != nil {
		return false
	}
	return true
}
