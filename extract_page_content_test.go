package main

import "testing"

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
