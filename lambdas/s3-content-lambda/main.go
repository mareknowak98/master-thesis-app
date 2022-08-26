package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"mylearnproject/lambdas/s3-content-lambda/cmd"
	"os"
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
	region := os.Getenv("REGION")
	s3Materials := os.Getenv("S3_MATERIALS")

	fmt.Printf("Request: %#v\n", request)

	c := cmd.NewClient(region)

	switch request.Path {
	case "/materials":
		switch request.HTTPMethod {
		case "POST":
			resp, err := c.SaveMaterials(request, s3Materials)
			if err != nil {
				return responseGenerator(500, err.Error()), nil
			}
			return responseGenerator(200, resp), nil
		case "GET":
			resp, err := c.GetMaterials(request, s3Materials)
			if err != nil {
				return responseGenerator(500, err.Error()), nil
			}
			return responseGenerator(200, resp), nil

		default:
			return responseGenerator(400, "No such method"), nil
		}
	default:
		return responseGenerator(400, "No such endpoint"), nil
	}
}

func main() {
	lambda.Start(HandleRequest)
}
