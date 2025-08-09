package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type SelectedcontentElement struct {
	children []htemel.Node
}

// Selectedcontent creates a tag <selectedcontent> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Selectedcontent(children ...htemel.Node) *SelectedcontentElement {
	node := &SelectedcontentElement{
		children: children,
	}

	return node
}

func SelectedcontentIf(condition bool, children ...htemel.Node) *SelectedcontentElement {
	if condition {
		return Selectedcontent(children...)
	}

	return nil
}

func (e *SelectedcontentElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<selectedcontent")); err != nil {
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

	if _, err := w.Write([]byte("</selectedcontent>")); err != nil {
		return err
	}

	return nil
}
