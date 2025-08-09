package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type FigureElement struct {
	children []htemel.Node
}

// Figure creates a tag <figure> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Figure(children ...htemel.Node) *FigureElement {
	node := &FigureElement{
		children: children,
	}

	return node
}

func FigureIf(condition bool, children ...htemel.Node) *FigureElement {
	if condition {
		return Figure(children...)
	}

	return nil
}

func (e *FigureElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<figure")); err != nil {
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

	if _, err := w.Write([]byte("</figure>")); err != nil {
		return err
	}

	return nil
}
