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
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}
	urls := []string{}
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}
		u, _ := url.Parse(href)
		if u.Scheme != "" && u.Host != "" {
			urls = append(urls, href)
		} else {
			baseURL.Path = href
			href = baseURL.String()
			urls = append(urls, href)
		}
	})
	return urls, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}
	urls := []string{}
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if !exists {
			return
		}
		u, _ := url.Parse(src)
		if u.Scheme != "" && u.Host != "" {
			urls = append(urls, src)
		} else {
			baseURL.Path = src
			src = baseURL.String()
			urls = append(urls, src)
		}
	})
	return urls, nil

}
