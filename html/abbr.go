package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type AbbrElement struct {
	children []htemel.Node
}

// Abbr creates a tag <abbr> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Abbr(children ...htemel.Node) *AbbrElement {
	node := &AbbrElement{
		children: children,
	}

	return node
}

func AbbrIf(condition bool, children ...htemel.Node) *AbbrElement {
	if condition {
		return Abbr(children...)
	}

	return nil
}

func (e *AbbrElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<abbr")); err != nil {
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

	if _, err := w.Write([]byte("</abbr>")); err != nil {
		return err
	}

	return nil
}
