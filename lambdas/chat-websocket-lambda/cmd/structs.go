package cmd

import "github.com/aws/aws-lambda-go/events"

type MessageAction struct {
	Type    string          `json:"type"`
	Payload MessageWithInfo `json:"payload"`
}

// MessageWithInfo ..
type MessageWithInfo struct {
	To      string      `json:"to"`
	From    string      `json:"from"`
	Message interface{} `json:"message"`
}

type MessageSave struct {
	UserFromTo string `json:"userFromTo"`
	Timestamp  string `json:"timestamp"`
	Message    string `json:"message"`
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
