package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type ObjectElement struct {
	children []htemel.Node
}

// Object creates a tag <object> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Object(children ...htemel.Node) *ObjectElement {
	node := &ObjectElement{
		children: children,
	}

	return node
}

func ObjectIf(condition bool, children ...htemel.Node) *ObjectElement {
	if condition {
		return Object(children...)
	}

	return nil
}

func (e *ObjectElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<object")); err != nil {
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

	if _, err := w.Write([]byte("</object>")); err != nil {
		return err
	}

	return nil
}
