package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type PElement struct {
	children []htemel.Node
}

// P creates a tag <p> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func P(children ...htemel.Node) *PElement {
	node := &PElement{
		children: children,
	}

	return node
}

func PIf(condition bool, children ...htemel.Node) *PElement {
	if condition {
		return P(children...)
	}

	return nil
}

func (e *PElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<p")); err != nil {
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

	if _, err := w.Write([]byte("</p>")); err != nil {
		return err
	}

	return nil
}
