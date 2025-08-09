package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type BdoElement struct {
	children []htemel.Node
}

// Bdo creates a tag <bdo> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Bdo(children ...htemel.Node) *BdoElement {
	node := &BdoElement{
		children: children,
	}

	return node
}

func BdoIf(condition bool, children ...htemel.Node) *BdoElement {
	if condition {
		return Bdo(children...)
	}

	return nil
}

func (e *BdoElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<bdo")); err != nil {
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

	if _, err := w.Write([]byte("</bdo>")); err != nil {
		return err
	}

	return nil
}
