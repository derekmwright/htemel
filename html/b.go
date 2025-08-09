package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type BElement struct {
	children []htemel.Node
}

// B creates a tag <b> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func B(children ...htemel.Node) *BElement {
	node := &BElement{
		children: children,
	}

	return node
}

func BIf(condition bool, children ...htemel.Node) *BElement {
	if condition {
		return B(children...)
	}

	return nil
}

func (e *BElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<b")); err != nil {
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

	if _, err := w.Write([]byte("</b>")); err != nil {
		return err
	}

	return nil
}
