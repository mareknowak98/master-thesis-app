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
//	_ = os.Setenv("USER_POOL_ID", "eu-central-1_NT3RKgZjO")
//	_ = os.Setenv("REGION", "eu-central-1")
//	_ = os.Setenv("COGNITO_CLIENT_ID", "72bjtaumn8kqm3s3o51lt7sis1")
//
//	request.HTTPMethod = "POST"
//	request.Path = "/login"
//
//	request.Body = "{\n   \"username\" : \"testuser3\",\n   \"password\" : \"Passw0rd!\"\n}"
//
//	return request
//}

func getDebugInput() events.APIGatewayProxyRequest {
	var request events.APIGatewayProxyRequest
	_ = os.Setenv("USER_TABLE", "cognito-users")
	_ = os.Setenv("USER_POOL_ID", "eu-central-1_LsCgMRDvD")
	_ = os.Setenv("REGION", "eu-central-1")
	_ = os.Setenv("COGNITO_CLIENT_ID", "2jc39495jj7r6rpltst6b8ps80")

	request.HTTPMethod = "POST"
	request.Path = "/register"

	request.Body = "{\n   \"username\" : \"testuser4\",\n   \"email\" : \"mark.now997@gmail.com\",\n   \"password1\" : \"Passw0rd!\",\n   \"password2\" : \"Passw0rd!\"\n}"

	return request
}
