package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type MeterElement struct {
	children []htemel.Node
}

// Meter creates a tag <meter> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Meter(children ...htemel.Node) *MeterElement {
	node := &MeterElement{
		children: children,
	}

	return node
}

func MeterIf(condition bool, children ...htemel.Node) *MeterElement {
	if condition {
		return Meter(children...)
	}

	return nil
}

func (e *MeterElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<meter")); err != nil {
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

	if _, err := w.Write([]byte("</meter>")); err != nil {
		return err
	}

	return nil
}
