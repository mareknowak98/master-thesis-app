package cmd

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
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

// DeleteLessonOrSlide function delete given slide or all lesson
func (c *Client) DeleteLessonOrSlide(request events.APIGatewayProxyRequest, tableName string) (string, error) {

	return "", nil
}

// UpdateLessonSlide function update given slide from lesson
func (c *Client) UpdateLessonSlide(request events.APIGatewayProxyRequest, tableName string) (string, error) {

	return "", nil
}

// GetLessons function return list of lessons/ details of given lesson/slides from given lesson
func (c *Client) GetLessons(request events.APIGatewayProxyRequest, tableName string) (string, error) {

	return "", nil
}
