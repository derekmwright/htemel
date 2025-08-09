package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type BaseElement struct {
	children []htemel.Node
}

// Base creates a tag <base> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Base(children ...htemel.Node) *BaseElement {
	node := &BaseElement{
		children: children,
	}

	return node
}

func BaseIf(condition bool, children ...htemel.Node) *BaseElement {
	if condition {
		return Base(children...)
	}

	return nil
}

func (e *BaseElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<base")); err != nil {
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

	if _, err := w.Write([]byte("</base>")); err != nil {
		return err
	}

	return nil
}
