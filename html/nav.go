package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type NavElement struct {
	children []htemel.Node
}

// Nav creates a tag <nav> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Nav(children ...htemel.Node) *NavElement {
	node := &NavElement{
		children: children,
	}

	return node
}

func NavIf(condition bool, children ...htemel.Node) *NavElement {
	if condition {
		return Nav(children...)
	}

	return nil
}

func (e *NavElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<nav")); err != nil {
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

	if _, err := w.Write([]byte("</nav>")); err != nil {
		return err
	}

	return nil
}
