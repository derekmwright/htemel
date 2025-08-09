package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type FigcaptionElement struct {
	children []htemel.Node
}

// Figcaption creates a tag <figcaption> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Figcaption(children ...htemel.Node) *FigcaptionElement {
	node := &FigcaptionElement{
		children: children,
	}

	return node
}

func FigcaptionIf(condition bool, children ...htemel.Node) *FigcaptionElement {
	if condition {
		return Figcaption(children...)
	}

	return nil
}

func (e *FigcaptionElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<figcaption")); err != nil {
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

	if _, err := w.Write([]byte("</figcaption>")); err != nil {
		return err
	}

	return nil
}
