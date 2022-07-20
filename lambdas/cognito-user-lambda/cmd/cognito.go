package cmd

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (c *Client) getUserToken(login UserLogin, clientId string) (UserLoginResponse, error) {
	resp := UserLoginResponse{}

	// convert UserLogin struct to map[string]string
	var authMap map[string]string
	data, err := json.Marshal(login)
	if err != nil {
		return resp, nil
	}

	err = json.Unmarshal(data, &authMap)
	if err != nil {
		return resp, nil
	}

	// call to cognito user pools
	auth, err := c.CognitoCl.InitiateAuth(context.Background(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       "USER_PASSWORD_AUTH",
		ClientId:       aws.String(clientId),
		AuthParameters: authMap,
	})
	if err != nil {
		return resp, err
	}

	resp.AccessToken = *auth.AuthenticationResult.AccessToken
	resp.RefreshToken = *auth.AuthenticationResult.RefreshToken

	return resp, nil
}
