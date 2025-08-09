package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type ProgressElement struct {
	children []htemel.Node
}

// Progress creates a tag <progress> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Progress(children ...htemel.Node) *ProgressElement {
	node := &ProgressElement{
		children: children,
	}

	return node
}

func ProgressIf(condition bool, children ...htemel.Node) *ProgressElement {
	if condition {
		return Progress(children...)
	}

	return nil
}

func (e *ProgressElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<progress")); err != nil {
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

	if _, err := w.Write([]byte("</progress>")); err != nil {
		return err
	}

	return nil
}
