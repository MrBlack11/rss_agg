package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracs an API Key from
// the headers of an HTTP request
// Example: Authorization: ApiKey <insert_key_here>
func GetAPIKey(headers http.Header) (string, error) {
	authString := headers.Get("Authorization")
	if authString == "" {
		return "", errors.New("no authentication info found")
	}

	splittedAuthString := strings.Split(authString, " ")
	if len(splittedAuthString) != 2 {
		return "", errors.New("malformed auth header")
	}

	if splittedAuthString[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return splittedAuthString[1], nil
}
