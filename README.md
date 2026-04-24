# BootCrawler

## Overview

BootCrawler is a concurrent web crawler written in Go. It starts from a base URL, crawls pages within the same domain, and extracts structured data such as headings, first paragraphs, links, and images. Results are saved as a JSON report.

---

## Installation

```bash
git clone https://your.repo.url/bootcrawler.git
cd bootcrawler
go mod tidy
go build -o bootcrawler
```

---

## Usage

```bash
./bootcrawler <baseURL> [maxConcurrency] [maxPages]
```

* `baseURL` (required): Starting URL
* `maxConcurrency` (default: 5): Parallel requests
* `maxPages` (default: 50): Crawl limit

**Example:**

```bash
./bootcrawler https://example.com 10 100
```

---

## Features

* Concurrent crawling with goroutines
* Domain-restricted traversal
* Extracts:

  * Headings
  * First paragraph
  * Outgoing links
  * Image URLs
* JSON report output (`report.json`)

---

## Project Structure

* `main.go` – CLI entry point
* `crawl_page.go` – Core crawling logic
* `configure.go` – State & concurrency control
* `extract_page_*` – HTML parsing helpers
* `normalize_url.go` – URL normalization
* `json_report.go` – JSON output

---

## Output Format

```json
[
  {
    "url": "...",
    "heading": "...",
    "first_paragraph": "...",
    "outgoing_links": [],
    "image_urls": []
  }
]
```

---

## Testing

```bash
go test ./...
```

---

## Notes / Improvements

* Add URL validation & sanitization
* Implement retries and structured logging
* Respect robots.txt and rate limiting
* Expand test coverage

---

## Summary

A lightweight, concurrent Go crawler for extracting structured website data into JSON—simple, fast, and extensible.
