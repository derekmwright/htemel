package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type HeadElement struct {
	children []htemel.Node
}

// Head creates a tag <head> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Head(children ...htemel.Node) *HeadElement {
	node := &HeadElement{
		children: children,
	}

	return node
}

func HeadIf(condition bool, children ...htemel.Node) *HeadElement {
	if condition {
		return Head(children...)
	}

	return nil
}

func (e *HeadElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<head")); err != nil {
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

	if _, err := w.Write([]byte("</head>")); err != nil {
		return err
	}

	return nil
}
