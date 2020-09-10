package renderer

import (
	"bytes"
	"context"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
	my_app "github.com/wafuwafu/Hatena-Intern-2020/services/renderer-go/app"
)

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string, app my_app.App) (string, error) {
	markdown := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithASTTransformers(
				util.Prioritized(NewLinker(ctx, app), 999),
			),
		),
		goldmark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			TaskList,
		),
	)

	var buf bytes.Buffer
	if err := markdown.Convert([]byte(src), &buf); err != nil {
		return "", err
	} 
	return buf.String(), nil
}
