package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type DatalistElement struct {
	children []htemel.Node
}

// Datalist creates a tag <datalist> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Datalist(children ...htemel.Node) *DatalistElement {
	node := &DatalistElement{
		children: children,
	}

	return node
}

func DatalistIf(condition bool, children ...htemel.Node) *DatalistElement {
	if condition {
		return Datalist(children...)
	}

	return nil
}

func (e *DatalistElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<datalist")); err != nil {
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

	if _, err := w.Write([]byte("</datalist>")); err != nil {
		return err
	}

	return nil
}
