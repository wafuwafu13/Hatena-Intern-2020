package renderer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuin/goldmark"
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
