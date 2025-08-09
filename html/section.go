package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type SectionElement struct {
	children []htemel.Node
}

// Section creates a tag <section> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Section(children ...htemel.Node) *SectionElement {
	node := &SectionElement{
		children: children,
	}

	return node
}

func SectionIf(condition bool, children ...htemel.Node) *SectionElement {
	if condition {
		return Section(children...)
	}

	return nil
}

func (e *SectionElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<section")); err != nil {
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

	if _, err := w.Write([]byte("</section>")); err != nil {
		return err
	}

	return nil
}
