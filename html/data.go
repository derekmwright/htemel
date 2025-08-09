package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type DataElement struct {
	children []htemel.Node
}

// Data creates a tag <data> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Data(children ...htemel.Node) *DataElement {
	node := &DataElement{
		children: children,
	}

	return node
}

func DataIf(condition bool, children ...htemel.Node) *DataElement {
	if condition {
		return Data(children...)
	}

	return nil
}

func (e *DataElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<data")); err != nil {
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

	if _, err := w.Write([]byte("</data>")); err != nil {
		return err
	}

	return nil
}
