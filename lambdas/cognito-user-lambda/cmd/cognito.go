package cmd

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (c *Client) getUserToken(login UserLogin, clientId string) (string, error) {
	// convert UserLogin struct to map[string]string
	var authMap map[string]string
	data, err := json.Marshal(login)
	if err != nil {
		return "", nil
	}

	err = json.Unmarshal(data, &authMap)
	if err != nil {
		return "", nil
	}

	// call to cognito user pools
	resp, err := c.CognitoCl.InitiateAuth(context.Background(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       "USER_PASSWORD_AUTH",
		ClientId:       aws.String(clientId),
		AuthParameters: authMap,
	})
	if err != nil {
		return "", err
	}
	return *resp.AuthenticationResult.AccessToken, nil
}
