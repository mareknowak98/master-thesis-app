package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"os"
	"testing"
)

// File to test running lambda locally

func TestHandleRequest(t *testing.T) {

	input := getDebugInput()

	res, _ := HandleRequest(context.Background(), input)

	fmt.Println("Response:")
	fmt.Println(res)
}

func getDebugInput() events.APIGatewayV2CustomAuthorizerV1Request {
	var request events.APIGatewayV2CustomAuthorizerV1Request
	request.AuthorizationToken = "eyJraWQiOiIyYXhZRWQzRzBLNXZSNVRoRWppalREV3BWSG1PalF1TlFiamJvZTdrMGs0PSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiI3OGE4MjVhOC0zMThhLTRkNDUtOGNiOC1hNjUwOGVhZTNhNWUiLCJjb2duaXRvOmdyb3VwcyI6WyJ0ZWFjaGVyLWdyb3VwIl0sImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLmV1LWNlbnRyYWwtMS5hbWF6b25hd3MuY29tXC9ldS1jZW50cmFsLTFfa3d3SEhwd01MIiwiY29nbml0bzp1c2VybmFtZSI6InRlc3R1c2VyMiIsIm9yaWdpbl9qdGkiOiI3MTgwOGE2OS0wOGRkLTQ5YjAtYjQzNS05MmU0ZTc0OGIyMjIiLCJhdWQiOiI2ZXNjbGVscjN1aDFzaXA1djc5ZGhjcW9vdSIsImV2ZW50X2lkIjoiOWZhN2I2MjItZTVlZi00OWU2LTk0ZjktZjZjZjU4YTI0M2UwIiwidG9rZW5fdXNlIjoiaWQiLCJhdXRoX3RpbWUiOjE2NTYxMDY2OTcsImV4cCI6MTY1NjExMDI5NywiaWF0IjoxNjU2MTA2Njk3LCJqdGkiOiJmYWU1OTUxNy1mYTFiLTQ3M2YtYWIwOS01OWIzMWRlMTljM2YiLCJlbWFpbCI6ImphbmVAZXhhbXBsZS5jb20ifQ.a1itqTwAt3IuqGuxidprWG6wqp6VwWH6nP1yYz9He4ltJwklIpGbdhXUrjJDufzhgT0OnFCBFYe-v59hwNMYW4XfEy1eKmJIAD4oejaoMx8KiUqlVr1mhhG2Ke4CqcOjliOPpA0pZtlUiXjcUtFZzHabghYUHYaJs1vb-LfADJLUjgaOept6CF3Q256qKCJnrv4jTiN2lq_nIx8jC_94Q7_D5jImrrnFfzptSTOs0ttQq0MvnprEb1evlQH00Ch3V12li8_xTvvqCvWnxvKaMlt6hRg86jgFzIwKdSwGrkAOj-YnggXuwPoWz6PgIAWGMfLx5eoz4uFxi1A2fCBZdQ"
	err := os.Setenv("REGION", "eu-central-1")
	if err != nil {
		return events.APIGatewayV2CustomAuthorizerV1Request{}
	}
	err = os.Setenv("USER_POOL_ID", "eu-central-1_kwwHHpwML")
	if err != nil {
		return events.APIGatewayV2CustomAuthorizerV1Request{}
	}

	return request
}
