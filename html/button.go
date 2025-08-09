package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type ButtonElement struct {
	children []htemel.Node
}

// Button creates a tag <button> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Button(children ...htemel.Node) *ButtonElement {
	node := &ButtonElement{
		children: children,
	}

	return node
}

func ButtonIf(condition bool, children ...htemel.Node) *ButtonElement {
	if condition {
		return Button(children...)
	}

	return nil
}

func (e *ButtonElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<button")); err != nil {
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

	if _, err := w.Write([]byte("</button>")); err != nil {
		return err
	}

	return nil
}
