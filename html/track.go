package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type TrackElement struct {
	children []htemel.Node
}

// Track creates a tag <track> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Track(children ...htemel.Node) *TrackElement {
	node := &TrackElement{
		children: children,
	}

	return node
}

func TrackIf(condition bool, children ...htemel.Node) *TrackElement {
	if condition {
		return Track(children...)
	}

	return nil
}

func (e *TrackElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<track")); err != nil {
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

	if _, err := w.Write([]byte("</track>")); err != nil {
		return err
	}

	return nil
}
