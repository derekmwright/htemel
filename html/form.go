package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type FormElement struct {
	children []htemel.Node
}

// Form creates a tag <form> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Form(children ...htemel.Node) *FormElement {
	node := &FormElement{
		children: children,
	}

	return node
}

func FormIf(condition bool, children ...htemel.Node) *FormElement {
	if condition {
		return Form(children...)
	}

	return nil
}

func (e *FormElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<form")); err != nil {
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

	if _, err := w.Write([]byte("</form>")); err != nil {
		return err
	}

	return nil
}
