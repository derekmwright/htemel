package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type SummaryElement struct {
	children []htemel.Node
}

// Summary creates a tag <summary> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Summary(children ...htemel.Node) *SummaryElement {
	node := &SummaryElement{
		children: children,
	}

	return node
}

func SummaryIf(condition bool, children ...htemel.Node) *SummaryElement {
	if condition {
		return Summary(children...)
	}

	return nil
}

func (e *SummaryElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<summary")); err != nil {
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

	if _, err := w.Write([]byte("</summary>")); err != nil {
		return err
	}

	return nil
}
