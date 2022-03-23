package model

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"go/server"
)

type Admin struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	PrivateKey string `json:"private_key"`
	Password string `json:"password"`
}

func adminPasswordEncrypt(password string, privateKey string) string {
	passrand := password + privateKey
	shaHash := sha256.New224()

	passEncryByte := shaHash.Sum([]byte(passrand))

	return hex.EncodeToString(passEncryByte)
}

func (admin *Admin) IsAdmin(password string) error {
	needSignPassword := adminPasswordEncrypt(password, admin.PrivateKey)

	if needSignPassword != admin.Password {
		return errors.New("password is error")
	}

	return nil
}

func (admin *Admin) GetAdmin(phione string) error {
	mysql := server.GetMysql()

	mysql.GormDB.Where("phione = ?", phione).Find(admin)

	server.PutMysql(mysql)

	if admin.Name == "" {
		return errors.New("no admin")
	}

	return nil
}

