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
	request.AuthorizationToken = "eyJraWQiOiJVRUpkZ3hjU0VRUFhUS2twK3dOQnVlUkZpQzd1SXpnTTFHOW9cL3JBQk14RT0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiI3OGE4MjVhOC0zMThhLTRkNDUtOGNiOC1hNjUwOGVhZTNhNWUiLCJjb2duaXRvOmdyb3VwcyI6WyJ0ZWFjaGVyLWdyb3VwIl0sImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC5ldS1jZW50cmFsLTEuYW1hem9uYXdzLmNvbVwvZXUtY2VudHJhbC0xX2t3d0hIcHdNTCIsImNsaWVudF9pZCI6IjZlc2NsZWxyM3VoMXNpcDV2NzlkaGNxb291Iiwib3JpZ2luX2p0aSI6ImI4MWM2YjlmLTdlYjUtNGUyNi04Y2JjLTMyMjUwYjk4MGYxNSIsImV2ZW50X2lkIjoiZGY4MTJmM2EtYmFjZS00Mzk1LTliYTYtMDNkZDJmNzgwNTdjIiwidG9rZW5fdXNlIjoiYWNjZXNzIiwic2NvcGUiOiJhd3MuY29nbml0by5zaWduaW4udXNlci5hZG1pbiIsImF1dGhfdGltZSI6MTY1NjMyODIxNywiZXhwIjoxNjU2MzMxODE3LCJpYXQiOjE2NTYzMjgyMTcsImp0aSI6IjE1NWI1ZTU2LTFiMWQtNDU2Yi04NmE1LTVjZDc1MWI3MTBkYyIsInVzZXJuYW1lIjoidGVzdHVzZXIyIn0.qN3Ubs3CpwtpSZxeWVpf_7IdnOeMM1uytTHi9vC1_2Znr2oO2kw4QBuKKurRCD6xiMeRQGhGFNcmMbHPTCbfYWKH0oA1oPshc7AN2EbLuihCR6EoAqFtIvESkg8CQVVWhRoxGAoSw8VlCPfpQnKi9Y7_qGh7qUeXTORVefgDc7YDguYrMFNkOPlnF2raxOfXIhFoIWCCzrH_70ielwBM0swyUli5NIhNkflaYkO3zzTTYFa-Fv_753-YG47lElOtJOuUc6gQNTDPzmq7mMngioKUSohHCfKis7rNRM5Am5aIOeoKE9fydMHlWfs8KXjWt8_sAjaHY7MLFXMNJcH_6A"
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
