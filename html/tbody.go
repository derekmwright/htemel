package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type TbodyElement struct {
	children []htemel.Node
}

// Tbody creates a tag <tbody> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Tbody(children ...htemel.Node) *TbodyElement {
	node := &TbodyElement{
		children: children,
	}

	return node
}

func TbodyIf(condition bool, children ...htemel.Node) *TbodyElement {
	if condition {
		return Tbody(children...)
	}

	return nil
}

func (e *TbodyElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<tbody")); err != nil {
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

	if _, err := w.Write([]byte("</tbody>")); err != nil {
		return err
	}

	return nil
}
