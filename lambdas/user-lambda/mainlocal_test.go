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
//	_ = os.Setenv("USER_TABLE", "cognito-users")
//	var inp cmd.AddUserInput
//
//	err := json.Unmarshal([]byte(`
//			{
//				"UserClass": "1AB",
//				"Username": "testuser12"
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
	_ = os.Setenv("USER_TABLE", "cognito-users")

	var inp cmd.AddUserInput

	err := json.Unmarshal([]byte(`
			{
				"UserClass": "1AB",
				"Username": "testuser10"
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

//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//	_ = os.Setenv("CLASS_TABLE", "mylearn-classes")
//	_ = os.Setenv("REGION", "eu-central-1")
//	_ = os.Setenv("USER_POOL_ID", "eu-central-1_LsCgMRDvD")
//	_ = os.Setenv("USER_TABLE", "cognito-users")
//
//	m := make(map[string]string)
//	m["info"] = "class"
//	request.QueryStringParameters = m
//	request.HTTPMethod = "GET"
//	request.Path = "/me"
//	m2 := make(map[string]string)
//	m2["Authorization"] = "eyJraWQiOiIrRUdhblhVdndMOUh6czVDUnRlcktxS282ZXVMUlo3Sm9aVE1VVzdzOFdrPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiI3ZTM1YTFkMy1jYmQyLTQ4OTctOTZiZi1iNjVhYjczYTc3Y2EiLCJjb2duaXRvOmdyb3VwcyI6WyJzdHVkZW50LWdyb3VwIl0sImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC5ldS1jZW50cmFsLTEuYW1hem9uYXdzLmNvbVwvZXUtY2VudHJhbC0xX0xzQ2dNUkR2RCIsImNsaWVudF9pZCI6IjJqYzM5NDk1amo3cjZycGx0c3Q2YjhwczgwIiwib3JpZ2luX2p0aSI6Ijc2NmJhMDQ4LWZjMjgtNDgxYy04OGM3LWIxMTU0ZmI3ZTYzNyIsImV2ZW50X2lkIjoiNWE4OTUwOGMtZmE1Yy00ZWFmLWEzYTgtODhlMzEyMWY1MTA2IiwidG9rZW5fdXNlIjoiYWNjZXNzIiwic2NvcGUiOiJhd3MuY29nbml0by5zaWduaW4udXNlci5hZG1pbiIsImF1dGhfdGltZSI6MTY2MTY5ODA2NCwiZXhwIjoxNjYxNzAxNjY0LCJpYXQiOjE2NjE2OTgwNjQsImp0aSI6IjIzY2RmMmE4LWYzNjQtNDE5MS05YTFhLTRiZjNmNzlmYzgxNCIsInVzZXJuYW1lIjoidGVzdHVzZXIxMiJ9.Yedso0IXQnfmaVT5Q-WOsR9FDuKeoA8j18rI44Zhl_Jt-L-1YrUsC4TmE6vmrRjEQjXbc5IcbEor8S6LidHMCwqxgb2VQXlBkMuktzn9C8Ch-MXd5t0R6hmzGU5_qUAm926raoFr338A6B1zOKGobRWX-Etr2pNbnALhsmyEeBO7U1X1NOEsCaZs1ffxMXpWR8rilaIQFjMP82KiR8Mg79LbflVS7LE7gvXZaFFDuH0vXrsv-W6lRMrfilHweQ9RYDi3dNvUS1ZrP9Nmwz0_7gD3QRYQf5nVF7nNu5nnrZHFCK5LFuS2GWx6cHYgbw9mfi-jwdb_9yHB59tAKUO8ug"
//	request.Headers = m2
//
//	return request
//}
