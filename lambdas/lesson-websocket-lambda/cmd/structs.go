package cmd

import "github.com/aws/aws-lambda-go/events"

type LessonSlide struct {
	LessonId        string `json:"lessonId"`
	SlideId         string `json:"slideId"`
	SlideType       string `json:"slideType"`
	SlideContent    string `json:"slideContent"`
	QuestionAnswers string `json:"questionAnswers"`
}

type SlideSwap struct {
	LessonId string `json:"lessonId"`
	SlideId  string `json:"slideId"`
}

type LessonSocket struct {
	LessonId     string `json:"LessonId"`
	ConnectionID string `json:"ConnectionID"`
}

// APIGatewayWebsocketProxyRequest modified a little from default struct from package
type APIGatewayWebsocketProxyRequest struct {
	MethodArn                       string                                        `json:"methodArn"`
	Resource                        string                                        `json:"resource"` // The resource path defined in API Gateway
	Path                            string                                        `json:"path"`     // The url path for the caller
	HTTPMethod                      string                                        `json:"httpMethod"`
	Headers                         map[string]string                             `json:"headers"`
	MultiValueHeaders               map[string][]string                           `json:"multiValueHeaders"`
	QueryStringParameters           map[string]string                             `json:"queryStringParameters"`
	MultiValueQueryStringParameters map[string][]string                           `json:"multiValueQueryStringParameters"`
	PathParameters                  map[string]string                             `json:"pathParameters"`
	StageVariables                  map[string]string                             `json:"stageVariables"`
	RequestContext                  events.APIGatewayWebsocketProxyRequestContext `json:"requestContext"`
	Body                            string                                        `json:"body"`
	IsBase64Encoded                 bool                                          `json:"isBase64Encoded,omitempty"`
}
