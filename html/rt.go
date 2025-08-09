package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type RtElement struct {
	children []htemel.Node
}

// Rt creates a tag <rt> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Rt(children ...htemel.Node) *RtElement {
	node := &RtElement{
		children: children,
	}

	return node
}

func RtIf(condition bool, children ...htemel.Node) *RtElement {
	if condition {
		return Rt(children...)
	}

	return nil
}

func (e *RtElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<rt")); err != nil {
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

	if _, err := w.Write([]byte("</rt>")); err != nil {
		return err
	}

	return nil
}
