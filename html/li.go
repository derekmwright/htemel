package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type LiElement struct {
	children []htemel.Node
}

// Li creates a tag <li> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Li(children ...htemel.Node) *LiElement {
	node := &LiElement{
		children: children,
	}

	return node
}

func LiIf(condition bool, children ...htemel.Node) *LiElement {
	if condition {
		return Li(children...)
	}

	return nil
}

func (e *LiElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<li")); err != nil {
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

	if _, err := w.Write([]byte("</li>")); err != nil {
		return err
	}

	return nil
}
