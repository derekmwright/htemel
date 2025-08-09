package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type RubyElement struct {
	children []htemel.Node
}

// Ruby creates a tag <ruby> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Ruby(children ...htemel.Node) *RubyElement {
	node := &RubyElement{
		children: children,
	}

	return node
}

func RubyIf(condition bool, children ...htemel.Node) *RubyElement {
	if condition {
		return Ruby(children...)
	}

	return nil
}

func (e *RubyElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<ruby")); err != nil {
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

	if _, err := w.Write([]byte("</ruby>")); err != nil {
		return err
	}

	return nil
}
