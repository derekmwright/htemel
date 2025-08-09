package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type StyleElement struct {
	children []htemel.Node
}

// Style creates a tag <style> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Style(children ...htemel.Node) *StyleElement {
	node := &StyleElement{
		children: children,
	}

	return node
}

func StyleIf(condition bool, children ...htemel.Node) *StyleElement {
	if condition {
		return Style(children...)
	}

	return nil
}

func (e *StyleElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<style")); err != nil {
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

	if _, err := w.Write([]byte("</style>")); err != nil {
		return err
	}

	return nil
}
