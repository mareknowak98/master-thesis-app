package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"time"
)

// SaveGrade Function save grade to DynamoDB with timestamp
func (c *Client) SaveGrade(grade InputGrades, tableName string) (error) {
	// Get current date and attach it to struct
	dt := time.Now()
	grade.Date = dt.Format("01-02-2006 15:04:05")

	// Marshall each element to 'map[string]types.AttributeValue' format
	av, err := attributevalue.MarshalMap(grade)

	if err != nil {
		fmt.Println(err)
		return err
	}

	// Create dynamoDB item input
	inp := dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: av,
	}

	// Put item in dynamoDB
	_, err = c.DynamoCl.PutItem(context.Background(), &inp)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
