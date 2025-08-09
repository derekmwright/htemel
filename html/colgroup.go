package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type ColgroupElement struct {
	children []htemel.Node
}

// Colgroup creates a tag <colgroup> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Colgroup(children ...htemel.Node) *ColgroupElement {
	node := &ColgroupElement{
		children: children,
	}

	return node
}

func ColgroupIf(condition bool, children ...htemel.Node) *ColgroupElement {
	if condition {
		return Colgroup(children...)
	}

	return nil
}

func (e *ColgroupElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<colgroup")); err != nil {
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

	if _, err := w.Write([]byte("</colgroup>")); err != nil {
		return err
	}

	return nil
}
