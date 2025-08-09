package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type BodyElement struct {
	children []htemel.Node
}

// Body creates a tag <body> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Body(children ...htemel.Node) *BodyElement {
	node := &BodyElement{
		children: children,
	}

	return node
}

func BodyIf(condition bool, children ...htemel.Node) *BodyElement {
	if condition {
		return Body(children...)
	}

	return nil
}

func (e *BodyElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<body")); err != nil {
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

	if _, err := w.Write([]byte("</body>")); err != nil {
		return err
	}

	return nil
}
