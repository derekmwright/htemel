package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type PictureElement struct {
	children []htemel.Node
}

// Picture creates a tag <picture> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Picture(children ...htemel.Node) *PictureElement {
	node := &PictureElement{
		children: children,
	}

	return node
}

func PictureIf(condition bool, children ...htemel.Node) *PictureElement {
	if condition {
		return Picture(children...)
	}

	return nil
}

func (e *PictureElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<picture")); err != nil {
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

	if _, err := w.Write([]byte("</picture>")); err != nil {
		return err
	}

	return nil
}
