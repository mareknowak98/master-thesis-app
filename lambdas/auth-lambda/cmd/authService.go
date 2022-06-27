package cmd

import (
	"strings"
)

// IsAllowed function return bool value if given group is authorized for exact request
//data is stored in hash map for better performance (comparing e.g. to read from json file)
func IsAllowed(method string, userGroup string) bool {
	permissions := map[string]map[string]bool{
		"GET/grades":  map[string]bool{"teacher-group": true, "student-group": true},
		"POST/grades": map[string]bool{"teacher-group": true},
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
