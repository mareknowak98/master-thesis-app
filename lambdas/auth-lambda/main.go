package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-jwt/jwt/v4"
	"mylearnproject/lambdas/auth-lambda/cmd"
)

func HandleRequest(ctx context.Context, request events.APIGatewayV2CustomAuthorizerV1Request) (events.APIGatewayCustomAuthorizerResponse, error) {
	//var res events.APIGatewayCustomAuthorizerResponse
	fmt.Printf("url: %#v\n", cmd.MethodArnToUrl("arn:aws:execute-api:eu-central-1:614695401085:y1llo92zmi/test/GET/grades"))

	fmt.Printf("Request: %#v\n", request)

	parsedToken, err := cmd.DecodeToken(request.AuthorizationToken)
	if err != nil {
		return authPolicyResponse("user", "Deny", request.MethodArn, nil), nil
	}

	// check if claims are present and token is valid.
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		// return Allow authResponse with userEntity in authorizer context for next lambda in chain.
		return authPolicyResponse("user", "Allow", request.MethodArn, map[string]interface{}{"userEntity": claims["sub"].(string)}), nil
	}

	return authPolicyResponse("user", "Deny", request.MethodArn, nil), nil
}

// authPolicyResponse for authorization/ de-authorization of given principal with supplied context
func authPolicyResponse(principalID, effect, resource string, context map[string]interface{}) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalID}
	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}
	authResponse.Context = context
	return authResponse
}

func main() {
	lambda.Start(HandleRequest)
}
