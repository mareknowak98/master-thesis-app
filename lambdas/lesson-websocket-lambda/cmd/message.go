package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"

	//"github.com/aws/aws-sdk-go/service/dynamodb"

	"os"
)

// Default method - send new slides to all connected users
func (c *Client) Default(request APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	tableNameConnections := os.Getenv("CONNECTIONS_TABLE")
	tableNameLessons := os.Getenv("LESSONS_TABLE")
	websocketApi := os.Getenv("WEBSOCKET_API")
	deploymentType := os.Getenv("DEPLOYMENT_TYPE")

	fmt.Println(tableNameConnections, tableNameLessons, websocketApi)
	fmt.Println(request.Body)

	// slide to send
	var slideToSwap SlideSwap
	err := json.Unmarshal([]byte(request.Body), &slideToSwap)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	fmt.Println(slideToSwap)

	hashKey := map[string]string{"keyName": "LessonId", "value": slideToSwap.LessonId}
	sortKey := map[string]string{"keyName": "SlideId", "value": slideToSwap.SlideId}

	queryResult, err := c.queryDynamo(hashKey, sortKey, tableNameLessons)
	var lessonSlide LessonSlide
	err = attributevalue.UnmarshalMap(queryResult.Items[0], &lessonSlide)

	fmt.Println("======9")
	fmt.Println(err)
	fmt.Println(lessonSlide)

	jsonOut, err := json.Marshal(lessonSlide)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	fmt.Println("jsonout")
	fmt.Println(string(jsonOut))

	for _, e := range queryResult.Items {
		fmt.Printf("%#v\n", e)
	}
	//
	fmt.Println("======1")
	// query connections table
	hashKey = map[string]string{"keyName": "LessonId", "value": slideToSwap.LessonId}
	sortKey = map[string]string{}
	connections, err := c.queryDynamo(hashKey, sortKey, tableNameConnections)

	for _, e := range connections.Items {
		fmt.Printf("%#v\n", e)
	}
	//

	fmt.Println("======2")

	var lessonSockets []LessonSocket
	err = attributevalue.UnmarshalListOfMaps(connections.Items, &lessonSockets)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	for _, lessonSocket := range lessonSockets {
		fmt.Println(lessonSocket)

		//apigateway, err := c.WebsocketApiCL.PostToConnection(context.Background(), &apigatewaymanagementapi.PostToConnectionInput{
		//	ConnectionId: aws.String(request.RequestContext.ConnectionID),
		//	Data:         jsonOut,
		//})

		// TODO there was issue with IAM
		sess := GetSession()

		input := &apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: aws.String(request.RequestContext.ConnectionID),
			Data:         jsonOut,
		}

		apigateway := apigatewaymanagementapi.New(sess, aws.NewConfig().WithEndpoint(fmt.Sprintf("%s.execute-api.eu-central-1.amazonaws.com/%s", websocketApi, deploymentType)))
		fmt.Printf("\ninput config %#v\n", apigateway.Endpoint)

		_, err = apigateway.PostToConnection(input)
		if err != nil {
			fmt.Printf("ERROR %#v\n", err)
		}

		if err != nil {
			fmt.Println(err)
			return events.APIGatewayProxyResponse{}, err
		}
		//fmt.Println(apigateway)

	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
