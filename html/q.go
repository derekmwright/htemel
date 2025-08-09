package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type QElement struct {
	children []htemel.Node
}

// Q creates a tag <q> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Q(children ...htemel.Node) *QElement {
	node := &QElement{
		children: children,
	}

	return node
}

func QIf(condition bool, children ...htemel.Node) *QElement {
	if condition {
		return Q(children...)
	}

	return nil
}

func (e *QElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<q")); err != nil {
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

	if _, err := w.Write([]byte("</q>")); err != nil {
		return err
	}

	return nil
}
