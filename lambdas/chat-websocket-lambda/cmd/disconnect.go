package cmd

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"os"
)

// Disconnect will receive the $disconnect requests
func Disconnect(request APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	connectionID := request.RequestContext.ConnectionID

	RemoveSocket(connectionID)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

// RemoveSocket will remove the id,connectionId socket from dynamodb
func RemoveSocket(connectionID string) {
	tableName := os.Getenv("CONNECTIONS_TABLE")

	fmt.Println(connectionID)
	inputScan := &dynamodb.ScanInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {
				S: aws.String(connectionID),
			},
		},
		FilterExpression: aws.String("connectionID = :a"),
		TableName:        aws.String(tableName),
	}

	db := dynamodb.New(GetSession())

	result, err := db.Scan(inputScan)

	fmt.Println(err)
	fmt.Println(result)

	fmt.Println("---------------")
	userId := result.Items[0]["userId"].S
	fmt.Println(*userId)

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": &dynamodb.AttributeValue{
				S: aws.String(*userId),
			},
		},
	}

	_, err = db.DeleteItem(input)
	if err != nil {
		log.Fatalln("Unable to remove user socket map", err.Error())
	}
}
