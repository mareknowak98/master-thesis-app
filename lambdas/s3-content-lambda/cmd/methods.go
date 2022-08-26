package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
)

// ImageRequestBody is the image request body
type ImageRequestBody struct {
	FileName string `json:"filename"`
	Body     string `json:"body"`
}

func (c *Client) SaveMaterials(request events.APIGatewayProxyRequest, s3MaterialsBucketName string) (string, error) {
	filename := request.Headers["fileName"]
	fmt.Println(filename)

	jsonBod, err := json.Marshal(request.Body)
	if err != nil {
		log.Println(err.Error())
		return "can't create json response", errors.New("Error: can't create json response")
	}

	fmt.Println("after")
	//fmt.Println(string())
	//
	//_, err = c.S3Cl.PutObject(context.Background(), &s3.PutObjectInput{
	//	Bucket: aws.String(s3MaterialsBucketName),
	//	Key:    aws.String(filename),
	//	Body:   bytes.NewReader(jsonBod),
	//})
	//fmt.Println("after")

	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("hexaco"),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(jsonBod),
	})
	fmt.Println("after2")

	if err != nil {
		log.Println(err.Error())
		return "can't put item", errors.New("Error: can't create json response")
	}

	return "", nil
}

func (c *Client) GetMaterials(request events.APIGatewayProxyRequest, s3MaterialsBucketName string) (string, error) {

	return "", nil
}
