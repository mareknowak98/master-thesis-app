package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"strings"
)

// GetFiles return list of files from given s3 bucket
// params:
// folder (string) - folder name (optional)
func (c *Client) GetFiles(request events.APIGatewayProxyRequest, s3Name string) (string, error) {
	// If query string empty
	if len(request.QueryStringParameters) == 0 {
		return "", fmt.Errorf("not enough parameters\n")
	}

	folder := request.QueryStringParameters["folder"]
	files, err := c.getFilesFromFolder(s3Name, folder)
	if err != nil {
		return "", err
	}

	var resp []FileData

	for _, item := range files.Contents {
		var tmp FileData
		tmp.FileName = *item.Key
		resp = append(resp, tmp)
	}

	jsonOut, err := json.Marshal(resp)

	if err != nil {
		return "", err
	}

	return string(jsonOut), nil
}

// GetFolders return list of files from given s3 bucket
// params:
// subfolder (string) - subfolder name to show nested folders (optional)
func (c *Client) GetFolders(request events.APIGatewayProxyRequest, s3Name string) (string, error) {
	// If query string empty
	if len(request.QueryStringParameters) == 0 {
		return "", fmt.Errorf("not enough parameters\n")
	}
	subfolder := request.QueryStringParameters["subfolder"]
	files, err := c.getFilesFromFolder(s3Name, subfolder)

	if err != nil {
		return "", err
	}

	var resp []FileData
	for _, item := range files.Contents {
		//find last index of '/'
		lastInd := strings.LastIndex(*item.Key, "/")
		if lastInd >= 0 {
			var tmp FileData
			tmp.FileName = (*item.Key)[:lastInd]
			resp = append(resp, tmp)
		}
	}
	jsonOut, err := json.Marshal(resp)

	if err != nil {
		return "", err
	}

	return string(jsonOut), nil
}
