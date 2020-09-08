package renderer

import (
	"fmt"
	"bytes"
	"context"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"regexp"
)

var markdown = goldmark.New(
	goldmark.WithRendererOptions(
		html.WithXHTML(),
		html.WithUnsafe(),
	),
	goldmark.WithExtensions(
		TaskList,
	),
)

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
	var buf bytes.Buffer
	if err := markdown.Convert([]byte(src), &buf); err != nil {
		return "", err
	} 
	return buf.String(), nil
}

var taskListRegexp = regexp.MustCompile(`^\[([\sxXyYzZ])\]\s*`)

type taskCheckBoxParser struct {
}

var defaultTaskCheckBoxParser = &taskCheckBoxParser{}

// NewTaskCheckBoxParser returns a new  InlineParser that can parse
// checkboxes in list items.
// This parser must take precedence over the parser.LinkParser.
func NewTaskCheckBoxParser() parser.InlineParser {
	return defaultTaskCheckBoxParser
}

func (s *taskCheckBoxParser) Trigger() []byte {
	return []byte{'['}
}

func (s *taskCheckBoxParser) Parse(parent gast.Node, block text.Reader, pc parser.Context) gast.Node {
	// Given AST structure must be like
	// - List
	//   - ListItem         : parent.Parent
	//     - TextBlock      : parent
	//       (current line)
	if parent.Parent() == nil || parent.Parent().FirstChild() != parent {
		return nil
	}

	if _, ok := parent.Parent().(*gast.ListItem); !ok {
		return nil
	}
	line, _ := block.PeekLine()
	m := taskListRegexp.FindSubmatchIndex(line)
	if m == nil {
		return nil
	}
	value := line[m[2]:m[3]][0]
	block.Advance(m[1])
	checked := value == 'x' || value == 'X'
	ranged := value == 'y' || value == 'Y'
	colored := value == 'z' || value == 'Z'
	return NewTaskCheckBox(checked, ranged, colored)
}

func (s *taskCheckBoxParser) CloseBlock(parent gast.Node, pc parser.Context) {
	// nothing to do
}

// TaskCheckBoxHTMLRenderer is a renderer.NodeRenderer implementation that
// renders checkboxes in list items.
type TaskCheckBoxHTMLRenderer struct {
	html.Config
}

// NewTaskCheckBoxHTMLRenderer returns a new TaskCheckBoxHTMLRenderer.
func NewTaskCheckBoxHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &TaskCheckBoxHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *TaskCheckBoxHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindTaskCheckBox, r.renderTaskCheckBox)
}

func (r *TaskCheckBoxHTMLRenderer) renderTaskCheckBox(w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	if !entering {
		return gast.WalkContinue, nil
	}
	n := node.(*TaskCheckBox)

	if n.IsChecked {
		w.WriteString(`<input type="checkbox"`)
	} else if n.IsRanged {
		w.WriteString(`<input type="range"`)
	} else if n.IsColored {
		w.WriteString(`<input type="color"`)
	}else {
		w.WriteString(`<input disabled="" type="checkbox"`)
	}
	if r.XHTML {
		w.WriteString(" /> ")
	} else {
		w.WriteString("> ")
	}
	return gast.WalkContinue, nil
}

type taskList struct {
}

// TaskList is an extension that allow you to use GFM task lists.
var TaskList = &taskList{}

func (e *taskList) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewTaskCheckBoxParser(), 0),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewTaskCheckBoxHTMLRenderer(), 500),
	))
}

// returns a new TaskCheckBox node.

type TaskCheckBox struct {
	gast.BaseInline
	IsChecked bool
	IsRanged bool
	IsColored bool
}

// Dump implements Node.Dump.
func (n *TaskCheckBox) Dump(source []byte, level int) {
	m := map[string]string{
		"Checked": fmt.Sprintf("%v", n.IsChecked),
	}
	gast.DumpHelper(n, source, level, m, nil)
}

// KindTaskCheckBox is a NodeKind of the TaskCheckBox node.
var KindTaskCheckBox = gast.NewNodeKind("TaskCheckBox")

// Kind implements Node.Kind.
func (n *TaskCheckBox) Kind() gast.NodeKind {
	return KindTaskCheckBox
}

// NewTaskCheckBox returns a new TaskCheckBox node.
func NewTaskCheckBox(checked bool, ranged bool, colored bool) *TaskCheckBox {
	return &TaskCheckBox{
		IsChecked: checked,
		IsRanged: ranged,
		IsColored: colored,
	}
}