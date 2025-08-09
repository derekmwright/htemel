package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type InsElement struct {
	children []htemel.Node
}

// Ins creates a tag <ins> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Ins(children ...htemel.Node) *InsElement {
	node := &InsElement{
		children: children,
	}

	return node
}

func InsIf(condition bool, children ...htemel.Node) *InsElement {
	if condition {
		return Ins(children...)
	}

	return nil
}

func (e *InsElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<ins")); err != nil {
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

	if _, err := w.Write([]byte("</ins>")); err != nil {
		return err
	}

	return nil
}
