package history

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Item struct {
	Original string `json:"original"`
	Short    string `json:"short"`
	Service  string `json:"service"`
}

var (
	data     []Item
	filePath string
)

// init runs once when the package is imported
func init() {
	var err error
	filePath, err = getHistoryFilePath()
	if err != nil {
		fmt.Println("⚠️ Could not determine history file path:", err)
		return
	}

	load()
}

// -------------------- Public API --------------------

func Add(original, short, service string) {
	data = append(data, Item{
		Original: original,
		Short:    short,
		Service:  service,
	})

	save()
}

func Print() {
	if len(data) == 0 {
		fmt.Println("📭 History is empty")
		return
	}

	fmt.Println("\n📜 Shortened URLs history:")
	for i, h := range data {
		fmt.Printf("%d) [%s]\n", i+1, h.Service)
		fmt.Println("   🔗", h.Original)
		fmt.Println("   ⭐", h.Short)
	}
}

func Clear() error {
	data = []Item{}
	return save()
}

// -------------------- Internal helpers --------------------

func load() {
	file, err := os.ReadFile(filePath)
	if err != nil {
		// File does not exist → first run
		return
	}

	_ = json.Unmarshal(file, &data)
}

func save() error {
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, jsonData, 0644)
}

func getHistoryFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(
		home,
		"Documents",
		"url-shortener",
		"history.json",
	), nil
}
