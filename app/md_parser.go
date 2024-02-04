package app

import (
	"bytes"
	"html"
	"github.com/yuin/goldmark"
	html_renderer "github.com/yuin/goldmark/renderer/html"
)
func ParseMarkdown(text string) string {
	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithRendererOptions(html_renderer.WithUnsafe()),
	)
	if err := md.Convert([]byte(text), &buf); err != nil {
		return html.EscapeString(text)
	}
	return buf.String()
}