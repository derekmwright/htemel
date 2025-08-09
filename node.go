package htemel

import (
	"fmt"
	"io"
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
	children []Node
}

func Generic(tag string, attrs map[string]any, children ...Node) *GenericElement {
	return &GenericElement{
		tag:      tag,
		attrs:    attrs,
		children: children,
	}
}

func (e *GenericElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<" + e.tag)); err != nil {
		return err
	}

	for key, val := range e.attrs {
		if _, err := w.Write([]byte(fmt.Sprintf(" %s=\"%v\"", key, val))); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte(">")); err != nil {
		return err
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
