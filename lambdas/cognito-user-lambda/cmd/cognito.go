package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
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

func (c *Client) registerUser(register UserRegister, clientId, userPoolId string) (UserRegisterResponse, error) {
	resp := UserRegisterResponse{}
	if register.Password1 != register.Password2 {
		return resp, fmt.Errorf("Passwords are not equal")
	}

	// create user
	isRegistered, err := c.CognitoCl.SignUp(context.Background(), &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(clientId),
		Username: aws.String(register.Username),
		Password: aws.String(register.Password1),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(register.Email),
			},
		},
	})
	fmt.Printf("%#v\n", isRegistered)
	if err != nil {
		return resp, err
	}

	// confirm email automatically : TEMPORARY
	confirm, err := c.CognitoCl.AdminConfirmSignUp(context.Background(), &cognitoidentityprovider.AdminConfirmSignUpInput{
		UserPoolId: aws.String(userPoolId),
		Username:   aws.String(register.Username),
	})
	fmt.Printf("%#v\n", confirm)
	if err != nil {
		return resp, err
	}

	// add new user to default student group
	group, err := c.CognitoCl.AdminAddUserToGroup(context.Background(), &cognitoidentityprovider.AdminAddUserToGroupInput{
		UserPoolId: aws.String(userPoolId),
		Username:   aws.String(register.Username),
		GroupName:  aws.String("student-group"),
	})
	fmt.Printf("%#v\n", group)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
