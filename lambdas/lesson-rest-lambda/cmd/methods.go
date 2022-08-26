package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

// SaveLessonSlide function parse request message to LessonSlide struct and save record to dynamoDB
func (c *Client) SaveLessonSlide(request events.APIGatewayProxyRequest, tableName string) (string, error) {
	// Format given request body to LessonSlide struct
	var slide LessonSlide
	err := json.Unmarshal([]byte(request.Body), &slide)
	if err != nil {
		return "", err
	}

	// Save lesson slide to dynamoDB
	err = c.saveToDynamo(slide, tableName)
	if err != nil {
		return "", err
	}
	return "", nil
}

// DeleteSlide function delete given slide
func (c *Client) DeleteSlide(request events.APIGatewayProxyRequest, tableName string) (string, error) {
	// If query string empty
	if len(request.QueryStringParameters) == 0 {
		return "", fmt.Errorf("not enough parameters\n")
	}

	hashKey := map[string]string{"keyName": "LessonId", "value": request.QueryStringParameters["lesson"]}
	sortKey := map[string]string{"keyName": "SlideId", "value": request.QueryStringParameters["slide"]}

	err := c.deleteFromDynamo(hashKey, sortKey, tableName)
	if err != nil {
		return "", err
	}
	return "", nil
}

// GetSlide function return given slide
func (c *Client) GetSlide(request events.APIGatewayProxyRequest, tableName string) (string, error) {
	// If query string empty
	if len(request.QueryStringParameters) == 0 {
		return "", fmt.Errorf("not enough parameters\n")
	}

	hashKey := map[string]string{"keyName": "LessonId", "value": request.QueryStringParameters["lesson"]}
	sortKey := map[string]string{"keyName": "SlideId", "value": request.QueryStringParameters["slide"]}

	out, err := c.queryDynamo(hashKey, sortKey, tableName)

	var resp []LessonSlide
	err = attributevalue.UnmarshalListOfMaps(out.Items, &resp)

	jsonOut, err := json.Marshal(resp)

	if err != nil {
		return "", err
	}

	return string(jsonOut), nil
}

// GetLessons function return list of lessons/ details of given lesson/slides from given lesson
func (c *Client) GetLessons(request events.APIGatewayProxyRequest, tableName string) (string, error) {
	// If query string empty
	if len(request.QueryStringParameters) == 0 {
		out, err := c.scanDynamo(make([]map[string]string, 0), []string{"LessonId"}, tableName)
		if err != nil {
			return "", err
		}

		var resp []map[string]string

		err = attributevalue.UnmarshalListOfMaps(out.Items, &resp)
		uniqueResp := unique(resp, "LessonId")

		jsonOut, err := json.Marshal(uniqueResp)

		if err != nil {
			return "", err
		}

		return string(jsonOut), nil
	}

	return "", fmt.Errorf("bad parameters\n")
}

func unique(interfaceSlice []map[string]string, key string) map[string]string {
	var resp = map[string]string{}
	for _, v := range interfaceSlice {
		resp[v[key]] = key
	}

	return resp
}

//	for _, v := range out.Items {
//		var esa LessonSlide
//		err = attributevalue.UnmarshalMap(v, &esa)
//		fmt.Printf("%#v\n", esa)
//	}
