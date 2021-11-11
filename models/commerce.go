package models

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"

	"github.com/agp/push-notifications-api/entity"
	"github.com/agp/push-notifications-api/services"
	"github.com/guregu/dynamo"
)

//GetAllcommerces Fetch all commerce data
func GetAllCommerces(sess client.ConfigProvider, uuID string) ([]entity.Commerce, error) {

	db := dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("region"))})
	table := db.Table(os.Getenv("dynamodb_table"))
	var commerces []entity.Commerce
	errGet := table.Get("UUID", uuID).Range("UserData", dynamo.BeginsWith, "commerce-").All(dynamo.AWSEncoding(&commerces))
	if errGet != nil {
		fmt.Println("Error Get: ", errGet)
		return commerces, errGet
	}

	fmt.Println("Results ", commerces)
	return commerces, nil

}

//Createcommerce ... Insert New data
func CreateCommerce(sess client.ConfigProvider, companyName string, commerce *entity.Commerce, uuID string) (err error) {

	commerce.UUID = uuID

	commerce.UserData = "commerce-" + services.GenerateUUID()
	companyNameLower := strings.ToLower(companyName)

	commerce.QRCode = "http://127.0.0.1/" + companyNameLower + "/" + commerce.UserData

	currentTime := time.Now()

	commerce.CreatedAt = currentTime.Format("2017-09-07 17:06:06")
	commerce.UpdatedAt = currentTime.Format("2017-09-07 17:06:06")

	db := dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("region"))})
	table := db.Table(os.Getenv("dynamodb_table"))

	fmt.Println("commerce: ", commerce)

	// put item
	errPut := table.Put(commerce).Run()
	if errPut != nil {
		fmt.Println("Error Put: ", errPut)
		return errPut
	}

	return nil

}

//GetcommerceByID ... Fetch only one commerce by Id
func GetCommerceByID(sess client.ConfigProvider, uuID, id string) (entity.Commerce, error) {

	db := dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("region"))})
	table := db.Table(os.Getenv("dynamodb_table"))

	var commerce entity.Commerce

	fmt.Println("ID ", id)
	fmt.Println("UserID ", uuID)

	errGet := table.Get("UUID", uuID).Range("UserData", dynamo.Equal, id).One(&commerce)
	if errGet != nil {
		fmt.Println("Error Get: ", errGet)
		return commerce, errGet
	}

	fmt.Println("Results ", commerce)
	return commerce, nil
}

//GetcommerceByEmail ... Fetch only one commerce by Id
func GetCommerceByEmail(commerce *entity.Commerce, email string) (err error) {
	return nil
}

//Updatecommerce ... Update commerce
func UpdateCommerce(sess client.ConfigProvider, commerce *entity.Commerce, uuID, id string) (err error) {
	commerce.UUID = uuID

	commerce.UserData = id

	currentTime := time.Now()

	commerce.UpdatedAt = currentTime.Format("2017-09-07 17:06:06")

	db := dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("region"))})
	table := db.Table(os.Getenv("dynamodb_table"))

	fmt.Println("commerce: ", commerce)

	// Putting an item
	errPut := table.Put(dynamo.AWSEncoding(&commerce)).Run()

	if errPut != nil {
		fmt.Println("Error Put: ", errPut)
		return errPut
	}

	return nil

}

//Deletecommerce ... Delete commerce
func DeleteCommerce(sess client.ConfigProvider, commerce *entity.Commerce, uuID, id string) (err error) {
	commerce.UUID = uuID

	commerce.UserData = id

	currentTime := time.Now()

	commerce.CreatedAt = currentTime.Format("2017-09-07 17:06:06")
	commerce.UpdatedAt = currentTime.Format("2017-09-07 17:06:06")

	db := dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("region"))})
	table := db.Table(os.Getenv("dynamodb_table"))

	fmt.Println("commerce: ", commerce)
	fmt.Println("commerceID: ", id)
	fmt.Println("UUID: ", uuID)

	// Putting an item
	errDel := table.Delete("UUID", uuID).Range("UserData", id).Run()

	if errDel != nil {
		fmt.Println("Error Put: ", errDel)
		return errDel
	}

	return nil

}
