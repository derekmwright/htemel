package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type FieldsetElement struct {
	children []htemel.Node
}

// Fieldset creates a tag <fieldset> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Fieldset(children ...htemel.Node) *FieldsetElement {
	node := &FieldsetElement{
		children: children,
	}

	return node
}

func FieldsetIf(condition bool, children ...htemel.Node) *FieldsetElement {
	if condition {
		return Fieldset(children...)
	}

	return nil
}

func (e *FieldsetElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<fieldset")); err != nil {
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

	if _, err := w.Write([]byte("</fieldset>")); err != nil {
		return err
	}

	return nil
}
