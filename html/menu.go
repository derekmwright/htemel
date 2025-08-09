package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type MenuElement struct {
	children []htemel.Node
}

// Menu creates a tag <menu> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Menu(children ...htemel.Node) *MenuElement {
	node := &MenuElement{
		children: children,
	}

	return node
}

func MenuIf(condition bool, children ...htemel.Node) *MenuElement {
	if condition {
		return Menu(children...)
	}

	return nil
}

func (e *MenuElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<menu")); err != nil {
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

	if _, err := w.Write([]byte("</menu>")); err != nil {
		return err
	}

	return nil
}
