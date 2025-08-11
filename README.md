⚠️ This project is a **work in progress**, check back often and throw a ⭐ on it if you are interested!

# htemel

Pronounced HTML, htemel is a library focused on generating documents for the web using pure Go functions and not a templating language.
The library is being written in a way that makes it easy to extend and add additional functionality, or provide support for any custom web components you may have.

## Installation

```shell
go get github.com/derekmwright/htemel
go get github.com/derekmwright/htemel/html
```

## Examples

Using a `.` (dot) import on a namespaced element package helps readability, but it's not required.

### Basic Usage

```go
package main

import (
	"os"

	. "github.com/derekmwright/htemel"
	. "github.com/derekmwright/htemel/html"
)

func MainLayout(children ...Node) Node {
	return Group(
		GenericVoid("!DOCTYPE", map[string]any{"html": nil}),
		Html(
			Head(
				Meta().Charset("UTF-8"),
				Meta().Content("width=device-width, initial-scale=1.0"),
				Title(Text("Example htemel Page")),
				Link().Href("site.css").Rel("stylesheet"),
			),
			Body(
				children...,
			).Id("app-content"),
		).Class("h-dvh bg-gray-200").Lang("en"),
	)
}

func Navigation(menuItems ...Node) Node {
	c := Nav(
		Ul(
			Group(menuItems...),
		).Class("flex list-none"),
	).Id("main-navigation")

	return c
}

func main() {
	loggedIn := true
	MainLayout(
		DivTernary(
			loggedIn,
			Group(
				Navigation(),
				Div(
					P(Text("Welcome back!")),
				),
			),
			Div(
				P(Text("Please login.")),
			),
		),
	).Render(os.Stdout)
}
```

Outputs (I formatted the result to help readability of the output):
```html
<!DOCTYPE html>
<html class="h-dvh bg-gray-200" lang="en">
<head>
    <meta charset="UTF-8">
    <meta content="width=device-width, initial-scale=1.0">
    <title>Example htemel Page</title>
    <link href="site.css" rel="stylesheet">
</head>
<body id="app-content">
<div>
    <nav id="main-navigation">
        <ul class="flex list-none"></ul>
    </nav>
    <div>
        <p>Welcome back!</p>
    </div>
</div>
</body>
</html>
```

## The Goal

The aim of this package is to provide a robust and **compile time safe** way to generate HTML.
HTML elements are represented by functions that return a related struct that can then be chain called for further modification.
An element has a limited set of attributes available to it and are unique to an instance of it.
This means, when using code completion or hinting in your editor, you will only see attributes allowed for the specific element.
If an attribute is an enumerated-type, special types are generated to conform to the HTML spec for you.

Example:
```go
html.Link().HttpEquiv(html.MetaHttpEquivEnumContentType)
```

Additionally, due to the way function chain calling works, and HTML attributes are mutually exclusive on an element, the most recent chain call takes precedence.

Example:
```go
// You can also pass a list of strings to space-seperated-token type attributes
// ie: Class("active", "hover:bg-indigo-400")
html.A().Class("active hover:bg-indigo-400").Class("this-takes-precedence")
```

Resuls:
```text
<a class="this-takes-precedence">
```

## Why

I looked at `templ`, `gomponents`, and `gostar` but they just felt like none of them fully met my needs.

With templ, it was still too much writing html and interpolating in templates.

Gomponents, definitely piqued my interest but seemed a little loose with what attributes were allowed on certain elements.
Given it was hand-crafted, trying to adjust global attributes vs element-specific attributes would be quite burdensome, so I get it.

Gostar, seemed a bit unmaintained but definitely felt like it could have potential, but I had some issues trying to figure out how to implement some custom TailwindPlus components with it.

So, I figured, I can write Go, lets give it a shot.

## Contributing

Feel free to fork and submit a PR if you want to help development.
My only ask is that you try to consider the feel and direction of the API before implementing new features.

## Generating

There's couple generators here...
I've included the HTML spec generator that pulls from the living HTML standard spec @ https://html.spec.whatwg.org/ - caveat: The parser is a WIP and there are a bunch of hand-crafted attributes to fill in gaps.
I'd like the generator to be able to also generate a spec for SVG and MathML or other markup languages that you may feel would be beneficial.
Once a spec is generated, it can be used to generate package namespaced elements based on the provided spec.

You don't HAVE to generate a spec, if you have a custom web component, you can quickly write your own spec, and feed it to the element generator.
Then use the generated elements for your own use-cases.
