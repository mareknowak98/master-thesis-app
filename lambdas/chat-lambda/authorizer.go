package main

import (
	"fmt"
	"github.com/MicahParks/keyfunc"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"os"
	"time"
)

// fetchJWKS Function used to connect user with exact cognito distribution
func fetchJWKS() (*keyfunc.JWKS, error) {
	// Read environmental variables
	region := os.Getenv("REGION")
	userPoolId := os.Getenv("USER_POOL_ID")

	refreshInterval := time.Hour
	refreshRateLimit := time.Minute * 5
	refreshTimeout := time.Second * 10
	refreshUnknownKID := true
	options := keyfunc.Options{
		RefreshErrorHandler: func(err error) {
			log.Printf("There was an error with the jwt.KeyFunc\nError:%s\n", err.Error())
		},
		RefreshInterval:   refreshInterval,
		RefreshRateLimit:  refreshRateLimit,
		RefreshTimeout:    refreshTimeout,
		RefreshUnknownKID: refreshUnknownKID,
	}
	return keyfunc.Get(fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPoolId), options)
}

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

//
//// Authorizer custom api authorizer
//func Authorizer(request APIGatewayWebsocketProxyRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
//	tokenString := request.QueryStringParameters["token"]
//
//	// Fetch all keys
//	jwkSet, err := fetchJWKS()
//	//jwkSet, err := jwk.Fetch("https://cognito-idp.us-east-1.amazonaws.com/us-east-1_vvx4f42sK/.well-known/jwks.json")
//	if err != nil {
//		log.Fatalln("Unable to fetch keys")
//	}
//
//	token, err := jwt.Parse(tokenString, jwkSet.Keyfunc)
//	if err != nil {
//		if verr, ok := err.(*jwt.ValidationError); ok {
//			if verr.Errors == jwt.ValidationErrorMalformed {
//				return events.APIGatewayCustomAuthorizerResponse{}, fmt.Errorf("unauthorized")
//			}
//			if verr.Errors == jwt.ValidationErrorExpired {
//				return events.APIGatewayCustomAuthorizerResponse{}, fmt.Errorf("token is expired")
//			}
//		}
//	}
//	//// Verify
//	//t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
//	//	keys := jwkSet.LookupKeyID(t.Header["kid"].(string))
//	//	return keys[0].Materialize()
//	//})
//	if err != nil || !token.Valid {
//		log.Fatalln("Unauthorized")
//	}
//
//	claims := token.Claims.(jwt.MapClaims)
//
//	fmt.Printf("Token :%#v\n", token)
//
//	return events.APIGatewayCustomAuthorizerResponse{
//		PrincipalID: "me",
//		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
//			Version: "2012-10-17",
//			Statement: []events.IAMPolicyStatement{
//				events.IAMPolicyStatement{
//					Action:   []string{"execute-api:*"},
//					Effect:   "Allow",
//					Resource: []string{request.MethodArn},
//				},
//			},
//		},
//		Context: claims,
//	}, nil
//}
