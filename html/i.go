package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type IElement struct {
	children []htemel.Node
}

// I creates a tag <i> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func I(children ...htemel.Node) *IElement {
	node := &IElement{
		children: children,
	}

	return node
}

func IIf(condition bool, children ...htemel.Node) *IElement {
	if condition {
		return I(children...)
	}

	return nil
}

func (e *IElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<i")); err != nil {
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

	if _, err := w.Write([]byte("</i>")); err != nil {
		return err
	}

	return nil
}
