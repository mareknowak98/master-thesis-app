package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-jwt/jwt/v4"
	"mylearnproject/lambdas/auth-lambda/cmd"
)

// HandleRequest Handle lambda request (from api gateway)
func HandleRequest(ctx context.Context, request events.APIGatewayV2CustomAuthorizerV1Request) (events.APIGatewayCustomAuthorizerResponse, error) {
	endpointUrl := cmd.MethodArnToUrl(request.MethodArn)
	parsedToken, err := cmd.DecodeToken(request.AuthorizationToken)

	// Return deny if token cannot be parsed
	if err != nil {
		return authPolicyResponse("user", "Deny", request.MethodArn, nil), nil
	}

	//get user group from token claims, then map []interface {} to string
	userGroup := parsedToken.Claims.(jwt.MapClaims)["cognito:groups"].([]interface{})[0].(string)
	isAllowed := cmd.IsAllowed(endpointUrl, userGroup)

	// Return deny if user has no permission to access endpoint (based on user group)
	if !isAllowed {
		return authPolicyResponse("user", "Deny", request.MethodArn, nil), nil
	}

	// Check if claims are present and token is valid.
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		// return Allow authResponse with userEntity in authorizer context for next lambda in chain.
		return authPolicyResponse("user", "Allow", request.MethodArn, map[string]interface{}{"userEntity": claims["sub"].(string)}), nil
	}

	return authPolicyResponse("user", "Deny", request.MethodArn, nil), nil
}

// authPolicyResponse for authorization/ de-authorization of given principal with supplied context
// Used as built-in map for better performance
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
