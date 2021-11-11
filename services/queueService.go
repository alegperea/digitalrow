package services

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func getQueueUrl(queueName string) (*sqs.GetQueueUrlOutput, error) {
	sess := ConnectAws()

	svc := sqs.New(sess)

	inputQUrl := &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	}

	resultQUrl, errQUrl := svc.GetQueueUrl(inputQUrl)
	if errQUrl != nil {
		return nil, errQUrl
	}

	return resultQUrl, nil

}

func SetQueue(body, queueName string) (*sqs.SendMessageOutput, error) {
	sess := ConnectAws()
	svc := sqs.New(sess)

	queueUrl, errQUrl := getQueueUrl(queueName)

	if errQUrl != nil {
		return nil, errQUrl
	}

	fmt.Println("QueueUrl", *queueUrl.QueueUrl)

	inputQueue := &sqs.SendMessageInput{
		MessageBody: aws.String(body),
		QueueUrl:    aws.String(*queueUrl.QueueUrl),
	}
	resultQueue, errQueue := svc.SendMessage(inputQueue)

	if errQueue != nil {
		return nil, errQueue
	}

	return resultQueue, nil

}
