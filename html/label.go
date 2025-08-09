package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type LabelElement struct {
	children []htemel.Node
}

// Label creates a tag <label> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Label(children ...htemel.Node) *LabelElement {
	node := &LabelElement{
		children: children,
	}

	return node
}

func LabelIf(condition bool, children ...htemel.Node) *LabelElement {
	if condition {
		return Label(children...)
	}

	return nil
}

func (e *LabelElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<label")); err != nil {
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

	if _, err := w.Write([]byte("</label>")); err != nil {
		return err
	}

	return nil
}
