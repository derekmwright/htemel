package views

import (
	. "github.com/derekmwright/htemel"
	. "github.com/derekmwright/htemel/html"
)

func Contact() Node {
	return PageContent("Contact",
		P().Text("This is the contact page."),
	)
}
