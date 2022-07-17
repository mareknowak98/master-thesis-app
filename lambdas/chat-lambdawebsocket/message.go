package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"os"
	"strings"
)

type MessageAction struct {
	Type    string          `json:"type"`
	Payload MessageWithInfo `json:"payload"`
}

//// MessagePayload ..
//type MessagePayload struct {
//	Message MessageWithInfo `json:"message"`
//}

// MessageWithInfo ..
type MessageWithInfo struct {
	To      string      `json:"to"`
	From    string      `json:"from"`
	Message interface{} `json:"message"`
}

// Default ..
func Default(request APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	tableName := os.Getenv("MESSAGES_TABLE")
	websocketApi := os.Getenv("WEBSOKET_API")

	b := MessageAction{}
	if err := json.NewDecoder(strings.NewReader(request.Body)).Decode(&b); err != nil {
		log.Println("Unable to decode body", err.Error())
	}

	fmt.Printf("decoded body %#v\n", b)

	data, _ := json.Marshal(b)
	fmt.Printf("MessageAction %#v\n", string(data))

	sess := GetSession()
	//fmt.Printf("Session %#v\n", sess)

	fmt.Printf("RequestContext.ConnectionID %#v\n", request.RequestContext.ConnectionID)

	db := dynamodb.New(sess)

	//queryInput := &dynamodb.QueryInput{
	//	TableName: aws.String(tableName),
	//	ExpressionAttributeNames: map[string]*string{
	//		"#id": aws.String("id"),
	//	},
	//	ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
	//		":id": &dynamodb.AttributeValue{
	//			S: aws.String(b.Payload.To),
	//			//S: aws.String("id"),
	//		},
	//	},
	//	KeyConditionExpression: aws.String("#id=:id"),
	//}
	//queryInput := &dynamodb.QueryInput{
	//	TableName: aws.String(tableName),
	//	ExpressionAttributeNames: map[string]*string{
	//		"#connectionID": aws.String("connectionID"),
	//	},
	//	ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
	//		":connectionID": &dynamodb.AttributeValue{
	//			S: aws.String(request.RequestContext.ConnectionID),
	//		},
	//	},
	//	KeyConditionExpression: aws.String("#connectionID=:connectionID"),
	//}

	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		ExpressionAttributeNames: map[string]*string{
			"#userId": aws.String("userId"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":userId": &dynamodb.AttributeValue{
				S: aws.String(b.Payload.To),
			},
		},
		KeyConditionExpression: aws.String("#userId=:userId"),
	}
	fmt.Printf("QueryInput %#v\n", queryInput)

	output, err := db.Query(queryInput)
	if err != nil {
		log.Println("Unable to find connection ID", err.Error())
		return events.APIGatewayProxyResponse{}, err
	}
	fmt.Println(output)

	conndectionID := output.Items[0]["connectionID"].S
	fmt.Printf("conndectionID %#v\n", *conndectionID)

	userSock := UserSocket{}
	userSock.ConnectionID = *conndectionID
	fmt.Printf("userSocks %#v\n", userSock)

	/////

	input := &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(userSock.ConnectionID),
		Data:         data,
	}
	fmt.Printf("userSock %#v\n", userSock)
	fmt.Printf("input %#v\n", input)
	fmt.Printf("input data %#v\n", string(input.Data))

	fmt.Printf("%s.execute-api.eu-central-1.amazonaws.com/test", websocketApi)
	apigateway := apigatewaymanagementapi.New(sess, aws.NewConfig().WithEndpoint(fmt.Sprintf("%s.execute-api.eu-central-1.amazonaws.com/test", websocketApi)))
	fmt.Printf("\ninput config %#v\n", apigateway.Endpoint)

	_, err = apigateway.PostToConnection(input)
	if err != nil {
		fmt.Printf("ERROR %#v\n", err)
		log.Println("ERROR", err.Error())
	}

	/////

	//userSocks := make([]UserSocket, *output.Count)
	//fmt.Printf("userSocks %#v\n", userSocks)

	//dynamodbattribute.UnmarshalListOfMaps(output.Items, &userSocks)
	//fmt.Printf("userSocks %#v\n", userSocks)

	//for _, userSock := range userSocks {
	//	input := &apigatewaymanagementapi.PostToConnectionInput{
	//		ConnectionId: aws.String(userSock.ConnectionID),
	//		Data:         data,
	//	}
	//	fmt.Printf("userSock %#v\n", userSock)
	//	fmt.Printf("input %#v\n", input)
	//	fmt.Printf("input data %#v\n", string(input.Data))
	//
	//	fmt.Printf("%s.execute-api.eu-central-1.amazonaws.com/test", websocketApi)
	//	apigateway := apigatewaymanagementapi.New(sess, aws.NewConfig().WithEndpoint(fmt.Sprintf("%s.execute-api.eu-central-1.amazonaws.com/test", websocketApi)))
	//	fmt.Printf("\ninput config %#v\n", apigateway.Endpoint)
	//
	//	_, err = apigateway.PostToConnection(input)
	//	if err != nil {
	//		fmt.Printf("ERROR %#v\n", err)
	//		log.Println("ERROR", err.Error())
	//	}
	//}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
