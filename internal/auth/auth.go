package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extracts API key from headers of HTTP request
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth info found")
	}

	valArr := strings.Split(val, " ")
	if len(valArr) != 2 {
		return "", errors.New("malformed auth header")
	}

	if valArr[0] != "ApiKey" {
		return "", errors.New("malformed first path of auth header")
	}

	return valArr[1], nil
}
