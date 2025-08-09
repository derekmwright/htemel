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

	"github.com/derekmwright/htemel"
	. "github.com/derekmwright/htemel/html"
)

func main() {
	loggedIn := true
	htemel.Group(
		htemel.GenericVoid("!DOCTYPE", map[string]any{"html": nil}),
		Html(
			BodyIf(
				!loggedIn,
				Div(),
			),
			BodyIf(
				loggedIn,
				Nav(),
			),
		),
	).
		Render(os.Stdout)
}
```

Outputs:
```text
<!DOCTYPE html><html><body><nav></nav></body></html>
```

I have no plans on supporting indentation levels for nested elements, plus it's a waste of electrons for all that whitespace :P

## Why

I looked at `templ`, `gomponents`, and `gostar` but they just felt like there was too many rough edges.

With templ, it was still too much writing html and interpolating in templates.

Gomponents, definitely piqued my interest but seemed a little loose with what attributes were allowed on certain elements.
Given it was hand-crafted, trying to adjust global attributes vs element-specific attributes would be quite burdensome, so I get it.

Gostar, seemed a bit unmaintained but definitely felt like it could have potential, until I tried to figure out how to implement some custom TailwindPlus components with it.

So, I figured, I can write Go, lets give it a shot.

## Contributing

Feel free to fork and submit a PR if you want to help development.
My only ask is that you try to consider the feel and direction of the API before implementing new features.

## Generating

There's couple generators here...
I've included the HTML spec generator that pulls from the living HTML standard spec @ https://html.spec.whatwg.org/
I'd like the generator to be able to also generate a spec for SVG and MathML or other markup languages that you may feel would be beneficial.
Once a spec is generated, it can be used to generate package namespaced elements based on the provided spec.

You don't HAVE to generate a spec, if you have a custom web component, you can quickly write your own spec, and feed it to the element generator.
You can then use the generated elements for your own use-cases.
