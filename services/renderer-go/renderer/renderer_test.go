package renderer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/testutil"
)

func Test_Render(t *testing.T) {
	tests := []struct {
		input string
		expected string
	}{
		{
			`# foo`, 
`<h1>foo</h1>
`,
},
		{
			`[Google](https://www.google.com/)`,
`<p><a href="https://www.google.com/">Google</a></p>
`,
},
		{
			`- bar`, 
`<ul>
<li>bar</li>
</ul>
`,
},
	}

	for _, tt := range tests {
		var buf bytes.Buffer
	    goldmark.Convert([]byte(tt.input), &buf)
	    assert.Equal(t, tt.expected, buf.String())
	}
}

func Test_Expansion_Render(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			TaskList,
		),
	)
	testutil.DoTestCaseFile(markdown, "./tasklist.txt", t, testutil.ParseCliCaseArg()...)
}
