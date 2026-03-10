package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	header := doc.Find("h1").Text()
	if header == "" {
		header = doc.Find("h2").Text()
	}
	if header == "" {
		return ""
	}
	return header
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	main := doc.Find("main")
	var paragraph string
	if main.Length() != 0 {
		paragraph = main.Find("p").First().Text()
	} else {
		paragraph = doc.Find("p").First().Text()
	}
	return paragraph
}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {

}
