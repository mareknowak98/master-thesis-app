package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"os"
)

func (c *Client) GetUsers(request events.APIGatewayProxyRequest, tableName string) (string, error) {
	if len(request.QueryStringParameters) == 0 {
		return "", fmt.Errorf("not enough parameters\n")
	}

	if request.QueryStringParameters["group"] == "" {
		return "", fmt.Errorf("bad query string parameter\n")
	}

	if request.QueryStringParameters["group"] == "all" {
		resp, err := c.GetAllUsers()
		if err != nil {
			return "", err
		}
		return resp, nil
	}

	resp, err := c.GetGroupUsers(request.QueryStringParameters["group"])
	if err != nil {
		return "", err
	}
	return resp, nil
}

func (c *Client) GetAllUsers() (string, error) {
	var userList []OutputUsers

	userPoolID := os.Getenv("USER_POOL_ID")

	usersInGroup, err := c.CognitoCl.ListUsers(context.Background(), &cognitoidentityprovider.ListUsersInput{
		UserPoolId: aws.String(userPoolID),
	})

	if err != nil {
		return "", err
	}

	for _, element := range usersInGroup.Users {
		var tmp OutputUsers
		tmp.Username = *element.Username
		tmp.Email = *(element.Attributes[2].Value)
		userList = append(userList, tmp)
	}

	//Marshall to json format
	jsonOut, err := json.Marshal(userList)

	if err != nil {
		return "", err
	}
	return string(jsonOut), nil
}

func (c *Client) GetGroupUsers(groupName string) (string, error) {
	var userList []OutputUsers
	userPoolID := os.Getenv("USER_POOL_ID")

	usersInGroup, err := c.CognitoCl.ListUsersInGroup(context.Background(), &cognitoidentityprovider.ListUsersInGroupInput{
		GroupName:  aws.String(groupName),
		UserPoolId: aws.String(userPoolID),
	})

	if err != nil {
		return "", err
	}

	for _, element := range usersInGroup.Users {
		var tmp OutputUsers
		tmp.Username = *element.Username
		tmp.Email = *(element.Attributes[2].Value)
		userList = append(userList, tmp)
	}

	//Marshall to json format
	jsonOut, err := json.Marshal(userList)

	if err != nil {
		return "", err
	}
	return string(jsonOut), nil
}

//func (c *Client) GetUsers(request events.APIGatewayProxyRequest, tableName string) (string, error) {
//	fmt.Println(request.QueryStringParameters)
//
//	users, err := c.getAllItems(context.Background(), tableName)
//	// Parse response to Json
//	resp, err := parseDynamoToUsers(users)
//	if err != nil {
//		return "", err
//	}
//	return resp, nil
//}
//
//// Scan all items in DynamoDB table
//// WARNING: can be very slow, do not use on large tables
//// TODO: Add batching (max limit of scan output is 1MB)
//func (c *Client) getAllItems(context context.Context, tableName string) (dynamodb.ScanOutput, error) {
//	out, err := c.DynamoCl.Scan(context, &dynamodb.ScanInput{
//		TableName: aws.String(tableName),
//	})
//	if err != nil {
//		return dynamodb.ScanOutput{}, err
//	}
//
//	return *out, nil
//}
//
//// Parse DynamoDB response to json
//func parseDynamoToUsers(out dynamodb.ScanOutput) (string, error) {
//	var items []OutputUsers
//
//	// Unmarshall list of Maps to OutputUsers structure
//	err := attributevalue.UnmarshalListOfMaps(out.Items, &items)
//
//	//Marshall to json format
//	jsonOut, err := json.Marshal(items)
//
//	if err != nil {
//		return "", err
//	}
//	return string(jsonOut), nil
//}

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
