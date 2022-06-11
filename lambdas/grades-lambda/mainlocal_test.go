package main

import (
	"context"
	"encoding/json"
	"fmt"
	"mylearnproject/lambdas/grades-lambda/cmd"
	"testing"
)

// File to test running lambda locally

func TestHandleRequest(t *testing.T) {

	input := getDebugInput()

	res, _ := HandleRequest(context.Background(), input)

	fmt.Println("Response:")
	fmt.Println(res.StatusCode)
	fmt.Println(res.Body)

}

func getDebugInput() cmd.InputGrades {
	var inp cmd.InputGrades

	err := json.Unmarshal([]byte(`
	{
		"UserId": "123",
		"ClassYear": "sometest",
		"Grade": "4+"
	}`), &inp)

	if err != nil {
		fmt.Println(err)
	}

	return inp
}
