package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"mylearnproject/lambdas/grades-lambda/cmd"
)

func HandleRequest(ctx context.Context, lambdaInp cmd.InputGrades) (events.APIGatewayProxyResponse, error) {
	//fmt.Println(lambdaInp)

	// Initialize AWS client
	c := cmd.NewClient("eu-central-1")

	// Save data to DynamoDB
	c.SaveGrade(lambdaInp, "mylearn-grades")

	var tmp events.APIGatewayProxyResponse
	return tmp, nil
}

func main() {
	lambda.Start(HandleRequest)
}

