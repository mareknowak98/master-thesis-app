package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"mylearnproject/lambdas/grades-lambda/cmd"
	"testing"
)

// File to test running lambda locally

func TestHandleRequest(t *testing.T) {

	input := getDebugInput()

	res, _ := HandleRequest(input)

	fmt.Println("Response:")
	fmt.Println(res.StatusCode)
	fmt.Println(res.Body)

}

func getDebugInput() events.APIGatewayProxyRequest {
	var inp cmd.InputGrades

	err := json.Unmarshal([]byte(`
	{
		"UserId": "123",
		"ClassYear": "ab22ba",
		"Grade": "4+"
	}`), &inp)

	var request events.APIGatewayProxyRequest
	if err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(inp)
	if err != nil {
		fmt.Println(err)
	}
	request.Body = string(b)
	request.HTTPMethod = "POST"
	fmt.Printf("Input %s\n", inp)
	fmt.Printf("Request body %s\n", request.Body)
	fmt.Printf("Request method %s\n", request.HTTPMethod)

	return request
}

//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//
//	request.HTTPMethod = "GET"
//	fmt.Printf("Request body %s\n", request.Body)
//	fmt.Printf("Request method %s\n", request.HTTPMethod)
//
//	return request
//}
