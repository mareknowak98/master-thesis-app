package cmd

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"os"
)

// Connect will receive the $connect request
func Connect(request APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	lessonId := request.QueryStringParameters["lessonId"]

	connectionID := request.RequestContext.ConnectionID
	StoreSocket(lessonId, connectionID)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

// StoreSocket will store the id,connectionid map in dynamodb
func StoreSocket(id, connectionID string) error {
	m := LessonSocket{
		LessonId:     id,
		ConnectionID: connectionID,
	}

	av, err := dynamodbattribute.MarshalMap(m)
	if err != nil {
		log.Fatalln("Unable to marshal user socket map", err.Error())
	}

	tableName := os.Getenv("CONNECTIONS_TABLE")

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	}

	sess := GetSession()

	db := dynamodb.New(sess)

	_, err = db.PutItem(input)
	if err != nil {
		log.Fatal("INSERT ERROR", err.Error())
	}

	return nil
}
