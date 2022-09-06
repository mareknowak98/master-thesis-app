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
//		   "correctAnswer" : "1.somecontent"
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

//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//	_ = os.Setenv("TESTS_TABLE", "mylearn-tests")
//	_ = os.Setenv("RESULTS_TABLE", "mylearn-tests-results")
//	_ = os.Setenv("REGION", "eu-central-1")
//
//	request.HTTPMethod = "POST"
//	request.Path = "/results"
//
//	var inp cmd.CheckTestInput
//	err := json.Unmarshal([]byte(`
//			{
//			   "testId" : "1",
//			   "questionAnswer" : [
//				{
//					"combinedKey" : "1AB:1",
//					"correctQuestionAnswer": "2.eluwina"
//				},
//				{
//					"combinedKey" : "1AB:2",
//					"correctQuestionAnswer": "1.somecontent"
//				}
//				]
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
//
//	m2 := make(map[string]string)
//	m2["Authorization"] = "eyJraWQiOiIrRUdhblhVdndMOUh6czVDUnRlcktxS282ZXVMUlo3Sm9aVE1VVzdzOFdrPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiJmMDJlNzE4Yi0xZGUzLTRmYTMtOGY0NS0wNmFjOGYxZGFlMDIiLCJjb2duaXRvOmdyb3VwcyI6WyJzdHVkZW50LWdyb3VwIl0sImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC5ldS1jZW50cmFsLTEuYW1hem9uYXdzLmNvbVwvZXUtY2VudHJhbC0xX0xzQ2dNUkR2RCIsImNsaWVudF9pZCI6IjJqYzM5NDk1amo3cjZycGx0c3Q2YjhwczgwIiwib3JpZ2luX2p0aSI6ImEzOWEwNzU2LWVmYjItNDI0ZC1hOGMwLTBlYjEyODMwMGFlZSIsImV2ZW50X2lkIjoiN2FkNmM1ZDItZmZmNC00YzRkLWFhZGEtNzljODNiYzE2YjY3IiwidG9rZW5fdXNlIjoiYWNjZXNzIiwic2NvcGUiOiJhd3MuY29nbml0by5zaWduaW4udXNlci5hZG1pbiIsImF1dGhfdGltZSI6MTY2MjUwNzA1MywiZXhwIjoxNjYyNTkzNDUzLCJpYXQiOjE2NjI1MDcwNTMsImp0aSI6ImI5NGFkMzRmLWY1NTYtNGEwOS05YjUzLTg3MWRhNTkyMDkzNyIsInVzZXJuYW1lIjoidGVzdHVzZXI2In0.eMlpuQZHTCTuluEsFheCZ3TKspOepd-v6lF2mXqSzhxDqfaefZxynZhZzGq_aOIJMm7e-RDaR3M689rIEB8XpMdwDku3A4CWyqo3TBzHhLAOh3C_nB0KFAhAqNqayfqZimFxm5MTCbwADItMlQ8OWyKHyvFx7suPRkLs1cnag7Z35UOkYV1W3iQw87sQE_sZI2tlV9te_sBG3sfYBNAAcN5skGJioD6hgjep0uI-Z2V5OVCayeAp_LOAI8wX7J9oqJ7spYtRk5jz5CLRfugrEffqdNVjmmeEctErD18rGGiOHEBb-OJUJoYCl9smE1FQ4wVJ9owKDFsmF37L2Z1esQ"
//	request.Headers = m2
//
//	return request
//}

//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//	_ = os.Setenv("TESTS_TABLE", "mylearn-tests")
//	_ = os.Setenv("RESULTS_TABLE", "mylearn-tests-results")
//	_ = os.Setenv("REGION", "eu-central-1")
//
//	m := make(map[string]string)
//	m["testId"] = "1"
//	request.QueryStringParameters = m
//	request.HTTPMethod = "GET"
//	request.Path = "/results"
//	return request
//}

func getDebugInput() events.APIGatewayProxyRequest {
	var request events.APIGatewayProxyRequest
	_ = os.Setenv("TESTS_TABLE", "mylearn-tests")
	_ = os.Setenv("RESULTS_TABLE", "mylearn-tests-results")
	_ = os.Setenv("REGION", "eu-central-1")

	m := make(map[string]string)
	m["testId"] = "1"
	m["userId"] = "testuser6"

	request.QueryStringParameters = m
	request.HTTPMethod = "GET"
	request.Path = "/results"
	return request
}
