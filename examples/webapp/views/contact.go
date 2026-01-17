package views

import (
	. "github.com/go-htemel/htemel"
	. "github.com/go-htemel/htemel/html"
)

func Contact() Node {
	return PageContent("Contact",
		P().Text("This is the contact page."),
	)
}
