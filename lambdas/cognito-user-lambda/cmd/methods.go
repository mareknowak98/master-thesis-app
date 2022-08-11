package cmd

import (
	"encoding/json"
	"fmt"
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

	data, _ := json.Marshal(token)
	return string(data), nil
}

func (c *Client) Register(request events.APIGatewayProxyRequest) (string, error) {
	fmt.Println("register")
	// Format given request body to UserRegister struct
	var register UserRegister
	err := json.Unmarshal([]byte(request.Body), &register)
	if err != nil {
		return "", err
	}
	fmt.Printf("%#v\n", register)

	cognitoClientId := os.Getenv("COGNITO_CLIENT_ID")
	userPoolID := os.Getenv("USER_POOL_ID")
	isRegistered, err := c.registerUser(register, cognitoClientId, userPoolID)
	if err != nil {
		return "", err
	}

	data, _ := json.Marshal(isRegistered)

	return string(data), nil
}
