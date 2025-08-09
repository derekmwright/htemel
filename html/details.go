package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type DetailsElement struct {
	children []htemel.Node
}

// Details creates a tag <details> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Details(children ...htemel.Node) *DetailsElement {
	node := &DetailsElement{
		children: children,
	}

	return node
}

func DetailsIf(condition bool, children ...htemel.Node) *DetailsElement {
	if condition {
		return Details(children...)
	}

	return nil
}

func (e *DetailsElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<details")); err != nil {
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

	if _, err := w.Write([]byte("</details>")); err != nil {
		return err
	}

	return nil
}
