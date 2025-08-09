package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type CanvasElement struct {
	children []htemel.Node
}

// Canvas creates a tag <canvas> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Canvas(children ...htemel.Node) *CanvasElement {
	node := &CanvasElement{
		children: children,
	}

	return node
}

func CanvasIf(condition bool, children ...htemel.Node) *CanvasElement {
	if condition {
		return Canvas(children...)
	}

	return nil
}

func (e *CanvasElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<canvas")); err != nil {
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

	if _, err := w.Write([]byte("</canvas>")); err != nil {
		return err
	}

	return nil
}
