package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type ArticleElement struct {
	children []htemel.Node
}

// Article creates a tag <article> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Article(children ...htemel.Node) *ArticleElement {
	node := &ArticleElement{
		children: children,
	}

	return node
}

func ArticleIf(condition bool, children ...htemel.Node) *ArticleElement {
	if condition {
		return Article(children...)
	}

	return nil
}

func (e *ArticleElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<article")); err != nil {
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

	if _, err := w.Write([]byte("</article>")); err != nil {
		return err
	}

	return nil
}
