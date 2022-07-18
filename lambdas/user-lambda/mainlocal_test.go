package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
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

	request.HTTPMethod = "GET"

	return request
}
