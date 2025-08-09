package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type TitleElement struct {
	children []htemel.Node
}

// Title creates a tag <title> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Title(children ...htemel.Node) *TitleElement {
	node := &TitleElement{
		children: children,
	}

	return node
}

func TitleIf(condition bool, children ...htemel.Node) *TitleElement {
	if condition {
		return Title(children...)
	}

	return nil
}

func (e *TitleElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<title")); err != nil {
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

	if _, err := w.Write([]byte("</title>")); err != nil {
		return err
	}

	return nil
}
