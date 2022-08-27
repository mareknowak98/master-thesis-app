package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func (c *Client) getFilesFromFolder(s3Name, folderName string) (s3.ListObjectsOutput, error) {
	files, err := c.S3Cl.ListObjects(context.Background(), &s3.ListObjectsInput{
		Bucket: aws.String(s3Name),
		Prefix: aws.String(folderName),
	})

	if err != nil {
		return s3.ListObjectsOutput{}, err
	}

	return *files, nil
}
