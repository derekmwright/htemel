package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type RpElement struct {
	children []htemel.Node
}

// Rp creates a tag <rp> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Rp(children ...htemel.Node) *RpElement {
	node := &RpElement{
		children: children,
	}

	return node
}

func RpIf(condition bool, children ...htemel.Node) *RpElement {
	if condition {
		return Rp(children...)
	}

	return nil
}

func (e *RpElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<rp")); err != nil {
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

	if _, err := w.Write([]byte("</rp>")); err != nil {
		return err
	}

	return nil
}
