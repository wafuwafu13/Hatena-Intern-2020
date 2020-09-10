package renderer

import (
	"context"

	my_app "github.com/wafuwafu/Hatena-Intern-2020/services/renderer-go/app"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type linker struct {
	ctx context.Context
	app my_app.App
}

func newLinker(app my_app.App) *linker {
	return &linker{app: app}
}

func (l *linker) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
	    if node, ok := node.(*ast.Link); ok && entering && node.ChildCount() == 0 {
	    	title, err := l.app.Fetch(l.ctx, string(node.Destination))
	    	if err != nil {
	    		title = string(node.Destination)
	    	}
	    	node.AppendChild(node, ast.NewString([]byte(title)))
	    }
	    return ast.WalkContinue, nil
    })
}