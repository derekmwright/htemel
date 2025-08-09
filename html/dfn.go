package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type DfnElement struct {
	children []htemel.Node
}

// Dfn creates a tag <dfn> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Dfn(children ...htemel.Node) *DfnElement {
	node := &DfnElement{
		children: children,
	}

	return node
}

func DfnIf(condition bool, children ...htemel.Node) *DfnElement {
	if condition {
		return Dfn(children...)
	}

	return nil
}

func (e *DfnElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<dfn")); err != nil {
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

	if _, err := w.Write([]byte("</dfn>")); err != nil {
		return err
	}

	return nil
}
