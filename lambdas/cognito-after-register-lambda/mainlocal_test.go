package main

import (
	"context"
	"fmt"
	"mylearnproject/lambdas/cognito-after-register-lambda/cmd"
	"testing"
)

// File to test running lambda locally

func TestHandleRequest(t *testing.T) {

	input := getDebugInput()

	res, _ := HandleRequest(context.Background(), input)

	fmt.Println("Response:")
	fmt.Println(res)

}

func getDebugInput() cmd.CognitoEventRequest {
	var request cmd.CognitoEventRequest
	tmp := make(map[string]map[string]string)
	tmp2 := make(map[string]string)
	tmp2["email"] = "jane@example.com"
	tmp2["sub"] = "ebfddb40-a7b9-4a25-8244-d49e6ee923b2"

	tmp["userAttributes"] = tmp2

	request.Request = tmp
	request.Username = "testuser2"
	return request
}
