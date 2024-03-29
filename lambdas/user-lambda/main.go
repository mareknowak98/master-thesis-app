package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"mylearnproject/lambdas/user-lambda/cmd"
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
	tableName := os.Getenv("USER_TABLE")
	region := os.Getenv("REGION")

	c := cmd.NewClient(region)

	switch request.Path {
	case "/users":
		switch request.HTTPMethod {
		case "GET":
			resp, err := c.GetUsers(request, tableName)
			if err != nil {
				return responseGenerator(500, err.Error()), nil
			}
			return responseGenerator(200, resp), nil

		default:
			return responseGenerator(400, "No such method"), nil
		}
	case "/classes":
		switch request.HTTPMethod {
		case "GET":
			resp, err := c.GetClasses(request)
			if err != nil {
				return responseGenerator(500, err.Error()), nil
			}
			return responseGenerator(200, resp), nil
		case "POST":
			resp, err := c.SetClasses(request)
			if err != nil {
				return responseGenerator(500, err.Error()), nil
			}
			return responseGenerator(200, resp), nil
		case "DELETE":
			resp, err := c.DeleteClasses(request)
			if err != nil {
				return responseGenerator(500, err.Error()), nil
			}
			return responseGenerator(200, resp), nil
		default:
			return responseGenerator(400, "No such method"), nil
		}

	case "/classUsers":
		switch request.HTTPMethod {
		case "POST":
			resp, err := c.AddUserToClass(request)
			if err != nil {
				return responseGenerator(500, err.Error()), nil
			}
			return responseGenerator(200, resp), nil
		case "DELETE":
			resp, err := c.DeleteUserFromClass(request)
			if err != nil {
				return responseGenerator(500, err.Error()), nil
			}
			return responseGenerator(200, resp), nil
		case "GET":
			resp, err := c.GetUsersFromClass(request)
			if err != nil {
				return responseGenerator(500, err.Error()), nil
			}
			return responseGenerator(200, resp), nil
		default:
			return responseGenerator(400, "No such method"), nil
		}
	case "/me":
		switch request.HTTPMethod {
		case "GET":
			resp, err := c.GetMe(request)
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
