package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

// Handler is the base handler that will receive all web socket request
func Handler(request APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Request :%#v\n", request)

	switch request.RequestContext.RouteKey {
	case "$connect":
		return Connect(request)
	case "$disconnect":
		return Disconnect(request)
	default:
		return Default(request)
	}
}

func main() {
	lambda.Start(Handler)
}

// GetSession creates a new aws session and returns it
func GetSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	})
	if err != nil {
		log.Fatalln("Unable to create AWS session", err.Error())
	}

	return sess
}
