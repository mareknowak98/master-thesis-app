package cmd

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
)

// Connect will receive the $connect request
// It will handle the authorization also
func Connect(request APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	parsedToken, err := DecodeToken(request.Headers["Authorization"])
	if err != nil {
		fmt.Println("eror")
		fmt.Println(err)
	}
	fmt.Println("Connection fnc after stareted")

	fmt.Printf("token claims: %#v\n", parsedToken.Claims.(jwt.MapClaims))

	id := parsedToken.Claims.(jwt.MapClaims)["username"].(string)

	fmt.Printf("User id %v\n", id)

	fmt.Printf("RequestContext %v\n", request.RequestContext)
	fmt.Printf("ConnectionID %v\n", request.RequestContext.ConnectionID)

	connectionID := request.RequestContext.ConnectionID
	StoreSocket(id, connectionID)

	fmt.Println("After store socket")

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

type UserSocket struct {
	UserID       string `json:"userId"`
	ConnectionID string `json:"connectionID"`
}

// StoreSocket will store the id,connectionid map in dynamodb
func StoreSocket(id, connectionID string) error {
	fmt.Println("In store socket")
	m := UserSocket{
		UserID:       id,
		ConnectionID: connectionID,
	}

	fmt.Printf("UserSocket %v\n", m)

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

	fmt.Println("EOF store socket")

	return nil
}
