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

//GetAllcustomers Fetch all customer data
func GetAllCustomers(customer *[]entity.Customer) (err error) {
	return nil

}

//Createcustomer ... Insert New data
func AddCustomer(sess client.ConfigProvider, uuID string, commerce *entity.Commerce, customer *entity.Customer) (err error) {

	fmt.Println("customer: ", customer)

	fmt.Println("uuID: ", uuID)

	fmt.Println("commerce: ", commerce)

	currentTime := time.Now()

	customer.CustomerID = "customer-" + services.GenerateUUID()
	customer.CreatedAt = currentTime.Format("2017-09-07 17:06:06")

	commerce.UUID = uuID

	commerce.Customers = append(commerce.Customers, *customer)

	db := dynamo.New(sess, &aws.Config{Region: aws.String("us-east-1")})
	table := db.Table(os.Getenv("dynamodb_table"))

	fmt.Println("customer: ", customer)

	// put item
	errPut := table.Put(commerce).Run()
	if errPut != nil {
		fmt.Println("Error Put: ", errPut)
		return errPut
	}
	return nil

}

//GetcustomerByID ... Fetch only one customer by Id
func GetCustomerByCommerce(sess client.ConfigProvider, uuID, commerceID string) ([]entity.Customer, error) {

	db := dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("region"))})
	table := db.Table(os.Getenv("dynamodb_table"))

	var commerce entity.Commerce

	fmt.Println("CommerceID ", commerceID)
	fmt.Println("UserID ", uuID)

	errGet := table.Get("UUID", uuID).Range("UserData", dynamo.Equal, commerceID).One(&commerce)
	if errGet != nil {
		fmt.Println("Error Get: ", errGet)
		return commerce.Customers, errGet
	}

	fmt.Println("Results ", commerce)
	return commerce.Customers, nil
}

//GetcustomerByID ... Fetch only one customer by Id
func GetCustomerByID(sess client.ConfigProvider, customerID string, customer *entity.Customer) (err error) {

	return nil
}

//GetcustomerByEmail ... Fetch only one customer by Id
func GetCustomerByEmail(sess client.ConfigProvider, email string) (customer entity.Customer, err error) {

	db := dynamo.New(sess, &aws.Config{Region: aws.String("us-east-1")})
	table := db.Table(os.Getenv("dynamodb_table"))

	fmt.Println("email ", email)

	errGet := table.Get("Email", email).Index("Email-Index").One(&customer)
	if errGet != nil {
		fmt.Println("Error Get: ", errGet)
		return customer, err
	}
	fmt.Println("customer Get ", customer)

	return customer, nil

}

//Updatecustomer ... Update customer
func UpdateCustomer(customer *entity.Customer, id string) (err error) {
	return nil

}

//Deletecustomer ... Delete customer
func DeleteCustomer(customer *entity.Customer, id string) (err error) {
	return nil

}
