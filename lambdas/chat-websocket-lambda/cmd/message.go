package cmd

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

// Default ..
func Default(request APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	tableName := os.Getenv("CONNECTIONS_TABLE")
	tableNameMessages := os.Getenv("MESSAGES_TABLE")

	websocketApi := os.Getenv("WEBSOCKET_API")
	deploymentType := os.Getenv("DEPLOYMENT_TYPE")

	b := MessageAction{}
	if err := json.NewDecoder(strings.NewReader(request.Body)).Decode(&b); err != nil {
		log.Println("Unable to decode body", err.Error())
	}

	fmt.Printf("decoded body %s\n", b)

	data, _ := json.Marshal(b)
	fmt.Printf("MessageAction %#v\n", string(data))

	sess := GetSession()
	//fmt.Printf("Session %#v\n", sess)

	fmt.Printf("RequestContext.ConnectionID %#v\n", request.RequestContext.ConnectionID)

	db := dynamodb.New(sess)

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

	err = storeMessageToDynamo(tableNameMessages, b.Payload)
	if err != nil {
		fmt.Printf("ERROR (saving message to dynamodb) %#v\n", err)
		return events.APIGatewayProxyResponse{}, err
	}

	if len(output.Items) > 0 {
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

		// TODO there was issue with IAM
		apigateway := apigatewaymanagementapi.New(sess, aws.NewConfig().WithEndpoint(fmt.Sprintf("%s.execute-api.eu-central-1.amazonaws.com/%s", websocketApi, deploymentType)))
		fmt.Printf("\ninput config %#v\n", apigateway.Endpoint)

		_, err = apigateway.PostToConnection(input)
		if err != nil {
			fmt.Printf("ERROR %#v\n", err)
			log.Println("ERROR", err.Error())
		}
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
