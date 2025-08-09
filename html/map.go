package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type MapElement struct {
	children []htemel.Node
}

// Map creates a tag <map> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Map(children ...htemel.Node) *MapElement {
	node := &MapElement{
		children: children,
	}

	return node
}

func MapIf(condition bool, children ...htemel.Node) *MapElement {
	if condition {
		return Map(children...)
	}

	return nil
}

func (e *MapElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<map")); err != nil {
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

	if _, err := w.Write([]byte("</map>")); err != nil {
		return err
	}

	return nil
}
