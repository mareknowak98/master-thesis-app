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

//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//	_ = os.Setenv("REGION", "eu-central-1")
//	_ = os.Setenv("MYLEARN_BUCKET", "mylearn-materials-eu-central-1")
//
//	request.HTTPMethod = "GET"
//	request.Path = "/files"
//
//	m := make(map[string]string)
//	m["folder"] = "eluwina"
//
//	request.QueryStringParameters = m
//
//	return request
//}

func getDebugInput() events.APIGatewayProxyRequest {
	var request events.APIGatewayProxyRequest
	_ = os.Setenv("REGION", "eu-central-1")
	_ = os.Setenv("MYLEARN_BUCKET", "mylearn-materials-eu-central-1")

	request.HTTPMethod = "GET"
	request.Path = "/folders"

	m := make(map[string]string)
	m["subfolder"] = ""

	request.QueryStringParameters = m

	return request
}
