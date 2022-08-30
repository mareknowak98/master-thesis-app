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
//	_ = os.Setenv("LESSONS_TABLE", "mylearn-lessons")
//	_ = os.Setenv("REGION", "eu-central-1")
//
//	request.HTTPMethod = "POST"
//	request.Path = "/slides"
//
//	var inp cmd.LessonSlide
//	err := json.Unmarshal([]byte(`
//		{
//		   "lessonId" : "12",
//		   "slideId" : "2",
//		   "slideType" : "sometypa",
//		   "slideContent" : "somecontent",
//		   "questionAnswers" : "someanswers"
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
//	_ = os.Setenv("LESSONS_TABLE", "mylearn-lessons")
//	_ = os.Setenv("REGION", "eu-central-1")
//
//	request.HTTPMethod = "DELETE"
//	request.Path = "/slides"
//
//	m := make(map[string]string)
//	m["lesson"] = "12"
//	m["slide"] = "2"
//
//	request.QueryStringParameters = m
//
//	return request
//}

//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//	_ = os.Setenv("LESSONS_TABLE", "mylearn-lessons")
//	_ = os.Setenv("REGION", "eu-central-1")
//
//	request.HTTPMethod = "GET"
//	request.Path = "/slides"
//
//	m := make(map[string]string)
//	m["lesson"] = "12"
//	//m["slide"] = "2"
//
//	request.QueryStringParameters = m
//
//	return request
//}

//func getDebugInput() events.APIGatewayProxyRequest {
//	var request events.APIGatewayProxyRequest
//	_ = os.Setenv("LESSONS_TABLE", "mylearn-lessons")
//	_ = os.Setenv("REGION", "eu-central-1")
//
//	request.HTTPMethod = "GET"
//	request.Path = "/lessons"
//
//	//m := make(map[string]string)
//	//m["lesson"] = "12"
//	//m["slide"] = "2"
//
//	//request.QueryStringParameters = m
//
//	return request
//}

func getDebugInput() events.APIGatewayProxyRequest {
	var request events.APIGatewayProxyRequest
	_ = os.Setenv("LESSONS_CONNECTIONS_TABLE", "mylearn-lessons-connections")
	_ = os.Setenv("REGION", "eu-central-1")

	m := make(map[string]string)
	m["ongoing"] = "true"
	request.QueryStringParameters = m
	request.HTTPMethod = "GET"
	request.Path = "/lessons"
	return request
}
