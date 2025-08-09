package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type SpanElement struct {
	children []htemel.Node
}

// Span creates a tag <span> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Span(children ...htemel.Node) *SpanElement {
	node := &SpanElement{
		children: children,
	}

	return node
}

func SpanIf(condition bool, children ...htemel.Node) *SpanElement {
	if condition {
		return Span(children...)
	}

	return nil
}

func (e *SpanElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<span")); err != nil {
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

	if _, err := w.Write([]byte("</span>")); err != nil {
		return err
	}

	return nil
}
