package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type TrElement struct {
	children []htemel.Node
}

// Tr creates a tag <tr> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Tr(children ...htemel.Node) *TrElement {
	node := &TrElement{
		children: children,
	}

	return node
}

func TrIf(condition bool, children ...htemel.Node) *TrElement {
	if condition {
		return Tr(children...)
	}

	return nil
}

func (e *TrElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<tr")); err != nil {
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

	if _, err := w.Write([]byte("</tr>")); err != nil {
		return err
	}

	return nil
}
