package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	tinyURLAPI = "http://tinyurl.com/api-create.php?url="
	isgdAPI    = "https://is.gd/create.php?format=simple&url="
)

func ShortenByName(original, service string) (string, string, error) {
	escaped := url.QueryEscape(original)

	var apiURL, name string

	switch service {
	case "isgd":
		apiURL = isgdAPI + escaped
		name = "is.gd"
	default:
		apiURL = tinyURLAPI + escaped
		name = "TinyURL"
	}

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", "", fmt.Errorf("Network error: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	result := strings.TrimSpace(string(body))

	if strings.HasPrefix(result, "Error") {
		return "", "", fmt.Errorf("API error: %s", result)
	}

	return result, name, nil
}
