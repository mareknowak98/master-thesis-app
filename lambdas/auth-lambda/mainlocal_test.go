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
	request.AuthorizationToken = "eyJraWQiOiJRZlo3TlNSdm1pWHFBYjllZVdZVGg4bDBTXC82UkRBRWZzQ2pkVmxBaGR4az0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJlZDY5ZDA1OS1hMzAyLTQ0ZjMtOTZhMC1iMmZmYmNhMWNkMGUiLCJjb2duaXRvOmdyb3VwcyI6WyJ0ZWFjaGVyLWdyb3VwIl0sImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC5ldS1jZW50cmFsLTEuYW1hem9uYXdzLmNvbVwvZXUtY2VudHJhbC0xX05UM1JLZ1pqTyIsImNsaWVudF9pZCI6IjcyYmp0YXVtbjhrcW0zczNvNTFsdDdzaXMxIiwib3JpZ2luX2p0aSI6ImNmNmQ2NTkzLWMzM2QtNGVlNC1iN2UwLTMwYjYxM2RjZDk1YiIsImV2ZW50X2lkIjoiYzNjNWU2YjMtNGQ3My00MWEwLWEwMWEtYzRiYTVkOWUxYWU3IiwidG9rZW5fdXNlIjoiYWNjZXNzIiwic2NvcGUiOiJhd3MuY29nbml0by5zaWduaW4udXNlci5hZG1pbiIsImF1dGhfdGltZSI6MTY1NzYzMDI2MywiZXhwIjoxNjU3NjMzODYzLCJpYXQiOjE2NTc2MzAyNjMsImp0aSI6ImUwOWQ2N2RiLTcwMjgtNDc3My05ZDYyLTliZTljNjMzYzQ1YyIsInVzZXJuYW1lIjoidGVzdHVzZXIzIn0.kaLhnwibTTy2GTav0zr6hqikZ9cPJSHLUNJo_JLSPtZHyW3yKH5uZGT3k3wbN39Ca4Ic7jsP85chJr6jBA-lkn_skFUdG854MzgUy_ZI3z42VC5kELWK1b8WawPcC4QDu98I0OHf4jNxIlcUnBqdYUrerxmQhBHuW3kmm-Q1WevNxmFH4-FkB0LEQQH8pJBUVkjkTepkok88oEUAdEVSwaVs7hrx4QttNVny8s8lTuROkseWYcqx8x2SWkWE405Bk_AiPgLC5W0pQF-k7MFNRzwzb5LPobzRXwd_Gkha14U3xLZyjTBSKOnb7cFX7wniB-plguq4CTRPbBKO7vfDvA"
	request.MethodArn = "arn:aws:execute-api:eu-central-1:614695401085:y1llo92zmi/test/POST/grades"

	err := os.Setenv("REGION", "eu-central-1")
	if err != nil {
		return events.APIGatewayV2CustomAuthorizerV1Request{}
	}
	err = os.Setenv("USER_POOL_ID", "eu-central-1_NT3RKgZjO")
	if err != nil {
		return events.APIGatewayV2CustomAuthorizerV1Request{}
	}

	return request
}
