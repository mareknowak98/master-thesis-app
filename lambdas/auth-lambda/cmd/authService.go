package cmd

import (
	"github.com/aws/aws-lambda-go/events"
	"strings"
)

// IsAllowed function return bool value if given group is authorized for exact request
// data is stored in map of maps for better performance (comparing e.g. to read from json file) o
func IsAllowed(method string, userGroup string) bool {
	permissions := map[string]map[string]bool{
		"GET/grades":        map[string]bool{"teacher-group": true, "student-group": true, "parent-group": true, "admin-group": true},
		"POST/grades":       map[string]bool{"teacher-group": true},
		"GET/users":         map[string]bool{"teacher-group": true, "student-group": true, "parent-group": true, "admin-group": true},
		"GET/messages":      map[string]bool{"teacher-group": true, "student-group": true, "parent-group": true, "admin-group": true},
		"GET/classes":       map[string]bool{"teacher-group": true},
		"POST/classes":      map[string]bool{"teacher-group": true},
		"DELETE/classes":    map[string]bool{"teacher-group": true},
		"GET/classUsers":    map[string]bool{"teacher-group": true, "student-group": true, "parent-group": true, "admin-group": true},
		"POST/classUsers":   map[string]bool{"teacher-group": true},
		"DELETE/classUsers": map[string]bool{"teacher-group": true},
		"GET/me":            map[string]bool{"teacher-group": true, "student-group": true, "parent-group": true, "admin-group": true},
		"GET/slides":        map[string]bool{"teacher-group": true, "student-group": true, "parent-group": true, "admin-group": true},
		"POST/slides":       map[string]bool{"teacher-group": true},
		"DELETE/slides":     map[string]bool{"teacher-group": true},
		"GET/lessons":       map[string]bool{"teacher-group": true, "student-group": true, "parent-group": true, "admin-group": true},
		"GET/files":         map[string]bool{"teacher-group": true, "student-group": true, "parent-group": true, "admin-group": true},
		"GET/folders":       map[string]bool{"teacher-group": true, "student-group": true, "parent-group": true, "admin-group": true},
		"GET/s3":            map[string]bool{"teacher-group": true, "student-group": true, "parent-group": true, "admin-group": true},
		"PUT/s3":            map[string]bool{"teacher-group": true, "student-group": true, "parent-group": true, "admin-group": true},
		"POST/manageGroups": map[string]bool{"admin-group": true},
	}

	if endpointPerm, ok := permissions[method]; ok {
		if isAllowed, ok := endpointPerm[userGroup]; ok {
			return isAllowed
		}
	}

	return false
}

// MethodArnToUrl Function get url from Method arn
// arn:aws:execute-api:eu-central-1:614695401085:y1llo92zmi/test/GET/grades -> GET/grades
func MethodArnToUrl(methodArn string) string {
	var validID = strings.Split(methodArn, "/")[2:]
	return strings.Join(validID, "/")
}

// authPolicyResponse for authorization/ de-authorization of given principal with supplied context
// Used as built-in map for better performance
func AuthPolicyResponse(principalID, effect, resource string, context map[string]interface{}) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalID}
	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}
	authResponse.Context = context
	return authResponse
}
