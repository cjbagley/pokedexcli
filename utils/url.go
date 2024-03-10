package utils

import (
	"net/url"
	"strings"
)

func GetEndpointFromUrl(raw_url string) (string, error) {
	if raw_url == "" {
		return "", nil
	}

	u, err := url.Parse(raw_url)
	if err != nil {
		return "", err
	}

	parts := strings.Split(u.Path, "/")
	path := parts[len(parts)-1]
	query := u.Query().Encode()
	if query != "" {
		path += "?" + query
	}

	return path, nil
}
