package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"mylearnproject/lambdas/grades-lambda/cmd"
)

//func HandleRequest(ctx context.Context, lambdaInp cmd.InputGrades) (events.APIGatewayProxyResponse, error) {
//	//fmt.Println(lambdaInp)
//
//	// Initialize AWS client
//	c := cmd.NewClient("eu-central-1")
//
//	// Save data to DynamoDB
//	c.SaveGrade(lambdaInp, "mylearn-grades")
//
//	var tmp events.APIGatewayProxyResponse
//	return tmp, nil
//}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//Initialize response object
	var resp events.APIGatewayProxyResponse

	// Initialize AWS client
	c := cmd.NewClient("eu-central-1")

	tableName := "mylearn-grades"

	// Save data to DynamoDB if POST
	if request.HTTPMethod == "POST" {
		err := c.SaveGrade(request, tableName)
		if err != nil {
			fmt.Println(err)
			resp.StatusCode = 400
			resp.Body = err.Error()
			return resp, nil
		}
		resp.StatusCode = 200
		return resp, nil
	}

	// GET method
	if request.HTTPMethod == "GET" {
		err := c.GetGrades(request, tableName)
		if err != nil {
			fmt.Println(err)
			resp.StatusCode = 400
			resp.Body = err.Error()
			return resp, nil
		}
		resp.StatusCode = 200
		return resp, nil
	}

	resp.Body = "Unhandled method request"
	return resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
