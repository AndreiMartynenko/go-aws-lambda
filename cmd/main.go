/*
------> API GATEAWAY -------> LAMBDA ------> DYNAMODB
*/

package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	//creating a session
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)})
	//handling error
	if err != nil {
		return
	}
	//DynamoDB
	dynaClient = dynamodb.New(awsSession)
	//Lambda package
	lambda.Start(handler)

}
