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

func getDebugInput() events.APIGatewayProxyRequest {
	var request events.APIGatewayProxyRequest
	//_ = os.Setenv("LESSONS_TABLE", "mylearn-lessons")
	_ = os.Setenv("REGION", "eu-central-1")
	_ = os.Setenv("S3_MATERIALS", "mylearn-materials-eu-central-1")

	request.HTTPMethod = "POST"
	request.Path = "/materials"

	//file, err := ioutil.ReadFile("meme.jpg")
	////fmt.Println(string(file))
	////fmt.Println("-------")
	//fmt.Println(err)
	//
	//esa := fmt.Sprintf("\t\t{\n\t\t   \"Item\" : %s,\n\t\t   \"ItemKey\" : \"costamKeysa\",\n\t\t}", file)
	//fmt.Println("-------")
	//
	//fmt.Println(esa)
	//
	//var inp cmd.S3Input
	//err = json.Unmarshal([]byte(esa), &inp)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//b, err := json.Marshal(inp)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//request.Body = string(b)

	m := make(map[string]string)
	m["filename"] = "new_file.jpg"

	request.Headers = m

	return request
}
