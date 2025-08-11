package htemel

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Node defines the interface that must be implemented in order to render elements.
type Node interface {
	Render(w io.Writer) error
}

// GroupElement is a struct that backs the Group function.
type GroupElement struct {
	children []Node
}

// Group is a generic wrapper that can be used to wrap one or more elements that may not have a suitable parent type.
func Group(children ...Node) *GroupElement {
	return &GroupElement{children}
}

// Render implements the Node interface by calling Render on all child nodes.
func (e *GroupElement) Render(w io.Writer) error {
	for _, child := range e.children {
		if err := child.Render(w); err != nil {
			return err
		}
	}

	return nil
}

// GenericElement struct backs the Generic element functions.
type GenericElement struct {
	tag      string
	attrs    map[string]any
	void     bool
	children []Node
}

// Generic element is provided as an escape-hatch for when the provided generated elements are not sufficient.
// Attributes can be passed as a map of strings with "any" type.
// The underlying type should implement the fmt.Stringer interface for predictable rendering.
func Generic(tag string, attrs map[string]any, children ...Node) *GenericElement {
	return &GenericElement{
		tag:      tag,
		attrs:    attrs,
		void:     false,
		children: children,
	}
}

// GenericVoid element is provided as an escape-hatch for when the provided generated elements are not sufficient.
// Attributes can be passed as a map of strings with "any" type.
// The underlying type should implement the fmt.Stringer interface for predictable rendering.
//
// Void elements are self-closing and therefore do not permit children.
func GenericVoid(tag string, attrs map[string]any) *GenericElement {
	return &GenericElement{
		tag:      tag,
		attrs:    attrs,
		void:     true,
		children: nil,
	}
}

func (e *GenericElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<" + e.tag)); err != nil {
		return err
	}

	for key, val := range e.attrs {
		if val == nil {
			if _, err := w.Write([]byte(fmt.Sprintf(" %s", key))); err != nil {
				return err
			}
		} else {
			if _, err := w.Write([]byte(fmt.Sprintf(" %s=\"%v\"", key, val))); err != nil {
				return err
			}
		}
	}

	if _, err := w.Write([]byte(">")); err != nil {
		return err
	}

	if e.void {
		return nil
	}

	for _, child := range e.children {
		if err := child.Render(w); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte("</" + e.tag + ">")); err != nil {
		return err
	}

	return nil
}

type TextElement struct {
	text     string
	children []Node
}

func UnsafeText(text string, children ...Node) *TextElement {
	return &TextElement{
		text:     text,
		children: children,
	}
}

func Text(text string, children ...Node) *TextElement {
	return &TextElement{
		text:     html.EscapeString(text),
		children: children,
	}
}

func (e *TextElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte(e.text)); err != nil {
		return err
	}

	for _, child := range e.children {
		if err := child.Render(w); err != nil {
			return err
		}
	}

	return nil
}
