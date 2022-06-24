package cmd

import (
	"regexp"
)

// function return bool value if given group is authorized for exact request
// data is stored in hash map for better performance (comparing e.g. to read from json file)
//func isAllowed(userGroup string, method string) bool {
//	permissions := map[string][]string{
//		"student-group": []string{},
//		"teacher-group": []string{},
//	}
//	return false
//}

func MethodArnToUrl(methodArn string) string {
	// TODO: to finish
	var validID = regexp.MustCompile(`/^(?:[^\/]*\/){2}\s*/`)

	return validID.FindString(methodArn)
}
