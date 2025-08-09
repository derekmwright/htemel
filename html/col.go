package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type ColElement struct {
	children []htemel.Node
}

// Col creates a tag <col> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Col(children ...htemel.Node) *ColElement {
	node := &ColElement{
		children: children,
	}

	return node
}

func ColIf(condition bool, children ...htemel.Node) *ColElement {
	if condition {
		return Col(children...)
	}

	return nil
}

func (e *ColElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<col")); err != nil {
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

	if _, err := w.Write([]byte("</col>")); err != nil {
		return err
	}

	return nil
}
