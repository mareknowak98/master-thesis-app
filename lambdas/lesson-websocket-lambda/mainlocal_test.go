package main

import (
	"encoding/json"
	"fmt"
	"mylearnproject/lambdas/lesson-websocket-lambda/cmd"
	"os"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	_ = os.Setenv("MESSAGES_TABLE", "chat-connections")
	_ = os.Setenv("WEBSOKET_API", "rjyithf86i")

	//tmp := cmd.APIGatewayWebsocketProxyRequest{}

	input := getDebugInput()

	res, _ := Handler(input)
	fmt.Println(res)
}

//func getDebugInput() cmd.APIGatewayWebsocketProxyRequest {
//	_ = os.Setenv("REGION", "eu-central-1")
//	_ = os.Setenv("CONNECTIONS_TABLE", "mylearn-lessons-connections")
//	_ = os.Setenv("LESSONS_TABLE", "mylearn-lessons")
//	_ = os.Setenv("WEBSOCKET_API", "pgcpqpf7w1")
//	_ = os.Setenv("USER_POOL_ID", "eu-central-1_LsCgMRDvD")
//
//	var request cmd.APIGatewayWebsocketProxyRequest
//
//	m := make(map[string]string)
//	m["lessonId"] = "123"
//	request.RequestContext.ConnectionID = "sdasdasdas"
//	request.QueryStringParameters = m
//
//	request.RequestContext.RouteKey = "$connect"
//
//	return request
//}

//func getDebugInput() cmd.APIGatewayWebsocketProxyRequest {
//	_ = os.Setenv("REGION", "eu-central-1")
//	_ = os.Setenv("CONNECTIONS_TABLE", "mylearn-lessons-connections")
//	_ = os.Setenv("LESSONS_TABLE", "mylearn-lessons")
//	_ = os.Setenv("WEBSOCKET_API", "pgcpqpf7w1")
//	_ = os.Setenv("USER_POOL_ID", "eu-central-1_LsCgMRDvD")
//
//	var request cmd.APIGatewayWebsocketProxyRequest
//
//	//m := make(map[string]string)
//	//m["lessonId"] = "123"
//	request.RequestContext.ConnectionID = "sdasdasdas"
//	//request.QueryStringParameters = m
//
//	request.RequestContext.RouteKey = "$disconnect"
//
//	return request
//}

func getDebugInput() cmd.APIGatewayWebsocketProxyRequest {
	_ = os.Setenv("REGION", "eu-central-1")
	_ = os.Setenv("CONNECTIONS_TABLE", "mylearn-lessons-connections")
	_ = os.Setenv("LESSONS_TABLE", "mylearn-lessons")
	_ = os.Setenv("WEBSOCKET_API", "pgcpqpf7w1")
	_ = os.Setenv("USER_POOL_ID", "eu-central-1_LsCgMRDvD")
	_ = os.Setenv("DEPLOYMENT_TYPE", "test")

	var request cmd.APIGatewayWebsocketProxyRequest

	m := make(map[string]string)
	request.RequestContext.ConnectionID = "sdasdasdas"
	request.QueryStringParameters = m

	var inp cmd.SlideSwap
	err := json.Unmarshal([]byte(`
		{
		   "lessonId" : "12",
		   "slideId" : "2"
		}`), &inp)

	if err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(inp)
	if err != nil {
		fmt.Println(err)
	}
	request.Body = string(b)

	request.RequestContext.RouteKey = "$message"

	return request
}
