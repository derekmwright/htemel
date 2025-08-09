package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type EmElement struct {
	children []htemel.Node
}

// Em creates a tag <em> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Em(children ...htemel.Node) *EmElement {
	node := &EmElement{
		children: children,
	}

	return node
}

func EmIf(condition bool, children ...htemel.Node) *EmElement {
	if condition {
		return Em(children...)
	}

	return nil
}

func (e *EmElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<em")); err != nil {
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

	if _, err := w.Write([]byte("</em>")); err != nil {
		return err
	}

	return nil
}
