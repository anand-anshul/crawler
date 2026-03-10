package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Printf("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl of: %s", args[0])
	}
	pages := make(map[string]int)
	crawlPage(args[0], args[0], pages)
	for normalizedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
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

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}
	if baseURL.Hostname() != currentURL.Hostname() {
		return
	}
	normCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}
	_, exists := pages[normCurrentURL]
	if exists {
		pages[normCurrentURL] += 1
		return
	} else {
		pages[normCurrentURL] = 1
	}
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		return
	}
	fmt.Printf(html)
	URLs, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		return
	}
	for _, url := range URLs {
		crawlPage(rawBaseURL, url, pages)
	}
}
