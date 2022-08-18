package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type OutputUsers struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
}

type ClassRecord struct {
	UserClass string   `json:"userClass"`
	UserList  []string `json:"serList"`
}

type AddClassInput struct {
	UserClass string `json:"userClass"`
}

type AddUserInput struct {
	UserClass string `json:"userClass"`
	Username  string `json:"username"`
}

type InputUsers struct {
	UserId    string `json:"UserId"`
	Username  string `json:"Username"`
	Email     string `json:"Email"`
	CreatedAt string `json:"CreatedAt"`
	UserGroup string `json:"UserGroup"`
	UserClass string `json:"UserClass"`
}

// Parse DynamoDB response to InputUsers
func parseDynamoToInputUsers(out dynamodb.QueryOutput) ([]InputUsers, error) {
	var items []InputUsers

	// Unmarshall list of Maps to Message structure
	err := attributevalue.UnmarshalListOfMaps(out.Items, &items)
	if err != nil {
		fmt.Println(err)
		return items, err
	}
	return items, nil
}
