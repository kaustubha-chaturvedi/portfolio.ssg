package app

import (
	"bytes"
	"html"

	mathjax "github.com/litao91/goldmark-mathjax"
	enclave "github.com/quail-ink/goldmark-enclave"
	"github.com/yuin/goldmark"
	extensions "github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	html_renderer "github.com/yuin/goldmark/renderer/html"
)
func ParseMarkdown(text string) string {
	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(
			mathjax.MathJax,
			enclave.New(&enclave.Config{
				DefaultImageAltPrefix: "Image",
			}),
			extensions.GFM,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html_renderer.WithXHTML(),
		),
	)
	if err := md.Convert([]byte(text), &buf); err != nil {
		return html.EscapeString(text)
	}
	return buf.String()
}