package cmd

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (c *Client) GetUsers(request events.APIGatewayProxyRequest, tableName string) (string, error) {
	users, err := c.getAllItems(context.Background(), tableName)
	// Parse response to Json
	resp, err := parseDynamoToUsers(users)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// Scan all items in DynamoDB table
// WARNING: can be very slow, do not use on large tables
// TODO: Add batching (max limit of scan output is 1MB)
func (c *Client) getAllItems(context context.Context, tableName string) (dynamodb.ScanOutput, error) {
	out, err := c.DynamoCl.Scan(context, &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		return dynamodb.ScanOutput{}, err
	}

	return *out, nil
}

// Parse DynamoDB response to json
func parseDynamoToUsers(out dynamodb.ScanOutput) (string, error) {
	var items []OutputUsers

	// Unmarshall list of Maps to OutputUsers structure
	err := attributevalue.UnmarshalListOfMaps(out.Items, &items)

	//Marshall to json format
	jsonOut, err := json.Marshal(items)

	if err != nil {
		return "", err
	}
	return string(jsonOut), nil
}

// TODO for later
//// QueryDynamo Query dynamoDB.go only with hashkey
//func (c *Client) QueryDynamo(context context.Context, tableName string, hashKeyName string, hashKeyValue string) (dynamodb.QueryOutput, error) {
//	out, err := c.DynamoCl.Query(context, &dynamodb.QueryInput{
//		TableName:              aws.String(tableName),
//		KeyConditionExpression: aws.String(fmt.Sprintf("%s = :hashKey", hashKeyName)),
//		ExpressionAttributeValues: map[string]types.AttributeValue{
//			":hashKey": &types.AttributeValueMemberS{Value: hashKeyValue},
//		},
//	})
//	if err != nil {
//		return *out, err
//	}
//
//	return *out, nil
//}
