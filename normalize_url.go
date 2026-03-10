package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	uHost := u.Host
	uPath := u.Path

	fullPath := uHost + uPath
	fullPath = strings.ToLower(fullPath)

	normalURL := strings.TrimSuffix(fullPath, "/")

	return normalURL, nil
}
