package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"mylearnproject/lambdas/lesson-websocket-lambda/cmd"
	"os"
)

// Handler is the base handler that will receive all web socket request
func Handler(request cmd.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Request :%#v\n", request)
	region := os.Getenv("REGION")
	c := cmd.NewClient(region)
	switch request.RequestContext.RouteKey {
	case "$connect":
		return cmd.Connect(request)
	case "$disconnect":
		return cmd.Disconnect(request)
	default:
		return c.Default(request)
	}
}

func main() {
	lambda.Start(Handler)
}
