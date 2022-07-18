package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"os"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	input := getDebugInput()

	res, _ := HandleRequest(input)

	fmt.Println("Response:")
	fmt.Println(res.StatusCode)
	fmt.Println(res.Body)
}

func getDebugInput() events.APIGatewayProxyRequest {
	var request events.APIGatewayProxyRequest
	_ = os.Setenv("USER_POOL_ID", "eu-central-1_NT3RKgZjO")
	_ = os.Setenv("REGION", "eu-central-1")
	_ = os.Setenv("MESSAGES_TABLE", "mylearn-chat-messages")

	request.HTTPMethod = "GET"
	request.Path = "/messages"
	m := make(map[string]string)
	m["UserTo"] = "testuser2"
	request.QueryStringParameters = m

	m2 := make(map[string]string)
	m2["Authorization"] = "eyJraWQiOiJRZlo3TlNSdm1pWHFBYjllZVdZVGg4bDBTXC82UkRBRWZzQ2pkVmxBaGR4az0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJlZDY5ZDA1OS1hMzAyLTQ0ZjMtOTZhMC1iMmZmYmNhMWNkMGUiLCJjb2duaXRvOmdyb3VwcyI6WyJ0ZWFjaGVyLWdyb3VwIl0sImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC5ldS1jZW50cmFsLTEuYW1hem9uYXdzLmNvbVwvZXUtY2VudHJhbC0xX05UM1JLZ1pqTyIsImNsaWVudF9pZCI6IjcyYmp0YXVtbjhrcW0zczNvNTFsdDdzaXMxIiwib3JpZ2luX2p0aSI6ImRhZTE0YzZlLWMyNjItNGFjYi04MmMzLTY1MDQ1MGE0NjZmNCIsImV2ZW50X2lkIjoiOTZmNmRiZTEtM2Q5ZS00ZmZkLWE5ZDgtYTFmMDEyNzBhNzllIiwidG9rZW5fdXNlIjoiYWNjZXNzIiwic2NvcGUiOiJhd3MuY29nbml0by5zaWduaW4udXNlci5hZG1pbiIsImF1dGhfdGltZSI6MTY1ODE2NDc5MSwiZXhwIjoxNjU4MTY4MzkxLCJpYXQiOjE2NTgxNjQ3OTEsImp0aSI6IjRmMzg2ODE3LTIwZGQtNDNkYS04ZWU5LTQ0YjI1YWE3ZTc5YyIsInVzZXJuYW1lIjoidGVzdHVzZXIzIn0.DGK-fBqR1HYBkmSBUJBppqyV9uVE9ybv6oeUv1K4WeTeV29CIpC13tFn-4hQ-1MhtycSy5jFMHlsQXB3JMMpnUkKDHzd0Ffkdz0SK7oczBc9of6Lc2dCSOImg2dflZTJ0rNdsIlk0uotcNVKgtHNq9Df9MyPscnu0WHcrrjEvLUe_Q-rqvODnyNtrSSvx00WFjgYqnti9B-ooR30YTwEj7tRlB4_yfFkWQwobv8PF9HyKJlsfGy0vpPGyK23BsjFXV2rSczt4mdPDh94B3pVWRAvHIl81RQan-wQB24_sui6mSF_cDMPnlIcFAfhKR7XjeSC2MZOiR3ss17wjA2skw"
	request.Headers = m2
	return request
}
