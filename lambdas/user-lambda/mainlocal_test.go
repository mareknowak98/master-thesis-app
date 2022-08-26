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
//	_ = os.Setenv("USER_TABLE", "cognito-users")
//	var inp cmd.AddUserInput
//
//	err := json.Unmarshal([]byte(`
//			{
//				"UserClass": "1AB",
//				"Username": "testuser10"
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

//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//	_ = os.Setenv("CLASS_TABLE", "mylearn-classes")
//	_ = os.Setenv("REGION", "eu-central-1")
//	_ = os.Setenv("USER_POOL_ID", "eu-central-1_LsCgMRDvD")
//	_ = os.Setenv("USER_TABLE", "cognito-users")
//
//	var inp cmd.AddUserInput
//
//	err := json.Unmarshal([]byte(`
//			{
//				"UserClass": "1AB",
//				"Username": "testuser10"
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
//	request.Path = "/classUsers"
//	return request
//}

//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//	_ = os.Setenv("CLASS_TABLE", "mylearn-classes")
//	_ = os.Setenv("REGION", "eu-central-1")
//	_ = os.Setenv("USER_POOL_ID", "eu-central-1_LsCgMRDvD")
//	_ = os.Setenv("USER_TABLE", "cognito-users")
//
//	m := make(map[string]string)
//	m["class"] = "1AB"
//	request.QueryStringParameters = m
//	request.HTTPMethod = "GET"
//	request.Path = "/classUsers"
//	return request
//}

func getDebugInput() events.APIGatewayProxyRequest {
	var request events.APIGatewayProxyRequest
	_ = os.Setenv("CLASS_TABLE", "mylearn-classes")
	_ = os.Setenv("REGION", "eu-central-1")
	_ = os.Setenv("USER_POOL_ID", "eu-central-1_LsCgMRDvD")
	_ = os.Setenv("USER_TABLE", "cognito-users")

	m := make(map[string]string)
	m["info"] = "class"
	request.QueryStringParameters = m
	request.HTTPMethod = "GET"
	request.Path = "/me"
	m2 := make(map[string]string)
	m2["Authorization"] = "eyJraWQiOiIrRUdhblhVdndMOUh6czVDUnRlcktxS282ZXVMUlo3Sm9aVE1VVzdzOFdrPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiJmOGI4NTE1Yi0zZTgwLTQ4YzItYTk3Ny1jNmM0ODYxMjk0NjkiLCJjb2duaXRvOmdyb3VwcyI6WyJzdHVkZW50LWdyb3VwIl0sImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC5ldS1jZW50cmFsLTEuYW1hem9uYXdzLmNvbVwvZXUtY2VudHJhbC0xX0xzQ2dNUkR2RCIsImNsaWVudF9pZCI6IjJqYzM5NDk1amo3cjZycGx0c3Q2YjhwczgwIiwib3JpZ2luX2p0aSI6IjRkYjgzMjJjLWViMjQtNGU0ZS1hZmQxLTE5OWNmZjBmZTc5OSIsImV2ZW50X2lkIjoiODRiMzFjNmEtMGI4YS00ZmI2LTgyMjYtOWQ4YjYyMzRhYjI5IiwidG9rZW5fdXNlIjoiYWNjZXNzIiwic2NvcGUiOiJhd3MuY29nbml0by5zaWduaW4udXNlci5hZG1pbiIsImF1dGhfdGltZSI6MTY2MDg1OTc3NCwiZXhwIjoxNjYwODYzMzc0LCJpYXQiOjE2NjA4NTk3NzQsImp0aSI6ImQ1NDUzZDhiLTAwMTItNGIzZS1iNmQ0LWUxY2U0MDAyMTY3MiIsInVzZXJuYW1lIjoidGVzdHVzZXIxMCJ9.sZO-MKGz8DuLnCj8fpZolmktJITOJdm8orVwGGfk-Isj-MWseYVRs92yIk27IkBeH8pJMiu-P-THhNZJZVQjDVX9ngPcncePNrdG2AdQVwV-LIL38V5Bps3JtrOmiHFk4S2iYvuor7ioh2XqMUtZexzxfJlLR2a4K2uiypVBw954r5nOZTGD5qdsh8bDhNRt96TBwXzhNZ5AoQtBh8aTpyhB_wcBu9MOb0VLIQYEm-tkTn8E54BHLYyDuHXlHHNKyYRld7Xojv6YaKEfXIYb9x54Ficgwr-RBm3IIQ7CYOH4Uq_OSRsWxQTTE9fga55ywE5iYRiM6klmZd9-ACDYdw"
	request.Headers = m2

	return request
}
