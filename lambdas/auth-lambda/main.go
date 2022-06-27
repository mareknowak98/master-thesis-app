package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-jwt/jwt/v4"
	"mylearnproject/lambdas/auth-lambda/cmd"
)

// HandleRequest Handle lambda request (from api gateway)
func HandleRequest(ctx context.Context, request events.APIGatewayV2CustomAuthorizerV1Request) (events.APIGatewayCustomAuthorizerResponse, error) {
	fmt.Println(request)
	endpointUrl := cmd.MethodArnToUrl(request.MethodArn)
	parsedToken, err := cmd.DecodeToken(request.AuthorizationToken)

	// Return deny if token cannot be parsed
	if err != nil {
		return cmd.AuthPolicyResponse("user", "Deny", request.MethodArn, nil), nil
	}

	//get user group from token claims, then map []interface {} to string
	userGroup := parsedToken.Claims.(jwt.MapClaims)["cognito:groups"].([]interface{})[0].(string)
	isAllowed := cmd.IsAllowed(endpointUrl, userGroup)

	// Return deny if user has no permission to access endpoint (based on user group)
	if !isAllowed {
		return cmd.AuthPolicyResponse("user", "Deny", request.MethodArn, nil), nil
	}

	// Check if claims are present and token is valid.
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		// return Allow authResponse with userEntity in authorizer context for next lambda in chain.
		return cmd.AuthPolicyResponse("user", "Allow", request.MethodArn, map[string]interface{}{"userEntity": claims["sub"].(string)}), nil
	}

	return cmd.AuthPolicyResponse("user", "Deny", request.MethodArn, nil), nil
}

func main() {
	lambda.Start(HandleRequest)
}
