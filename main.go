package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]
	// default values
	maxConcurrency := 5
	maxPages := 50

	// optional arg: maxConcurrency
	if len(os.Args) >= 3 {
		val, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid maxConcurrency value")
			os.Exit(1)
		}
		maxConcurrency = val
	}

	// optional arg: maxPages
	if len(os.Args) >= 4 {
		val, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("invalid maxPages value")
			os.Exit(1)
		}
		maxPages = val
	}
	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizedURL := range cfg.pages {
		fmt.Printf("found: %s\n", normalizedURL)
	}
	writeJSONReport(cfg.pages, "report.json")
}

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "BootCrawler/1.0")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("Client Error code: %d", res.StatusCode)
	}
	contentType := res.Header.Get("Content-Type")
	if contentType == "" {
		return "", fmt.Errorf("missing Content-Type header")
	}
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("invalid Content-Type: %s", contentType)
	}
	htmlBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	htmlString := string(htmlBytes)
	return htmlString, nil
}
