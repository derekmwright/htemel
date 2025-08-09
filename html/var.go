package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type VarElement struct {
	children []htemel.Node
}

// Var creates a tag <var> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Var(children ...htemel.Node) *VarElement {
	node := &VarElement{
		children: children,
	}

	return node
}

func VarIf(condition bool, children ...htemel.Node) *VarElement {
	if condition {
		return Var(children...)
	}

	return nil
}

func (e *VarElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<var")); err != nil {
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

	if _, err := w.Write([]byte("</var>")); err != nil {
		return err
	}

	return nil
}
