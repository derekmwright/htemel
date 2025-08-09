package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type DivElement struct {
	children []htemel.Node
}

// Div creates a tag <div> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Div(children ...htemel.Node) *DivElement {
	node := &DivElement{
		children: children,
	}

	return node
}

func DivIf(condition bool, children ...htemel.Node) *DivElement {
	if condition {
		return Div(children...)
	}

	return nil
}

func (e *DivElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<div")); err != nil {
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

	if _, err := w.Write([]byte("</div>")); err != nil {
		return err
	}

	return nil
}
