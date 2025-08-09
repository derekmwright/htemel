package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type DdElement struct {
	children []htemel.Node
}

// Dd creates a tag <dd> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Dd(children ...htemel.Node) *DdElement {
	node := &DdElement{
		children: children,
	}

	return node
}

func DdIf(condition bool, children ...htemel.Node) *DdElement {
	if condition {
		return Dd(children...)
	}

	return nil
}

func (e *DdElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<dd")); err != nil {
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

	if _, err := w.Write([]byte("</dd>")); err != nil {
		return err
	}

	return nil
}
