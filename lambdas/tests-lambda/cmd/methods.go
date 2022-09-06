package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"golang.org/x/exp/slices"
	"os"
)

func (c *Client) SaveTestQuestion(request events.APIGatewayProxyRequest) (string, error) {
	tableName := os.Getenv("TESTS_TABLE")

	// Format given request body to LessonSlide struct
	var question TestQuestionInput
	err := json.Unmarshal([]byte(request.Body), &question)
	if err != nil {
		return "", err
	}

	// Save lesson slide to dynamoDB
	err = c.saveToDynamo(question, tableName)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (c *Client) GetTest(request events.APIGatewayProxyRequest) (string, error) {
	tableName := os.Getenv("TESTS_TABLE")

	// If query string empty
	if len(request.QueryStringParameters) == 0 {
		out, err := c.scanDynamo(make([]map[string]string, 0), []string{"TestId"}, tableName)
		if err != nil {
			return "", err
		}

		var tests []TestId
		err = attributevalue.UnmarshalListOfMaps(out.Items, &tests)
		fmt.Printf("%#v\n", tests)

		var uniqueTests []TestId
		for _, elem := range tests {
			if !slices.Contains(uniqueTests, elem) {
				uniqueTests = append(uniqueTests, elem)
			}
		}

		jsonOut, err := json.Marshal(uniqueTests)

		if err != nil {
			return "", err
		}

		return string(jsonOut), nil
	}

	hashKey := map[string]string{"keyName": "TestId", "value": request.QueryStringParameters["testId"]}
	sortKey := map[string]string{}

	result, err := c.queryDynamo(hashKey, sortKey, tableName)

	var resp []TestQuestionOutput
	err = attributevalue.UnmarshalListOfMaps(result.Items, &resp)

	jsonOut, err := json.Marshal(resp)

	if err != nil {
		return "", err
	}

	return string(jsonOut), nil
}

func (c *Client) DeleteQuestion(request events.APIGatewayProxyRequest) (string, error) {
	tableName := os.Getenv("TESTS_TABLE")

	// If query string empty
	if len(request.QueryStringParameters) == 0 {
		return "", fmt.Errorf("not enough parameters\n")
	}

	hashKey := map[string]string{"keyName": "TestId", "value": request.QueryStringParameters["testId"]}
	sortKey := map[string]string{"keyName": "CombinedKey", "value": request.QueryStringParameters["combinedKey"]}

	err := c.deleteFromDynamo(hashKey, sortKey, tableName)
	if err != nil {
		return "", err
	}
	return "", nil
}
