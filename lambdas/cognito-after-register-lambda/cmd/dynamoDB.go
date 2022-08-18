package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"time"
)

// SaveUser This function is triggered each time when new user is created in Cognito
// Parse Cognito request body and save new user to dynamoDB.go
func (c *Client) SaveUser(request CognitoEventRequest, tableName string) error {
	// FInitialize InputUsers which will be saved into DynamoDB
	var user InputUsers

	// Fill user initial data
	user.UserId = request.Request["userAttributes"]["sub"]
	user.Username = request.Username
	user.Email = request.Request["userAttributes"]["email"]
	user.UserGroup = "student-group"
	user.UserGroup = ""

	// Get current date and attach it to struct
	dt := time.Now()
	user.CreatedAt = dt.Format("01-02-2006 15:04:05")

	// Marshall each element to 'map[string]types.AttributeValue' format
	av, err := attributevalue.MarshalMap(user)
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
