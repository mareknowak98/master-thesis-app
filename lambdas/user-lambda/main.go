package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"os"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var resp events.APIGatewayProxyResponse
	tableName := os.Getenv("USER_TABLE")
	region := os.Getenv("REGION")

	fmt.Println(tableName)
	fmt.Println(region)

	fmt.Printf("Request: %#v\n", request)

	//c := cmd.NewClient(region)

	return resp, nil
}
