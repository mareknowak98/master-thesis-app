package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

// GetClasses Function return list of classes
func (c *Client) GetClasses(request events.APIGatewayProxyRequest) (string, error) {
	tableName := os.Getenv("CLASS_TABLE")

	scanOutput, err := c.DynamoCl.Scan(context.Background(), &dynamodb.ScanInput{
		TableName:       aws.String(tableName),
		AttributesToGet: []string{"UserClass"},
	})

	if err != nil {
		return "", err
	}

	var classL []map[string]string
	//Unmarshall response to list of maps
	err = attributevalue.UnmarshalListOfMaps(scanOutput.Items, &classL)

	var classList []string
	for _, element := range classL {
		classList = append(classList, element["UserClass"])
	}

	//Marshall to json format
	jsonOut, err := json.Marshal(classList)

	if err != nil {
		return "", err
	}

	return string(jsonOut), nil
}

// SetClasses Function add new class to table
func (c *Client) SetClasses(request events.APIGatewayProxyRequest) (string, error) {
	var class AddClassInput
	tableName := os.Getenv("CLASS_TABLE")
	err := json.Unmarshal([]byte(request.Body), &class)
	if err != nil {
		return "", err
	}
	// Marshall each element to 'map[string]types.AttributeValue' format
	av, err := attributevalue.MarshalMap(class)
	if err != nil {
		return "", err
	}

	_, err = c.DynamoCl.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})

	if err != nil {
		return "", err
	}

	return "OK", nil
}

// DeleteClasses Function delete class from table
func (c *Client) DeleteClasses(request events.APIGatewayProxyRequest) (string, error) {
	var class AddClassInput
	tableName := os.Getenv("CLASS_TABLE")
	err := json.Unmarshal([]byte(request.Body), &class)
	if err != nil {
		return "", err
	}

	// Marshall each element to 'map[string]types.AttributeValue' format
	av, err := attributevalue.MarshalMap(class)
	if err != nil {
		return "", err
	}

	_, err = c.DynamoCl.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key:       av,
	})

	return "OK", nil
}

// QueryDynamo Query dynamoDB.go only with hashkey
func (c *Client) QueryDynamo(context context.Context, tableName string, hashKeyValue string) (dynamodb.QueryOutput, error) {
	out, err := c.DynamoCl.Query(context, &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("UserClass = :hashKey"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":hashKey": &types.AttributeValueMemberS{Value: hashKeyValue},
		},
	})
	if err != nil {
		return *out, err
	}

	return *out, nil
}

//TODO use UpdateItem
func (c *Client) AddUserToClass(request events.APIGatewayProxyRequest) (string, error) {
	var input AddUserInput
	tableName := os.Getenv("CLASS_TABLE")
	err := json.Unmarshal([]byte(request.Body), &input)
	if err != nil {
		return "", err
	}

	currentClassMembers, err := c.QueryDynamo(context.Background(), tableName, input.UserClass)
	if err != nil {
		return "", err
	}

	var userList []string
	//Unmarshall response to list
	err = attributevalue.Unmarshal(currentClassMembers.Items[0]["UserList"], &userList)
	if err != nil {
		return "", err
	}
	userList = append(userList, input.Username)
	var dynamoRecord ClassRecord
	dynamoRecord.UserClass = input.UserClass
	dynamoRecord.UserList = userList
	// Marshall each element to 'map[string]types.AttributeValue' format
	av, err := attributevalue.MarshalMap(dynamoRecord)
	if err != nil {
		return "", err
	}

	_, err = c.DynamoCl.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})

	return "OK", nil
}

//TODO use UpdateItem
func (c *Client) DeleteUserFromClass(request events.APIGatewayProxyRequest) (string, error) {
	var input AddUserInput
	tableName := os.Getenv("CLASS_TABLE")
	err := json.Unmarshal([]byte(request.Body), &input)
	if err != nil {
		return "", err
	}

	currentClassMembers, err := c.QueryDynamo(context.Background(), tableName, input.UserClass)
	if err != nil {
		return "", err
	}

	var userList []string
	//Unmarshall response to list
	err = attributevalue.Unmarshal(currentClassMembers.Items[0]["UserList"], &userList)
	if err != nil {
		return "", err
	}

	// Find and remove "two"
	for i, v := range userList {
		if v == input.Username {
			userList = append(userList[:i], userList[i+1:]...)
			break
		}
	}

	var dynamoRecord ClassRecord
	dynamoRecord.UserClass = input.UserClass
	dynamoRecord.UserList = userList
	// Marshall each element to 'map[string]types.AttributeValue' format
	av, err := attributevalue.MarshalMap(dynamoRecord)
	if err != nil {
		return "", err
	}

	_, err = c.DynamoCl.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})

	return "OK", nil
}
