package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type UlElement struct {
	children []htemel.Node
}

// Ul creates a tag <ul> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Ul(children ...htemel.Node) *UlElement {
	node := &UlElement{
		children: children,
	}

	return node
}

func UlIf(condition bool, children ...htemel.Node) *UlElement {
	if condition {
		return Ul(children...)
	}

	return nil
}

func (e *UlElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<ul")); err != nil {
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

	if _, err := w.Write([]byte("</ul>")); err != nil {
		return err
	}

	return nil
}
