/*
------> API GATEAWAY -------> LAMBDA ------> DYNAMODB
*/

package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	//creating a session
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession

}
