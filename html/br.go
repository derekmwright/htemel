package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type BrElement struct {
	children []htemel.Node
}

// Br creates a tag <br> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Br(children ...htemel.Node) *BrElement {
	node := &BrElement{
		children: children,
	}

	return node
}

func BrIf(condition bool, children ...htemel.Node) *BrElement {
	if condition {
		return Br(children...)
	}

	return nil
}

func (e *BrElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<br")); err != nil {
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

	if _, err := w.Write([]byte("</br>")); err != nil {
		return err
	}

	return nil
}
