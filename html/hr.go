package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type HrElement struct {
	children []htemel.Node
}

// Hr creates a tag <hr> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Hr(children ...htemel.Node) *HrElement {
	node := &HrElement{
		children: children,
	}

	return node
}

func HrIf(condition bool, children ...htemel.Node) *HrElement {
	if condition {
		return Hr(children...)
	}

	return nil
}

func (e *HrElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<hr")); err != nil {
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

	if _, err := w.Write([]byte("</hr>")); err != nil {
		return err
	}

	return nil
}
