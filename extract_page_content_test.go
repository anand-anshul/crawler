package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetHeadingFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputHTML string
		expected  string
	}{
		{
			name: "extract h1",
			inputHTML: `
<html>
  <body>
    <h1>Welcome to Boot.dev</h1>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
  </body>
</html>
			`,
			expected: "Welcome to Boot.dev",
		},
		// add more test cases here
		{
			name: "extract h2",
			inputHTML: `
<html>
  <body>
    <h2>Welcome to Boot.dev</h2>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
  </body>
</html>
			`,
			expected: "Welcome to Boot.dev",
		},
		{
			name: "no h1 or h2",
			inputHTML: `
<html>
  <body>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
  </body>
</html>
			`,
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.inputHTML)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputHTML string
		expected  string
	}{
		{
			name: "extract p in main",
			inputHTML: `
<html>
  <body>
    <h1>Welcome to Boot.dev</h1>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
  </body>
</html>
			`,
			expected: "Learn to code by building real projects.",
		},
		// add more test cases here
		{
			name: "extract p no main",
			inputHTML: `
<html>
  <body>
    <h2>Welcome to Boot.dev</h2>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
  </body>
</html>
			`,
			expected: "Learn to code by building real projects.",
		},
		{
			name: "no p",
			inputHTML: `
<html>
  <body>
    <main>
      
    </main>
  </body>
</html>
			`,
			expected: "",
		},
		{
			name: "Main paragraph.",
			inputHTML: `
<html>
	<body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body>
</html>
			`,
			expected: "Main paragraph.",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputHTML)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	inputURL := "https://crawler-test.com"

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Fatalf("couldn't parse input URL: %v", err)
	}

	tests := []struct {
		name      string
		inputHTML string
		expected  []string
	}{
		{
			name: "single absolute link",
			inputHTML: `<html><body>
				<a href="https://crawler-test.com">Boot.dev</a>
			</body></html>`,
			expected: []string{"https://crawler-test.com"},
		},
		{
			name: "multiple absolute links",
			inputHTML: `<html><body>
				<a href="https://crawler-test.com/page1">Page1</a>
				<a href="https://crawler-test.com/page2">Page2</a>
			</body></html>`,
			expected: []string{
				"https://crawler-test.com/page1",
				"https://crawler-test.com/page2",
			},
		},
		{
			name: "relative link",
			inputHTML: `<html><body>
				<a href="/about">About</a>
			</body></html>`,
			expected: []string{
				"https://crawler-test.com/about",
			},
		},
		{
			name: "mixed relative and absolute",
			inputHTML: `<html><body>
				<a href="/about">About</a>
				<a href="https://blog.boot.dev">Blog</a>
			</body></html>`,
			expected: []string{
				"https://crawler-test.com/about",
				"https://blog.boot.dev",
			},
		},
		{
			name: "no href attribute",
			inputHTML: `<html><body>
				<a>Broken link</a>
			</body></html>`,
			expected: []string{},
		},
		{
			name: "nested element inside anchor",
			inputHTML: `<html><body>
				<a href="https://crawler-test.com">
					<span>Boot.dev</span>
				</a>
			</body></html>`,
			expected: []string{
				"https://crawler-test.com",
			},
		},
		{
			name: "duplicate links",
			inputHTML: `<html><body>
				<a href="/about">About</a>
				<a href="/about">About again</a>
			</body></html>`,
			expected: []string{
				"https://crawler-test.com/about",
				"https://crawler-test.com/about",
			},
		},
		{
			name:      "empty html",
			inputHTML: `<html><body></body></html>`,
			expected:  []string{},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputHTML, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestGetImagesFromHTML(t *testing.T) {
	inputURL := "https://crawler-test.com"

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Fatalf("couldn't parse input URL: %v", err)
	}

	tests := []struct {
		name      string
		inputHTML string
		expected  []string
	}{
		{
			name: "relative image",
			inputHTML: `<html><body>
				<img src="/logo.png" alt="Logo">
			</body></html>`,
			expected: []string{
				"https://crawler-test.com/logo.png",
			},
		},
		{
			name: "absolute image",
			inputHTML: `<html><body>
				<img src="https://cdn.crawler-test.com/logo.png">
			</body></html>`,
			expected: []string{
				"https://cdn.crawler-test.com/logo.png",
			},
		},
		{
			name: "multiple images",
			inputHTML: `<html><body>
				<img src="/logo.png">
				<img src="/banner.png">
			</body></html>`,
			expected: []string{
				"https://crawler-test.com/logo.png",
				"https://crawler-test.com/banner.png",
			},
		},
		{
			name: "mixed relative and absolute",
			inputHTML: `<html><body>
				<img src="/logo.png">
				<img src="https://cdn.crawler-test.com/banner.png">
			</body></html>`,
			expected: []string{
				"https://crawler-test.com/logo.png",
				"https://cdn.crawler-test.com/banner.png",
			},
		},
		{
			name: "missing src attribute",
			inputHTML: `<html><body>
				<img alt="logo">
			</body></html>`,
			expected: []string{},
		},
		{
			name: "no images",
			inputHTML: `<html><body>
				<h1>Hello</h1>
				<p>No images here</p>
			</body></html>`,
			expected: []string{},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getImagesFromHTML(tc.inputHTML, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
