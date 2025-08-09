package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type ScriptElement struct {
	children []htemel.Node
}

// Script creates a tag <script> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Script(children ...htemel.Node) *ScriptElement {
	node := &ScriptElement{
		children: children,
	}

	return node
}

func ScriptIf(condition bool, children ...htemel.Node) *ScriptElement {
	if condition {
		return Script(children...)
	}

	return nil
}

func (e *ScriptElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<script")); err != nil {
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

	if _, err := w.Write([]byte("</script>")); err != nil {
		return err
	}

	return nil
}
