package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type TimeElement struct {
	children []htemel.Node
}

// Time creates a tag <time> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Time(children ...htemel.Node) *TimeElement {
	node := &TimeElement{
		children: children,
	}

	return node
}

func TimeIf(condition bool, children ...htemel.Node) *TimeElement {
	if condition {
		return Time(children...)
	}

	return nil
}

func (e *TimeElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<time")); err != nil {
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

	if _, err := w.Write([]byte("</time>")); err != nil {
		return err
	}

	return nil
}
