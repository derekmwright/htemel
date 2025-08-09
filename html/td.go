package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type TdElement struct {
	children []htemel.Node
}

// Td creates a tag <td> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Td(children ...htemel.Node) *TdElement {
	node := &TdElement{
		children: children,
	}

	return node
}

func TdIf(condition bool, children ...htemel.Node) *TdElement {
	if condition {
		return Td(children...)
	}

	return nil
}

func (e *TdElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<td")); err != nil {
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

	if _, err := w.Write([]byte("</td>")); err != nil {
		return err
	}

	return nil
}
