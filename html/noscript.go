package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type NoscriptElement struct {
	children []htemel.Node
}

// Noscript creates a tag <noscript> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Noscript(children ...htemel.Node) *NoscriptElement {
	node := &NoscriptElement{
		children: children,
	}

	return node
}

func NoscriptIf(condition bool, children ...htemel.Node) *NoscriptElement {
	if condition {
		return Noscript(children...)
	}

	return nil
}

func (e *NoscriptElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<noscript")); err != nil {
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

	if _, err := w.Write([]byte("</noscript>")); err != nil {
		return err
	}

	return nil
}
