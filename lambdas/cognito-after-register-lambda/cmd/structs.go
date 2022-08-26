package cmd

import "github.com/aws/aws-lambda-go/events"

type InputUsers struct {
	UserId    string `json:"UserId"`
	Username  string `json:"Username"`
	Email     string `json:"Email"`
	CreatedAt string `json:"CreatedAt"`
	UserGroup string `json:"UserGroup"`
	UserClass string `json:"UserClass"`
}

type CognitoEventRequest struct {
	DatasetName    string                                 `json:"datasetName"`
	DatasetRecords map[string]events.CognitoDatasetRecord `json:"datasetRecords"`
	TriggerSource  string                                 `json:"triggerSource"`
	IdentityID     string                                 `json:"identityId"`
	UserPoolID     string                                 `json:"userPoolId"`
	Region         string                                 `json:"region"`
	Version        string                                 `json:"version"`
	Request        map[string]map[string]string           `json:"request"`
	Username       string                                 `json:"userName"`
}

type CognitoEventResponse struct {
	TriggerSource string   `json:"triggerSource"`
	Username      string   `json:"userName"`
	UserPoolID    string   `json:"userPoolId"`
	Region        string   `json:"region"`
	Version       int      `json:"version"`
	Response      struct{} `json:"response"`
}
