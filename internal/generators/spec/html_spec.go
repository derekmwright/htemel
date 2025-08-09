package spec

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// GlobalAttributes are getting hardcoded right now as parsing them is a bit annoying.
func GlobalAttributes() []Attribute {
	return []Attribute{
		&AttributeTypeChar{
			Name:        "accesskey",
			Description: "The accesskey attribute's value is used by the user agent as a guide for creating a keyboard shortcut that activates or focuses the element.",
		},
		&AttributeTypeEnum{
			Name:        "autocapitalize",
			Description: "The autocapitalize attribute is an enumerated attribute whose states are the possible autocapitalization hints. The autocapitalization hint specified by the attribute's state combines with other considerations to form the used autocapitalization hint, which informs the behavior of the user agent.",
			Allowed: map[string]struct{}{
				"off":        {},
				"none":       {},
				"on":         {},
				"sentences":  {},
				"words":      {},
				"characters": {},
			},
		},
	}
}

func GenerateHTMLSpec(closer io.ReadCloser) (*Spec, error) {
	p := NewSpecParser(HTML)

	defer func(closer io.ReadCloser) {
		err := closer.Close()
		if err != nil {
			panic(err)
		}
	}(closer)

	doc, err := html.Parse(closer)
	if err != nil {
		panic(err)
	}

	var body *html.Node
	var ok bool
	if body, ok = findTag(doc, "body"); !ok {
		return nil, errors.New("could not find body")
	}

	start := false
	end := false
	for child := range body.ChildNodes() {
		if child.Data == "h2" {
			if start {
				end = true
				start = false
			}

			if _, ok := getIDIndex(child.Attr, "id", "semantics"); ok {
				start = true
			}
		}

		if end {
			break
		}

		if start {
			// Look for H4 elements and then check to see if their ID contains the term "element".
			// If so, then check the `code` tag for the text value.
			if child.Data == "h4" {
				var id string
				if id, ok = getAttribute(child.Attr, "id"); ok {
					if strings.Contains(id, "the-") && strings.Contains(id, "-element") {
						fmt.Println(id)
						var tagNode *html.Node
						if tagNode, ok = findTag(child, "code"); ok {
							p.Activate(tagNode.FirstChild.Data)
						}
					}
				}
			}

			if child.Data == "p" {
				if p.active && !p.descParsed {
					p.currElement.Description = gatherText(child, nil)
					p.descParsed = true
					p.Reset()
				}
			}
		}
	}

	return p.Spec, nil
}
