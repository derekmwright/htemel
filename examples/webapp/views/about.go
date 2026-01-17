package views

import (
	. "github.com/derekmwright/htemel"
	. "github.com/derekmwright/htemel/html"
)

func About() Node {
	return PageContent("About",
		P().Text("This is the about page."),
	)
}
