package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// save any item to given dynamodb table
func (c *Client) saveToDynamo(item interface{}, tableName string) error {
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = c.DynamoCl.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) deleteFromDynamo(hashKey, sortKey map[string]string, tableName string) error {
	if hashKey["value"] == "" || sortKey["value"] == "" {
		return fmt.Errorf("invalid parameters")
	}

	_, err := c.DynamoCl.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			hashKey["keyName"]: &types.AttributeValueMemberS{Value: hashKey["value"]},
			sortKey["keyName"]: &types.AttributeValueMemberS{Value: sortKey["value"]},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

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

func (c *Client) scanDynamo(scanKeys []map[string]string, attributesToGet []string, tableName string) (dynamodb.ScanOutput, error) {
	var scanInput dynamodb.ScanInput

	scanInput = dynamodb.ScanInput{
		TableName:       aws.String(tableName),
		AttributesToGet: attributesToGet,
	}

	out, err := c.DynamoCl.Scan(context.Background(), &scanInput)
	if err != nil {
		return dynamodb.ScanOutput{}, err
	}

	return *out, nil
}
