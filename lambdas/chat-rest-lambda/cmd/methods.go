package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func (c *Client) GetMessagesWithUser(request events.APIGatewayProxyRequest, tableName string) (string, error) {
	decodedToken, err := DecodeToken(request.Headers["Authorization"])
	if err != nil {
		return "Token malformed", err
	}

	//User sending request
	requestUser := decodedToken["username"]
	//User with whom we check messages
	consumerUser := request.QueryStringParameters["UserTo"]

	//TODO possibly can be multithreaded as lambda with 128mb has 2 cCPUs
	messages1, err := c.QueryDynamoHK(context.Background(), tableName, "UserFromTo", fmt.Sprintf("%s:%s", requestUser, consumerUser))
	if err != nil {
		return "Query error", err
	}

	messages2, err := c.QueryDynamoHK(context.Background(), tableName, "UserFromTo", fmt.Sprintf("%s:%s", consumerUser, requestUser))
	if err != nil {
		return "Query error", err
	}

	// parse dynamo query output to Message object
	parsedMessages1, _ := parseDynamoToMessage(messages1)
	parsedMessages2, _ := parseDynamoToMessage(messages2)

	// Append slice to slice
	parsedMessages1 = append(parsedMessages1, parsedMessages2...)

	//Marshall to json format
	jsonOut, err := json.Marshal(parsedMessages1)

	if err != nil {
		return "", err
	}
	return string(jsonOut), nil
}
