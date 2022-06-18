package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
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
	var inp events.APIGatewayProxyRequest

	err := json.Unmarshal([]byte(`
	{
		"UserId": "123",
		"ClassYear": "sometest",
		"Grade": "4+"
	}`), &inp)

	if err != nil {
		fmt.Println(err)
	}

	return inp
}
