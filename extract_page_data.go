package main

import (
	"net/url"
)

type PageData struct {
	URL            string   `json:"url"`
	Heading        string   `json:"heading"`
	FirstParagraph string   `json:"first_paragraph"`
	OutgoingLinks  []string `json:"outgoing_links"`
	ImageURLs      []string `json:"image_urls"`
}

func extractPageData(html, pageURL string) PageData {
	heading := getHeadingFromHTML(html)
	firstParagraph := getFirstParagraphFromHTML(html)
	u, err := url.Parse(pageURL)
	if err != nil {
		return PageData{}
	}
	outgoingLinks, err := getURLsFromHTML(html, u)
	if err != nil {
		return PageData{}
	}
	imageURLs, err := getImagesFromHTML(html, u)
	if err != nil {
		return PageData{}
	}
	pageData := PageData{
		URL:            pageURL,
		Heading:        heading,
		FirstParagraph: firstParagraph,
		OutgoingLinks:  outgoingLinks,
		ImageURLs:      imageURLs,
	}
	return pageData
}
