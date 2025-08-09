package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type TheadElement struct {
	children []htemel.Node
}

// Thead creates a tag <thead> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Thead(children ...htemel.Node) *TheadElement {
	node := &TheadElement{
		children: children,
	}

	return node
}

func TheadIf(condition bool, children ...htemel.Node) *TheadElement {
	if condition {
		return Thead(children...)
	}

	return nil
}

func (e *TheadElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<thead")); err != nil {
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

	if _, err := w.Write([]byte("</thead>")); err != nil {
		return err
	}

	return nil
}
