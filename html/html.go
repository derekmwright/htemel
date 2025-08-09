package html

import (
	"io"

	"github.com/derekmwright/htemel"
)

type HtmlElement struct {
	children []htemel.Node
}

// Html creates a tag <html> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Html(children ...htemel.Node) *HtmlElement {
	node := &HtmlElement{
		children: children,
	}

	return node
}

func HtmlIf(condition bool, children ...htemel.Node) *HtmlElement {
	if condition {
		return Html(children...)
	}

	return nil
}

func (e *HtmlElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<html")); err != nil {
		return err
	}

	// TODO: Attribute stuff here

	if _, err := w.Write([]byte(">")); err != nil {
		return err
	}

	for _, child := range e.children {
		if err := child.Render(w); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte("</html>")); err != nil {
		return err
	}

	return nil
}
