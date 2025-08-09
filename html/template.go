package html

import (
  "github.com/derekmwright/htemel"
  "io"
)

type TemplateElement struct {
	children []htemel.Node
}

// Template creates a tag <template> instance and returns it for further modification.
// Any children passed will be nested within the tag.
func Template(children ...htemel.Node) *TemplateElement {
	node := &TemplateElement{
		children: children,
	}

	return node
}

func TemplateIf(condition bool, children ...htemel.Node) *TemplateElement {
	if condition {
		return Template(children...)
	}

	return nil
}

func (e *TemplateElement) Render(w io.Writer) error {
	if _, err := w.Write([]byte("<template")); err != nil {
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

	if _, err := w.Write([]byte("</template>")); err != nil {
		return err
	}

	return nil
}
