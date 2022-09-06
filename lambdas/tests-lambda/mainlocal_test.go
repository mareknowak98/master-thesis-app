package main

import (
	"fmt"
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
//	_ = os.Setenv("TESTS_TABLE", "mylearn-tests")
//	_ = os.Setenv("REGION", "eu-central-1")
//
//	request.HTTPMethod = "POST"
//	request.Path = "/tests"
//
//	var inp cmd.TestQuestionInput
//	err := json.Unmarshal([]byte(`
//		{
//		   "testId" : "1",
//		   "combinedKey" : "1AB:2",
//		   "question" : "sometypa",
//		   "answers" : ["1.somecontent", "2.eluwina"],
//		   "correctAnswer" : "1"
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
//
//	return request
//}

//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//	_ = os.Setenv("TESTS_TABLE", "mylearn-tests")
//	_ = os.Setenv("REGION", "eu-central-1")
//
//	m := make(map[string]string)
//	m["testId"] = "1"
//	request.QueryStringParameters = m
//	request.HTTPMethod = "GET"
//	request.Path = "/tests"
//	return request
//}

//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//	_ = os.Setenv("TESTS_TABLE", "mylearn-tests")
//	_ = os.Setenv("REGION", "eu-central-1")
//
//	request.HTTPMethod = "GET"
//	request.Path = "/tests"
//	return request
//}

//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//	_ = os.Setenv("TESTS_TABLE", "mylearn-tests")
//	_ = os.Setenv("REGION", "eu-central-1")
//
//	request.HTTPMethod = "DELETE"
//	request.Path = "/tests"
//
//	m := make(map[string]string)
//	m["testId"] = "1"
//	m["combinedKey"] = "1AB:1"
//
//	request.QueryStringParameters = m
//
//	return request
//}
