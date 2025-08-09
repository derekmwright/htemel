package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type SmallElement struct {
	children []htemel.Node
}

// Small creates a tag <small> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Small(children ...htemel.Node) *SmallElement {
	node := &SmallElement{
		children: children,
	}

	return node
}

func SmallIf(condition bool, children ...htemel.Node) *SmallElement {
	if condition {
		return Small(children...)
	}

	return nil
}

func (e *SmallElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<small")); err != nil {
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

	if _, err := w.Write([]byte("</small>")); err != nil {
		return err
	}

	return nil
}
