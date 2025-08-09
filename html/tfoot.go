package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type TfootElement struct {
	children []htemel.Node
}

// Tfoot creates a tag <tfoot> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Tfoot(children ...htemel.Node) *TfootElement {
	node := &TfootElement{
		children: children,
	}

	return node
}

func TfootIf(condition bool, children ...htemel.Node) *TfootElement {
	if condition {
		return Tfoot(children...)
	}

	return nil
}

func (e *TfootElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<tfoot")); err != nil {
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

	if _, err := w.Write([]byte("</tfoot>")); err != nil {
		return err
	}

	return nil
}
