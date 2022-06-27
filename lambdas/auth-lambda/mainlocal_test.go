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
	request.AuthorizationToken = "eyJraWQiOiJVRUpkZ3hjU0VRUFhUS2twK3dOQnVlUkZpQzd1SXpnTTFHOW9cL3JBQk14RT0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiI3OGE4MjVhOC0zMThhLTRkNDUtOGNiOC1hNjUwOGVhZTNhNWUiLCJjb2duaXRvOmdyb3VwcyI6WyJ0ZWFjaGVyLWdyb3VwIl0sImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC5ldS1jZW50cmFsLTEuYW1hem9uYXdzLmNvbVwvZXUtY2VudHJhbC0xX2t3d0hIcHdNTCIsImNsaWVudF9pZCI6IjZlc2NsZWxyM3VoMXNpcDV2NzlkaGNxb291Iiwib3JpZ2luX2p0aSI6IjJlY2NkNGU4LWE5ODUtNDA0NS05MDJkLTgxZDA1MDg3NDkwNCIsImV2ZW50X2lkIjoiOGNmMDczNWUtOTFjMS00NDczLTgyMmEtNjJlNGZlOWU1YTA0IiwidG9rZW5fdXNlIjoiYWNjZXNzIiwic2NvcGUiOiJhd3MuY29nbml0by5zaWduaW4udXNlci5hZG1pbiIsImF1dGhfdGltZSI6MTY1NjM0NTgyMiwiZXhwIjoxNjU2MzQ5NDIyLCJpYXQiOjE2NTYzNDU4MjIsImp0aSI6IjVmODYxYzRjLWI4ZGMtNDYxYy1hZjE5LTk4NzA0OWI2YzZiNSIsInVzZXJuYW1lIjoidGVzdHVzZXIyIn0.sfYc60_-memts8lv5qp3dkkd6u6IX8oEhxRwyqi8wqXsP9EURqsOO9PB421R5atDknyxPrv-NEeM0LAUL-sq-xC8jBTrz8rSGGWByd8C42usRW9QKr5tedA-v1dWIu8XfaIghdEYDOP8uj9Xxnycz0P4zxFA7fBS9xvOrNqrxZ_FihnisG7t1xkJeTWIe5MxZG5zQefkp_vVz86H47-X0AmSSZB136zFHJw_c2xAauPnCDbkc4niBAwTtti06VkmovCYQpU60aTwI8Rtv11rDU_uaJ33ewikmFvckNdXVacebZJvBjdwKB8mz3lf-YUhaDlzFs2O4qOnmdi2y_RYpg"
	request.MethodArn = "arn:aws:execute-api:eu-central-1:614695401085:y1llo92zmi/test/GET/grades"

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
