package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (c *Client) queryDynamo(hashKey, sortKey map[string]string, tableName string) (dynamodb.QueryOutput, error) {
	var queryInput dynamodb.QueryInput

	if hashKey["value"] == "" && sortKey["value"] == "" {
		return dynamodb.QueryOutput{}, fmt.Errorf("invalid parameters")
	}

	if hashKey["value"] != "" && sortKey["value"] == "" {
		queryInput = dynamodb.QueryInput{
			TableName:              aws.String(tableName),
			KeyConditionExpression: aws.String(hashKey["keyName"] + "= :hashKey"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":hashKey": &types.AttributeValueMemberS{Value: hashKey["value"]},
			},
		}
	}

	if hashKey["value"] != "" && sortKey["value"] != "" {
		queryInput = dynamodb.QueryInput{
			TableName:              aws.String(tableName),
			KeyConditionExpression: aws.String(hashKey["keyName"] + "= :hashKey and #" + sortKey["keyName"] + "= :rangeKey"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":hashKey":  &types.AttributeValueMemberS{Value: hashKey["value"]},
				":rangeKey": &types.AttributeValueMemberS{Value: sortKey["value"]},
			},
			ExpressionAttributeNames: map[string]string{
				"#" + sortKey["keyName"]: sortKey["keyName"]},
		}
	}

	out, err := c.DynamoCl.Query(context.Background(), &queryInput)
	if err != nil {
		return dynamodb.QueryOutput{}, err
	}

	return *out, nil
}
