package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"os"
	"time"
)

func storeMessageToDynamo(tableName string, message MessageWithInfo) error {
	fmt.Println("storigntodynamo")
	msg := MessageSave{}

	msg.UserFromTo = fmt.Sprintf("%s:%s", message.From, message.To)

	// Get current date and attach it to struct
	dt := time.Now()
	msg.Timestamp = dt.Format("01-02-2006 15:04:05")

	// Attach message
	fmt.Printf("Mesgg: %s\n", message.Message.(string))
	msg.Message = message.Message.(string)

	fmt.Printf("tst: %#v\n", msg)

	// Marshall each element to 'map[string]types.AttributeValue' format
	av, err := attributevalue.MarshalMap(msg)
	if err != nil {
		return err
	}

	fmt.Printf("av: %#v\n", av)

	//TODO remake it to aws sdk v2
	region := os.Getenv("REGION")
	c := NewClient(region)

	// Create dynamoDB.go item input
	// Put item in dynamoDB.go
	_, err = c.DynamoCl.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av})

	if err != nil {
		return err
	}

	fmt.Println("ENDstorigntodynamo")

	return nil
}
