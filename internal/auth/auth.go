package auth

import (
	"errors"
	"net/http"
	"strings"
)

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
