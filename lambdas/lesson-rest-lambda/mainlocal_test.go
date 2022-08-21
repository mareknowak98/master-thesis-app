package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"mylearnproject/lambdas/lesson-rest-lambda/cmd"
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

func getDebugInput() events.APIGatewayProxyRequest {
	var request events.APIGatewayProxyRequest
	_ = os.Setenv("LESSONS_TABLE", "mylearn-lessons")
	_ = os.Setenv("REGION", "eu-central-1")

	request.HTTPMethod = "POST"
	request.Path = "/lessons"

	var inp cmd.LessonSlide
	err := json.Unmarshal([]byte(`
		{
		   "lessonId" : 12,
		   "slideId" : 2,
		   "slideType" : "sometypa",
		   "slideContent" : "somecontent",
		   "questionAnswers" : "someanswers"
		}`), &inp)

	if err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(inp)
	if err != nil {
		fmt.Println(err)
	}
	request.Body = string(b)

	return request
}
