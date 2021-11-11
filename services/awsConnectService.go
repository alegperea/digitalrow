package services

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var AccessKeyID string
var SecretAccessKey string
var MyRegion string

func ConnectAws() *session.Session {
	AccessKeyID = os.Getenv("aws_access_key_id")
	SecretAccessKey = os.Getenv("aws_secret_access_key")
	MyRegion = "us-east-1"
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}
