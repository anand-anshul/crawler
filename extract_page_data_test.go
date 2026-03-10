package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputHTML string
		expected  PageData
	}{
		{
			name:     "basic extraction",
			inputURL: "https://crawler-test.com",
			inputHTML: `<html><body>
				<h1>Test Title</h1>
				<p>This is the first paragraph.</p>
				<a href="/link1">Link 1</a>
				<img src="/image1.jpg">
			</body></html>`,
			expected: PageData{
				URL:            "https://crawler-test.com",
				Heading:        "Test Title",
				FirstParagraph: "This is the first paragraph.",
				OutgoingLinks: []string{
					"https://crawler-test.com/link1",
				},
				ImageURLs: []string{
					"https://crawler-test.com/image1.jpg",
				},
			},
		},
		{
			name:     "multiple links and images",
			inputURL: "https://crawler-test.com",
			inputHTML: `<html><body>
				<h1>Page Title</h1>
				<p>Intro paragraph.</p>
				<a href="/link1">Link 1</a>
				<a href="/link2">Link 2</a>
				<img src="/img1.png">
				<img src="/img2.png">
			</body></html>`,
			expected: PageData{
				URL:            "https://crawler-test.com",
				Heading:        "Page Title",
				FirstParagraph: "Intro paragraph.",
				OutgoingLinks: []string{
					"https://crawler-test.com/link1",
					"https://crawler-test.com/link2",
				},
				ImageURLs: []string{
					"https://crawler-test.com/img1.png",
					"https://crawler-test.com/img2.png",
				},
			},
		},
		{
			name:     "absolute urls",
			inputURL: "https://crawler-test.com",
			inputHTML: `<html><body>
				<h1>Title</h1>
				<p>Paragraph text.</p>
				<a href="https://example.com/page">External</a>
				<img src="https://cdn.site.com/img.jpg">
			</body></html>`,
			expected: PageData{
				URL:            "https://crawler-test.com",
				Heading:        "Title",
				FirstParagraph: "Paragraph text.",
				OutgoingLinks: []string{
					"https://example.com/page",
				},
				ImageURLs: []string{
					"https://cdn.site.com/img.jpg",
				},
			},
		},
		{
			name:     "missing heading",
			inputURL: "https://crawler-test.com",
			inputHTML: `<html><body>
				<p>First paragraph only.</p>
				<a href="/link1">Link</a>
			</body></html>`,
			expected: PageData{
				URL:            "https://crawler-test.com",
				Heading:        "",
				FirstParagraph: "First paragraph only.",
				OutgoingLinks: []string{
					"https://crawler-test.com/link1",
				},
				ImageURLs: []string{},
			},
		},
		{
			name:     "no links or images",
			inputURL: "https://crawler-test.com",
			inputHTML: `<html><body>
				<h1>Heading</h1>
				<p>Just text.</p>
			</body></html>`,
			expected: PageData{
				URL:            "https://crawler-test.com",
				Heading:        "Heading",
				FirstParagraph: "Just text.",
				OutgoingLinks:  []string{},
				ImageURLs:      []string{},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual := extractPageData(tc.inputHTML, tc.inputURL)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %+v, got %+v", tc.expected, actual)
			}
		})
	}
}
