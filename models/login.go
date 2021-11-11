package models

import (
	"fmt"

	"github.com/agp/push-notifications-api/entity"
	"github.com/aws/aws-sdk-go/aws/client"
	"golang.org/x/crypto/bcrypt"
)

//ValidateUser validate the username and password is valid or not
func ValidateCompany(sess client.ConfigProvider, loginVals entity.Login) (*entity.Company, bool) {
	email := loginVals.Email
	company, errFind := GetCompanyByEmail(sess, email)

	fmt.Println("Password ", string(company.Password))

	if errFind != nil {
		fmt.Println("UserNotExist")
		return &company, false
	}

	passwordBytes := []byte(loginVals.Password)
	passwordBD := []byte(company.Password)

	fmt.Println("passwordBytes:", passwordBytes)
	fmt.Println("passwordBD:", passwordBD)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		fmt.Println("err:", err.Error())
		return &company, false
	}

	return &company, true

}
