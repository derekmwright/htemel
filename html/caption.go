package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type CaptionElement struct {
	children []htemel.Node
}

// Caption creates a tag <caption> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Caption(children ...htemel.Node) *CaptionElement {
	node := &CaptionElement{
		children: children,
	}

	return node
}

func CaptionIf(condition bool, children ...htemel.Node) *CaptionElement {
	if condition {
		return Caption(children...)
	}

	return nil
}

func (e *CaptionElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<caption")); err != nil {
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

	if _, err := w.Write([]byte("</caption>")); err != nil {
		return err
	}

	return nil
}
