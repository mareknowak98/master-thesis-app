package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	//"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
)

// Disconnect will receive the $disconnect requests
func Disconnect(request APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	//parsedToken, err := cmd.DecodeToken(request.Headers["Authorization"])
	//if err != nil {
	//	fmt.Println("eror")
	//	fmt.Println(err)
	//}
	//fmt.Println("Connection fnc after stareted")
	//
	//fmt.Printf("token claims: %#v\n", parsedToken.Claims.(jwt.MapClaims))
	//
	//id := parsedToken.Claims.(jwt.MapClaims)["username"].(string)

	//id := request.RequestContext.Authorizer.(map[string]interface{})["cognito:username"].(string)
	connectionID := request.RequestContext.ConnectionID
	RemoveSocket(connectionID)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

// RemoveSocket will remove the id,connectionId socket from dynamodb
func RemoveSocket(connectionID string) {
	tableName := os.Getenv("MESSAGES_TABLE")

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"connectionID": &dynamodb.AttributeValue{
				S: aws.String(connectionID),
			},
		},
	}

	db := dynamodb.New(GetSession())

	_, err := db.DeleteItem(input)
	if err != nil {
		log.Fatalln("Unable to remove user socket map", err.Error())
	}
}
