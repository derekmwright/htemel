package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type PreElement struct {
	children []htemel.Node
}

// Pre creates a tag <pre> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Pre(children ...htemel.Node) *PreElement {
	node := &PreElement{
		children: children,
	}

	return node
}

func PreIf(condition bool, children ...htemel.Node) *PreElement {
	if condition {
		return Pre(children...)
	}

	return nil
}

func (e *PreElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<pre")); err != nil {
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

	if _, err := w.Write([]byte("</pre>")); err != nil {
		return err
	}

	return nil
}
