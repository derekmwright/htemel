package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type DialogElement struct {
	children []htemel.Node
}

// Dialog creates a tag <dialog> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Dialog(children ...htemel.Node) *DialogElement {
	node := &DialogElement{
		children: children,
	}

	return node
}

func DialogIf(condition bool, children ...htemel.Node) *DialogElement {
	if condition {
		return Dialog(children...)
	}

	return nil
}

func (e *DialogElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<dialog")); err != nil {
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

	if _, err := w.Write([]byte("</dialog>")); err != nil {
		return err
	}

	return nil
}
