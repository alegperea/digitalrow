package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	mail "github.com/xhit/go-simple-mail/v2"
)

var err error

type Company struct {
	UUID         string `json:"UUID" index:"UUID-index,hash" swaggerignore:"true"`
	UserData     string `json:"UserData" dynamo:",range" swaggerignore:"true"`
	CUIT         string `json:"CUIT"`
	Email        string `json:"Email" index:"Email-index,hash"`
	Password     string `json:"Password"`
	CompanyName  string `json:"CompanyName"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Phone        string `json:"Phone"`
	Logo         string `json:"Logo"`
	Subscription string `json:"Subscription"`
	CreatedAt    string `json:"CreatedAt" swaggerignore:"true"` // Range key, a.k.a. sort key
	UpdatedAt    string `json:"UpdatedAt" swaggerignore:"true"` // Range key, a.k.a. sort key
}

func handle(ctx context.Context, sqsEvent events.SQSEvent) error {
	// execute the lambda

	fmt.Println("sqsEvent", sqsEvent)
	fmt.Println("sqsEvent", sqsEvent)

	for _, message := range sqsEvent.Records {
		sendMail(message.Body)
		fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)
	}

	return nil

}

func main() {
	if err != nil {
		fmt.Println("Status:", err)
	}

	err := godotenv.Load()
	if err != nil {

		log.Fatal("Error loading .env file")
	}

	lambda.Start(handle)
}

func sendMail(message string) {

	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = os.Getenv("mailserver_user")
	server.Password = os.Getenv("mailserver_password")
	server.Encryption = mail.EncryptionSTARTTLS

	// Since v2.3.0 you can specified authentication type:
	// - PLAIN (default)
	// - LOGIN
	// - CRAM-MD5
	// - None
	// server.Authentication = mail.AuthPlain

	// Variable to keep alive connection
	server.KeepAlive = false

	// Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second

	// Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second

	// Set TLSConfig to provide custom TLS configuration. For example,
	// to skip TLS verification (useful for testing):
	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// SMTP client
	smtpClient, err := server.Connect()

	if err != nil {
		log.Fatal(err)
	}

	company := Company{}
	json.Unmarshal([]byte(message), &company)

	email := mail.NewMSG()
	email.SetFrom("agp Fila Digital <filadigital@agp.com>").
		AddTo(company.Email).
		SetSubject("Bienvenido a Fila Digital")

	htmlBody := `<html>
		<head>
			<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
			<title>Bienvenido a Fila Digital</title>
		</head>
		<body>
			<p>La cuenta de ` + company.FirstName + ` ` + company.LastName + ` ha sido creada con éxito </p>
		</body>
	</html>`

	email.SetBody(mail.TextHTML, htmlBody)
	email.SetSender(company.Email)

	// also you can add body from []byte with SetBodyData, example:
	// email.SetBodyData(mail.TextHTML, []byte(htmlBody))
	// or alternative part
	// email.AddAlternativeData(mail.TextHTML, []byte(htmlBody))

	// add inline
	//email.Attach(&mail.File{FilePath: "/path/to/image.png", Name: "Gopher.png", Inline: true})

	// always check error after send
	if email.Error != nil {
		log.Fatal(email.Error)
	}

	// Call Send and pass the client
	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent")
	}
}
