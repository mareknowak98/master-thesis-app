package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"mylearnproject/lambdas/grades-lambda/cmd"
)

func responseGenerator(code int, message string) events.APIGatewayProxyResponse {
	var resp events.APIGatewayProxyResponse
	resp.StatusCode = code
	resp.Body = message
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Access-Control-Allow-Origin"] = "*"
	resp.Headers = headers
	return resp
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Initialize AWS client
	c := cmd.NewClient("eu-central-1")

	tableName := "mylearn-grades"

	// Save data to DynamoDB if POST
	if request.HTTPMethod == "POST" {
		err := c.SaveGrade(request, tableName)
		if err != nil {
			return responseGenerator(400, err.Error()), nil
		}
		return responseGenerator(200, ""), nil
	}

	// GET method
	if request.HTTPMethod == "GET" {
		out, err := c.GetGrades(request, tableName)
		if err != nil {
			return responseGenerator(500, err.Error()), nil
		}
		return responseGenerator(200, out), nil
	}

	return responseGenerator(200, "Unhandled method request"), nil
}

func main() {
	lambda.Start(HandleRequest)
}
