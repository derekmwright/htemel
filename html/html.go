package html

import (
	"fmt"
	"github.com/derekmwright/htemel"
	"io"
)

type HtmlElement struct {
	children []htemel.Node
}

func Html(children ...htemel.Node) *HtmlElement {
	fmt.Println("Hello, World!")
	return nil
}
func (e *HtmlElement) Render(writer io.Writer) error {
	fmt.Println("Hello, World!")
	return nil
}
