package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type FooterElement struct {
	children []htemel.Node
}

// Footer creates a tag <footer> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Footer(children ...htemel.Node) *FooterElement {
	node := &FooterElement{
		children: children,
	}

	return node
}

func FooterIf(condition bool, children ...htemel.Node) *FooterElement {
	if condition {
		return Footer(children...)
	}

	return nil
}

func (e *FooterElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<footer")); err != nil {
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

	if _, err := w.Write([]byte("</footer>")); err != nil {
		return err
	}

	return nil
}
