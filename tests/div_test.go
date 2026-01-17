package tests

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/go-htemel/htemel/html"
)

func TestDiv(t *testing.T) {
	var basicDiv = `<div id="test-basicDiv" class="bg-black"><p>Test</p></div>`

	t.Run("Div Creation", func(t *testing.T) {
		d := html.Div().
			Id("test-basicDiv").
			Class("bg-black").
			Children(
				html.P().Text("Test"),
			)

		buf := &bytes.Buffer{}

		if err := d.Render(buf); err != nil {
			t.Errorf("Rendering failed: %v", err)
			return
		}

		if buf.String() != basicDiv {
			t.Errorf("Expected %s, got %s", basicDiv, buf.String())
		}
	})
}

func TestDivChildOrdering(t *testing.T) {
	var div = `<div><p>Start Text<span>Spanned Text</span>End Text</p></div>`
	d := html.Div().Children(
		html.P().Text("Start Text").Children(
			html.Span().Text("Spanned Text"),
		).Text("End Text"),
	)

	buf := &bytes.Buffer{}
	if err := d.Render(buf); err != nil {
		t.Errorf("Rendering failed: %v", err)
		return
	}

	if buf.String() != div {
		t.Errorf("Expected %s, got %s", div, buf.String())
	}
}

func BenchmarkRealisticPage(b *testing.B) {
	page := html.Html(
		html.Head(
			html.Title().Text("Benchmark Test"),
			html.Meta().Charset("utf-8"),
		),
		html.Body(
			html.Div().Class("container").With(func(d *html.DivElement) {
				for i := 0; i < 100; i++ { // simulate 100-item list
					d.Children(
						html.Div().Class("item").
							Children(
								html.H2().Text("Item "+fmt.Sprint(i)),
								html.P().Text("Some description..."),
							),
					)
				}
			}),
		),
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		_ = page.Render(buf)
	}
}
