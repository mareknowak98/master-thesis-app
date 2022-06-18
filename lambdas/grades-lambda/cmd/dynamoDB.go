package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"time"
)

// SaveGrade POST method function
// SaveGrade Function save grade to DynamoDB with timestamp
func (c *Client) SaveGrade(request events.APIGatewayProxyRequest, tableName string) error {
	// Format given request body to InputGrades struct
	var grade InputGrades
	err := json.Unmarshal([]byte(request.Body), &grade)
	if err != nil {
		fmt.Println(err)
	}

	// Get current date and attach it to struct
	dt := time.Now()
	grade.Date = dt.Format("01-02-2006 15:04:05")

	// Marshall each element to 'map[string]types.AttributeValue' format
	av, err := attributevalue.MarshalMap(grade)
	if err != nil {
		return err
	}

	// Create dynamoDB item input
	// Put item in dynamoDB
	_, err = c.DynamoCl.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av})

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetGrades(request events.APIGatewayProxyRequest, tableName string) error {
	fmt.Println(request.QueryStringParameters)

	return nil
}
