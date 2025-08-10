package htemel

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type Node interface {
	Render(w io.Writer) error
}

type GroupElement struct {
	children []Node
}

func Group(children ...Node) *GroupElement {
	return &GroupElement{children}
}

func (e *GroupElement) Render(w io.Writer) error {
	for _, child := range e.children {
		if err := child.Render(w); err != nil {
			return err
		}
	}

	return nil
}

type GenericElement struct {
	tag      string
	attrs    map[string]any
	void     bool
	children []Node
}

func Generic(tag string, attrs map[string]any, children ...Node) *GenericElement {
	return &GenericElement{
		tag:      tag,
		attrs:    attrs,
		void:     false,
		children: children,
	}
}

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
