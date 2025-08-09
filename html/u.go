package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type UElement struct {
	children []htemel.Node
}

// U creates a tag <u> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func U(children ...htemel.Node) *UElement {
	node := &UElement{
		children: children,
	}

	return node
}

func UIf(condition bool, children ...htemel.Node) *UElement {
	if condition {
		return U(children...)
	}

	return nil
}

func (e *UElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<u")); err != nil {
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

	if _, err := w.Write([]byte("</u>")); err != nil {
		return err
	}

	return nil
}
