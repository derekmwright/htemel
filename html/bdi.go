package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type BdiElement struct {
	children []htemel.Node
}

// Bdi creates a tag <bdi> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Bdi(children ...htemel.Node) *BdiElement {
	node := &BdiElement{
		children: children,
	}

	return node
}

func BdiIf(condition bool, children ...htemel.Node) *BdiElement {
	if condition {
		return Bdi(children...)
	}

	return nil
}

func (e *BdiElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<bdi")); err != nil {
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

	if _, err := w.Write([]byte("</bdi>")); err != nil {
		return err
	}

	return nil
}
