package cmd

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

// DecodeToken Function parse token without verification
// Request is previously authorized by custom lambda authorizer, so it's unnecessary at this step
func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, err
	}
	return nil, err
}
