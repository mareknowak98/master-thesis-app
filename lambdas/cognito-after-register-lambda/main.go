package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"mylearnproject/lambdas/cognito-after-register-lambda/cmd"
)

func HandleRequest(ctx context.Context, request cmd.CognitoEventRequest) (cmd.CognitoEventResponse, error) {
	//Initialize response object
	res := cmd.CognitoEventResponse{}
	res.Version = 1
	res.TriggerSource = request.TriggerSource
	res.Region = request.Region
	res.UserPoolID = request.UserPoolID
	res.Username = request.Username

	// Initialize AWS client
	c := cmd.NewClient("eu-central-1")
	tableName := "cognito-users"

	// Save user into dynamoDB
	err := c.SaveUser(request, tableName)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	return res, nil
}

func main() {
	lambda.Start(HandleRequest)
}
