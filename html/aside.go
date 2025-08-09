package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type AsideElement struct {
	children []htemel.Node
}

// Aside creates a tag <aside> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Aside(children ...htemel.Node) *AsideElement {
	node := &AsideElement{
		children: children,
	}

	return node
}

func AsideIf(condition bool, children ...htemel.Node) *AsideElement {
	if condition {
		return Aside(children...)
	}

	return nil
}

func (e *AsideElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<aside")); err != nil {
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

	if _, err := w.Write([]byte("</aside>")); err != nil {
		return err
	}

	return nil
}
