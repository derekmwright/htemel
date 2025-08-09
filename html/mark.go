package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type MarkElement struct {
	children []htemel.Node
}

// Mark creates a tag <mark> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Mark(children ...htemel.Node) *MarkElement {
	node := &MarkElement{
		children: children,
	}

	return node
}

func MarkIf(condition bool, children ...htemel.Node) *MarkElement {
	if condition {
		return Mark(children...)
	}

	return nil
}

func (e *MarkElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<mark")); err != nil {
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

	if _, err := w.Write([]byte("</mark>")); err != nil {
		return err
	}

	return nil
}
