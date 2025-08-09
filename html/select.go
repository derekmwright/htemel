package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type SelectElement struct {
	children []htemel.Node
}

// Select creates a tag <select> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Select(children ...htemel.Node) *SelectElement {
	node := &SelectElement{
		children: children,
	}

	return node
}

func SelectIf(condition bool, children ...htemel.Node) *SelectElement {
	if condition {
		return Select(children...)
	}

	return nil
}

func (e *SelectElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<select")); err != nil {
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

	if _, err := w.Write([]byte("</select>")); err != nil {
		return err
	}

	return nil
}
