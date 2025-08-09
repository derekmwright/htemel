package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type OlElement struct {
	children []htemel.Node
}

// Ol creates a tag <ol> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Ol(children ...htemel.Node) *OlElement {
	node := &OlElement{
		children: children,
	}

	return node
}

func OlIf(condition bool, children ...htemel.Node) *OlElement {
	if condition {
		return Ol(children...)
	}

	return nil
}

func (e *OlElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<ol")); err != nil {
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

	if _, err := w.Write([]byte("</ol>")); err != nil {
		return err
	}

	return nil
}
