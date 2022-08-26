package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"time"
)

// SaveGrade POST method function
// SaveGrade Function save grade to DynamoDB with timestamp
func (c *Client) SaveGrade(request events.APIGatewayProxyRequest, tableName string) error {
	// Format given request body to InputGrades struct
	var grade InputGrades
	err := json.Unmarshal([]byte(request.Body), &grade)
	if err != nil {
		return err
	}

	// Get current date and attach it to struct
	dt := time.Now()
	grade.Date = dt.Format("01-02-2006 15:04:05")

	// Marshall each element to 'map[string]types.AttributeValue' format
	av, err := attributevalue.MarshalMap(grade)
	if err != nil {
		return err
	}

	// Create dynamoDB.go item input
	// Put item in dynamoDB.go
	_, err = c.DynamoCl.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av})

	if err != nil {
		return err
	}

	return nil
}

// QueryDynamo Query dynamoDB.go only with hashkey
func (c *Client) QueryDynamo(context context.Context, tableName string, hashKeyValue string) (dynamodb.QueryOutput, error) {
	out, err := c.DynamoCl.Query(context, &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("UserId = :hashKey"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":hashKey": &types.AttributeValueMemberS{Value: hashKeyValue},
		},
	})
	if err != nil {
		return *out, err
	}

	return *out, nil
}

// Parse DynamoDB response to json
func parseDynamoToInputGrades(out dynamodb.QueryOutput) (string, error) {
	var items []InputGrades

	// Unmarshall list of Maps to InputGrades structure
	err := attributevalue.UnmarshalListOfMaps(out.Items, &items)

	//Marshall to json format
	jsonOut, err := json.Marshal(items)

	if err != nil {
		return "", err
	}
	return string(jsonOut), nil
}

// GetGrades This function fetches grades from grades table
func (c *Client) GetGrades(request events.APIGatewayProxyRequest, tableName string) (string, error) {
	// If query string empty
	if len(request.QueryStringParameters) == 0 {
		return "", fmt.Errorf("not enough parameters\n")
	}

	//Get by UserId
	out, err := c.QueryDynamo(context.Background(), tableName, request.QueryStringParameters["UserId"])
	if err != nil {
		return "", err
	}

	//Parse output to json
	jsonOut, err := parseDynamoToInputGrades(out)
	if err != nil {
		return "", err
	}

	return jsonOut, nil
}
