package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type CiteElement struct {
	children []htemel.Node
}

// Cite creates a tag <cite> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Cite(children ...htemel.Node) *CiteElement {
	node := &CiteElement{
		children: children,
	}

	return node
}

func CiteIf(condition bool, children ...htemel.Node) *CiteElement {
	if condition {
		return Cite(children...)
	}

	return nil
}

func (e *CiteElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<cite")); err != nil {
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

	if _, err := w.Write([]byte("</cite>")); err != nil {
		return err
	}

	return nil
}
