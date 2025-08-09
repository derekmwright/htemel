package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type HeaderElement struct {
	children []htemel.Node
}

// Header creates a tag <header> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Header(children ...htemel.Node) *HeaderElement {
	node := &HeaderElement{
		children: children,
	}

	return node
}

func HeaderIf(condition bool, children ...htemel.Node) *HeaderElement {
	if condition {
		return Header(children...)
	}

	return nil
}

func (e *HeaderElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<header")); err != nil {
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

	if _, err := w.Write([]byte("</header>")); err != nil {
		return err
	}

	return nil
}
