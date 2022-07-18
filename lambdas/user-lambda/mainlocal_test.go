package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"os"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	input := getDebugInput()

	res, _ := HandleRequest(input)

	fmt.Println("Response:")
	fmt.Println(res.StatusCode)
	fmt.Println(res.Body)
}

func getDebugInput() events.APIGatewayProxyRequest {
	var request events.APIGatewayProxyRequest
	_ = os.Setenv("USER_TABLE", "cognito-users")
	_ = os.Setenv("REGION", "eu-central-1")

	request.HTTPMethod = "GET"
	request.Path = "/users"
	return request
}
