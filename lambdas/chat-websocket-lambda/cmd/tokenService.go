package cmd

import (
	"fmt"
	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"time"
)

// DecodeToken Function decode given JWT token
func DecodeToken(tokenString string) (*jwt.Token, error) {
	jwks, err := fetchJWKS()

	token, err := jwt.Parse(tokenString, jwks.Keyfunc)
	if err != nil {
		if verr, ok := err.(*jwt.ValidationError); ok {
			if verr.Errors == jwt.ValidationErrorMalformed {
				return nil, fmt.Errorf("unauthorized")
			}
			if verr.Errors == jwt.ValidationErrorExpired {
				return nil, fmt.Errorf("token is expired")
			}
		}
	}
	return token, err
}

// fetchJWKS Function used to connect user with exact cognito distribution
func fetchJWKS() (*keyfunc.JWKS, error) {
	// Read environmental variables
	region := os.Getenv("REGION")
	userPoolId := os.Getenv("USER_POOL_ID")

	refreshInterval := time.Hour
	refreshRateLimit := time.Minute * 5
	refreshTimeout := time.Second * 10
	refreshUnknownKID := true
	options := keyfunc.Options{
		RefreshErrorHandler: func(err error) {
			log.Printf("There was an error with the jwt.KeyFunc\nError:%s\n", err.Error())
		},
		RefreshInterval:   refreshInterval,
		RefreshRateLimit:  refreshRateLimit,
		RefreshTimeout:    refreshTimeout,
		RefreshUnknownKID: refreshUnknownKID,
	}
	return keyfunc.Get(fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPoolId), options)
}
