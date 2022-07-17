package main

import (
	"github.com/aws/aws-lambda-go/events"
	"os"
	"testing"
)

//
//import (
//	"fmt"
//	"github.com/aws/aws-lambda-go/events"
//	"os"
//	"testing"
//)
//
// File to test running lambda locally

func TestHandleRequest(t *testing.T) {
	_ = os.Setenv("MESSAGES_TABLE", "chat-connections")
	_ = os.Setenv("WEBSOKET_API", "rjyithf86i")

	tmp := APIGatewayWebsocketProxyRequest{}
	tmp.RequestContext.ConnectionID = "Va_3vexLFiACHxA="
	tmp.Body = "{\n   \"type\" : \"message\",\n   \"payload\" : {\n        \"from\" : \"oneuser\",\n        \"to\" : \"testuser2\",\n        \"message\" : \"something\"\n   }\n}"

	Default(tmp)

}

//func TestHandleRequest(t *testing.T) {
//	_ = os.Setenv("MESSAGES_TABLE", "chat-connections")
//
//	//tableName := os.Getenv("MESSAGES_TABLE")
//
//	tmp := APIGatewayWebsocketProxyRequest{}
//	tmp.RequestContext.ConnectionID = "Va43ccv1FiACFFg="
//	Disconnect(tmp)
//}

//
//func TestHandleRequest(t *testing.T) {
//
//	//claims := jwt.MapClaims{"username": "testuser2"}
//	//fmt.Printf("token claims: %#v\n", claims)
//	////id1 := claims["cognito:username"].(interface{}).(string)
//	//id2 := claims["username"].(string)
//	////fmt.Println(id1)
//	//fmt.Println(id2)
//
//	b := MessageAction{}
//	if err := json.NewDecoder(strings.NewReader("{\n   \"type\" : \"message\",\n   \"payload\" : {\n        \"from\" : \"oneuser\",\n        \"to\" : \"otheruser\",\n        \"message\" : \"something\"\n   }\n}")).Decode(&b); err != nil {
//		fmt.Println(err)
//	}
//
//	data, _ := json.Marshal(b)
//	fmt.Printf("MessageAction %#v\n", string(data))
//
//	fmt.Printf("decoded body %#v\n", b)
//	_ = os.Setenv("MESSAGES_TABLE", "chat-connections")
//
//	tableName := os.Getenv("MESSAGES_TABLE")
//
//	sess := GetSession()
//	//fmt.Printf("Session %#v\n", sess)
//
//	db := dynamodb.New(sess)
//
//	queryInput := &dynamodb.QueryInput{
//		TableName: aws.String(tableName),
//		ExpressionAttributeNames: map[string]*string{
//			//"#id":           aws.String("id"),
//			"#connectionID": aws.String("connectionID"),
//		},
//		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
//			":connectionID": &dynamodb.AttributeValue{
//				S: aws.String("Va0-acZgliACEuQ="),
//			},
//			//":id": &dynamodb.AttributeValue{
//			//	S: aws.String(b.Payload.To),
//			//	S: aws.String("id"),
//			//},
//		},
//		KeyConditionExpression: aws.String("#connectionID=:connectionID"),
//	}
//
//	fmt.Printf("QueryInput %#v\n", queryInput)
//
//	output, err := db.Query(queryInput)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Printf("output %#v\n", output)
//
//	userSocks := make([]UserSocket, *output.Count)
//	fmt.Printf("userSocks %#v\n", userSocks)
//
//	dynamodbattribute.UnmarshalListOfMaps(output.Items, &userSocks)
//	fmt.Printf("userSocks %#v\n", userSocks)
//
//	for _, userSock := range userSocks {
//		input := &apigatewaymanagementapi.PostToConnectionInput{
//			ConnectionId: aws.String(userSock.ConnectionID),
//			Data:         data,
//		}
//		fmt.Printf("userSock %#v\n", userSock)
//		fmt.Printf("input %#v\n", input)
//		fmt.Printf("input data %#v\n", string(input.Data))
//
//		apigateway := apigatewaymanagementapi.New(sess, aws.NewConfig().WithEndpoint("rjyithf86i.execute-api.eu-central-1.amazonaws.com/test"))
//		fmt.Printf("input config %#v\n", apigateway.Endpoint)
//
//		_, err = apigateway.PostToConnection(input)
//		if err != nil {
//			fmt.Printf("ERROR %#v\n", err)
//
//			log.Println("ERROR", err.Error())
//		}
//	}
//}

type UserSocket2 struct {
	id           string `json:"id"`
	connectionID string `json:"connectionID"`
}

func getDebugInput() events.APIGatewayProxyRequest {
	err := os.Setenv("REGION", "eu-central-1")
	if err != nil {
		return events.APIGatewayProxyRequest{}
	}
	err = os.Setenv("MESSAGES_TABLE", "cognito-users")
	if err != nil {
		return events.APIGatewayProxyRequest{}
	}

	var request events.APIGatewayProxyRequest

	return request
}
