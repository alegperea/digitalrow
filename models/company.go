package models

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"

	"github.com/agp/push-notifications-api/entity"
	"github.com/agp/push-notifications-api/services"
	"github.com/guregu/dynamo"
)

//GetAllcompanys Fetch all company data
func GetAllCompanies(company *[]entity.Company) (err error) {
	return nil

}

//Createcompany ... Insert New data
func CreateCompany(sess client.ConfigProvider, company *entity.Company) (err error) {

	company.UUID = services.GenerateUUID()
	company.Password, _ = services.PasswordEncrypt(company.Password)
	company.UserData = "Profile"
	currentTime := time.Now()

	company.CreatedAt = currentTime.Format("2017-09-07 17:06:06")
	company.UpdatedAt = currentTime.Format("2017-09-07 17:06:06")

	db := dynamo.New(sess, &aws.Config{Region: aws.String("us-east-1")})
	table := db.Table(os.Getenv("dynamodb_table"))

	fmt.Println("company: ", company)

	// put item
	errPut := table.Put(company).Run()
	if errPut != nil {
		fmt.Println("Error Put: ", errPut)
		return errPut
	}
	// get the same item
	var result entity.Company

	errGet := table.Get("UUID", company.UUID).Range("UserData", dynamo.Equal, company.UserData).One(&result)
	if errGet != nil {
		fmt.Println("Error Get: ", errGet)
		return err
	}
	return nil

}

//GetcompanyByID ... Fetch only one company by Id
func GetCompanyByID(sess client.ConfigProvider, company *entity.Company) (err error) {

	return nil
}

//GetcompanyByEmail ... Fetch only one company by Id
func GetCompanyByEmail(sess client.ConfigProvider, email string) (company entity.Company, err error) {

	db := dynamo.New(sess, &aws.Config{Region: aws.String("us-east-1")})
	table := db.Table(os.Getenv("dynamodb_table"))

	fmt.Println("email ", email)

	errGet := table.Get("Email", email).Index("Email-Index").One(&company)
	if errGet != nil {
		fmt.Println("Error Get: ", errGet)
		return company, err
	}
	fmt.Println("company Get ", company)

	return company, nil

}

//Updatecompany ... Update company
func UpdateCompany(company *entity.Company, id string) (err error) {
	return nil

}

//Deletecompany ... Delete company
func DeleteCompany(company *entity.Company, id string) (err error) {
	return nil

}
