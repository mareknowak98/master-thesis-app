package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// QueryDynamoHK QueryDynamo Query dynamoDB only with hashkey
func (c *Client) QueryDynamoHK(context context.Context, tableName string, hashKeyName string, hashKeyValue string) (dynamodb.QueryOutput, error) {
	out, err := c.DynamoCl.Query(context, &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String(fmt.Sprintf("%s = :hashKey", hashKeyName)),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":hashKey": &types.AttributeValueMemberS{Value: hashKeyValue},
		},
	})
	if err != nil {
		return *out, err
	}

	return *out, nil
}
