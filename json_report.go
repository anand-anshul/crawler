package main

import (
	"encoding/json"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	keys := []string{}
	pageDati := []PageData{}
	for key, _ := range pages {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		pageDati = append(pageDati, pages[key])
	}
	jsonData, err := json.MarshalIndent(pageDati, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
