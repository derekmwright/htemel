package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type AElement struct {
	children []htemel.Node
}

// A creates a tag <a> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func A(children ...htemel.Node) *AElement {
	node := &AElement{
		children: children,
	}

	return node
}

func AIf(condition bool, children ...htemel.Node) *AElement {
	if condition {
		return A(children...)
	}

	return nil
}

func (e *AElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<a")); err != nil {
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

	if _, err := w.Write([]byte("</a>")); err != nil {
		return err
	}

	return nil
}
