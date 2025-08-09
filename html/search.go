package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type SearchElement struct {
	children []htemel.Node
}

// Search creates a tag <search> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Search(children ...htemel.Node) *SearchElement {
	node := &SearchElement{
		children: children,
	}

	return node
}

func SearchIf(condition bool, children ...htemel.Node) *SearchElement {
	if condition {
		return Search(children...)
	}

	return nil
}

func (e *SearchElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<search")); err != nil {
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

	if _, err := w.Write([]byte("</search>")); err != nil {
		return err
	}

	return nil
}
