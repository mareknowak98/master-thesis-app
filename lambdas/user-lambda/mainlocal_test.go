package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"mylearnproject/lambdas/user-lambda/cmd"
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
//	_ = os.Setenv("USER_TABLE", "cognito-users")
//	_ = os.Setenv("REGION", "eu-central-1")
//	_ = os.Setenv("USER_POOL_ID", "eu-central-1_LsCgMRDvD")
//
//	m := make(map[string]string)
//	m["group"] = "all"
//	request.QueryStringParameters = m
//	request.HTTPMethod = "GET"
//	request.Path = "/users"
//	return request
//}

//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//	_ = os.Setenv("CLASS_TABLE", "mylearn-classes")
//	_ = os.Setenv("REGION", "eu-central-1")
//	_ = os.Setenv("USER_POOL_ID", "eu-central-1_LsCgMRDvD")
//
//	request.HTTPMethod = "GET"
//	request.Path = "/classes"
//	return request
//}

//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//	_ = os.Setenv("CLASS_TABLE", "mylearn-classes")
//	_ = os.Setenv("REGION", "eu-central-1")
//	_ = os.Setenv("USER_POOL_ID", "eu-central-1_LsCgMRDvD")
//
//	var inp cmd.AddClassInput
//
//	err := json.Unmarshal([]byte(`
//		{
//			"UserClass": "1ABCD"
//		}`), &inp)
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	b, err := json.Marshal(inp)
//	if err != nil {
//		fmt.Println(err)
//	}
//	request.Body = string(b)
//	request.HTTPMethod = "POST"
//	fmt.Printf("Input %s\n", inp)
//	fmt.Printf("Request body %s\n", request.Body)
//
//	request.HTTPMethod = "POST"
//	request.Path = "/classes"
//	return request
//}

//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//	_ = os.Setenv("CLASS_TABLE", "mylearn-classes")
//	_ = os.Setenv("REGION", "eu-central-1")
//	_ = os.Setenv("USER_POOL_ID", "eu-central-1_LsCgMRDvD")
//
//	var inp cmd.AddClassInput
//
//	err := json.Unmarshal([]byte(`
//			{
//				"UserClass": "1ABCD"
//			}`), &inp)
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	b, err := json.Marshal(inp)
//	if err != nil {
//		fmt.Println(err)
//	}
//	request.Body = string(b)
//	request.HTTPMethod = "DELETE"
//	fmt.Printf("Input %s\n", inp)
//	fmt.Printf("Request body %s\n", request.Body)
//	request.Path = "/classes"
//	return request
//}
//
//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//	_ = os.Setenv("CLASS_TABLE", "mylearn-classes")
//	_ = os.Setenv("REGION", "eu-central-1")
//	_ = os.Setenv("USER_POOL_ID", "eu-central-1_LsCgMRDvD")
//
//	var inp cmd.AddUserInput
//
//	err := json.Unmarshal([]byte(`
//			{
//				"UserClass": "1AB",
//				"Username": "somerandomuser"
//			}`), &inp)
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	b, err := json.Marshal(inp)
//	if err != nil {
//		fmt.Println(err)
//	}
//	request.Body = string(b)
//	request.HTTPMethod = "POST"
//	fmt.Printf("Input %s\n", inp)
//	fmt.Printf("Request body %s\n", request.Body)
//	request.Path = "/classUsers"
//	return request
//}

func getDebugInput() events.APIGatewayProxyRequest {
	var request events.APIGatewayProxyRequest
	_ = os.Setenv("CLASS_TABLE", "mylearn-classes")
	_ = os.Setenv("REGION", "eu-central-1")
	_ = os.Setenv("USER_POOL_ID", "eu-central-1_LsCgMRDvD")

	var inp cmd.AddUserInput

	err := json.Unmarshal([]byte(`
			{
				"UserClass": "1AB",
				"Username": "somerandomuser"
			}`), &inp)

	if err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(inp)
	if err != nil {
		fmt.Println(err)
	}
	request.Body = string(b)
	request.HTTPMethod = "DELETE"
	fmt.Printf("Input %s\n", inp)
	fmt.Printf("Request body %s\n", request.Body)
	request.Path = "/classUsers"
	return request
}
