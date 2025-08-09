package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type AreaElement struct {
	children []htemel.Node
}

// Area creates a tag <area> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Area(children ...htemel.Node) *AreaElement {
	node := &AreaElement{
		children: children,
	}

	return node
}

func AreaIf(condition bool, children ...htemel.Node) *AreaElement {
	if condition {
		return Area(children...)
	}

	return nil
}

func (e *AreaElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<area")); err != nil {
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

	if _, err := w.Write([]byte("</area>")); err != nil {
		return err
	}

	return nil
}
