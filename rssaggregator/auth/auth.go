package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("authorization header not found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed Authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("authorization header is invalid")
	}

	return vals[1], nil
}
