package cmd

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"os"
)

func (c *Client) Login(request events.APIGatewayProxyRequest) (string, error) {
	// Format given request body to UserLogin struct
	var login UserLogin
	err := json.Unmarshal([]byte(request.Body), &login)
	if err != nil {
		return "", err
	}

	cognitoClientId := os.Getenv("COGNITO_CLIENT_ID")

	// Make call to Cognito user pools to retrieve token
	token, err := c.getUserToken(login, cognitoClientId)
	if err != nil {
		return "", err
	}
	return token, nil
}
