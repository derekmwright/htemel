package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type MainElement struct {
	children []htemel.Node
}

// Main creates a tag <main> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Main(children ...htemel.Node) *MainElement {
	node := &MainElement{
		children: children,
	}

	return node
}

func MainIf(condition bool, children ...htemel.Node) *MainElement {
	if condition {
		return Main(children...)
	}

	return nil
}

func (e *MainElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<main")); err != nil {
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

	if _, err := w.Write([]byte("</main>")); err != nil {
		return err
	}

	return nil
}
