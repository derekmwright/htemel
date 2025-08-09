package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type DlElement struct {
	children []htemel.Node
}

// Dl creates a tag <dl> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Dl(children ...htemel.Node) *DlElement {
	node := &DlElement{
		children: children,
	}

	return node
}

func DlIf(condition bool, children ...htemel.Node) *DlElement {
	if condition {
		return Dl(children...)
	}

	return nil
}

func (e *DlElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<dl")); err != nil {
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

	if _, err := w.Write([]byte("</dl>")); err != nil {
		return err
	}

	return nil
}
