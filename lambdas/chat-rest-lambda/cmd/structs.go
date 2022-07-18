package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Message struct {
	UserFromTo string `json:"userFromTo"`
	Timestamp  string `json:"timestamp"`
	Message    string `json:"message"`
}

// Parse DynamoDB response to Message list
func parseDynamoToMessage(out dynamodb.QueryOutput) ([]Message, error) {
	var items []Message

	// Unmarshall list of Maps to Message structure
	err := attributevalue.UnmarshalListOfMaps(out.Items, &items)
	if err != nil {
		fmt.Println(err)
		return items, err
	}
	return items, nil
}
