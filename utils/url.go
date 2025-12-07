package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func NormalizeURL(u string) string {
	if strings.HasPrefix(u, "http://") || strings.HasPrefix(u, "https://") {
		return u
	}
	return "https://" + u
}

func CheckAvailability(u string) bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("HEAD", u, nil)
	if err != nil {
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode < 400
}

func ShowResult(original, short string) {
	fmt.Println("\n✅ Successfully shortened!")
	fmt.Println("🔗 Original:", original)
	fmt.Println("⭐ Short:", short)

	saved := len(original) - len(short)
	fmt.Printf("📉 Characters saved: %d (%d → %d)\n",
		saved, len(original), len(short))
}
