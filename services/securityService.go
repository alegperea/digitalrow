package services

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func PasswordEncrypt(password string) (string, error) {
	lengh := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), lengh)
	return string(bytes), err
}

func GenerateUUID() string {

	uuidWithHyphen := uuid.New()
	fmt.Println(uuidWithHyphen)
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
	return uuid
}
