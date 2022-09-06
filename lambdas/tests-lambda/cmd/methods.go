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

func (c *Client) CheckTest(request events.APIGatewayProxyRequest) (string, error) {
	testsTableName := os.Getenv("TESTS_TABLE")
	resultsTableName := os.Getenv("RESULTS_TABLE")

	// Format given request body to LessonSlide struct
	var testAnswer CheckTestInput
	err := json.Unmarshal([]byte(request.Body), &testAnswer)
	if err != nil {
		return "", err
	}

	hashKey := map[string]string{"keyName": "TestId", "value": testAnswer.TestId}
	sortKey := map[string]string{}

	result, err := c.queryDynamo(hashKey, sortKey, testsTableName)

	var testQuestions []TestQuestionInput

	// Unmarshall list of Maps to InputGrades structure
	err = attributevalue.UnmarshalListOfMaps(result.Items, &testQuestions)
	if err != nil {
		return "", err
	}

	numberCorrect := 0

	for _, elem := range testQuestions {
		var tmpAnswer string
		for _, a := range testAnswer.QuestionAnswer {
			if a.CombinedKey == elem.CombinedKey {
				tmpAnswer = a.CorrectQuestionAnswer
			}
		}

		if tmpAnswer == elem.CorrectAnswer {
			numberCorrect++
		}
	}

	decodedToken, err := DecodeToken(request.Headers["Authorization"])
	if err != nil {
		return "Token malformed", err
	}
	//User sending request
	requestUser := decodedToken["username"]

	var dbInput TestResult
	dbInput.TestId = testAnswer.TestId
	tmp := float32(len(testQuestions)) / float32(numberCorrect)
	dbInput.Result = fmt.Sprintf("%f", tmp)
	dbInput.UserId = requestUser.(string)

	// Save lesson slide to dynamoDB
	err = c.saveToDynamo(dbInput, resultsTableName)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (c *Client) GetResults(request events.APIGatewayProxyRequest) (string, error) {
	resultsTableName := os.Getenv("RESULTS_TABLE")

	if len(request.QueryStringParameters) == 0 {
		return "", fmt.Errorf("not enough parameters\n")
	}

	if request.QueryStringParameters["userId"] != "" && request.QueryStringParameters["testId"] != "" {
		hashKey := map[string]string{"keyName": "TestId", "value": request.QueryStringParameters["testId"]}
		sortKey := map[string]string{"keyName": "UserId", "value": request.QueryStringParameters["userId"]}

		result, err := c.queryDynamo(hashKey, sortKey, resultsTableName)

		var resp []TestResult
		err = attributevalue.UnmarshalListOfMaps(result.Items, &resp)

		jsonOut, err := json.Marshal(resp)

		if err != nil {
			return "", err
		}

		return string(jsonOut), nil
	}

	if request.QueryStringParameters["testId"] != "" {
		hashKey := map[string]string{"keyName": "TestId", "value": request.QueryStringParameters["testId"]}
		sortKey := map[string]string{}

		result, err := c.queryDynamo(hashKey, sortKey, resultsTableName)

		var resp []TestResult
		err = attributevalue.UnmarshalListOfMaps(result.Items, &resp)

		jsonOut, err := json.Marshal(resp)

		if err != nil {
			return "", err
		}

		return string(jsonOut), nil
	}

	return "", nil
}
