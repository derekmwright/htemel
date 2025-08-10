package spec

import (
	"errors"
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
		&AttributeTypeEnum{
			Name:        "autocorrect",
			Description: "The autocorrect attribute can be used on an editing host to control autocorrection behavior for the hosted editable region, on an input or textarea element to control the behavior when inserting text into that element, or on a form element to control the default behavior for all autocapitalize-and-autocorrect inheriting elements associated with the form element.",
			Allowed: map[string]struct{}{
				"on":  {},
				"off": {},
			},
		},
		&AttributeTypeBool{
			Name:        "autofocus",
			Description: "The autofocus content attribute allows the author to indicate that an element is to be focused as soon as the page is loaded, allowing the user to just start typing without having to manually focus the main element.",
		},
		&AttributeTypeSST{
			Name:        "class",
			Description: "When specified on HTML elements, the class attribute must have a value that is a set of space-separated tokens representing the various classes that the element belongs to.",
		},
		&AttributeTypeEnum{
			Name:        "contenteditable",
			Description: "",
			Allowed: map[string]struct{}{
				"true":           {},
				"false":          {},
				"plaintext-only": {},
			},
		},
		&AttributeTypeEnum{
			Name:        "dir",
			Description: "",
			Allowed: map[string]struct{}{
				"ltr":  {},
				"rtl":  {},
				"auto": {},
			},
		},
		&AttributeTypeEnum{
			Name:        "draggable",
			Description: "All HTML elements may have the draggable content attribute set.",
			Allowed: map[string]struct{}{
				"true":  {},
				"false": {},
			},
		},
		&AttributeTypeEnum{
			Name:        "enterkeyhint",
			Description: "The enterkeyhint content attribute is an enumerated attribute that specifies what action label (or icon) to present for the enter key on virtual keyboards. This allows authors to customize the presentation of the enter key in order to make it more helpful for users.",
			Allowed: map[string]struct{}{
				"enter":    {},
				"done":     {},
				"go":       {},
				"next":     {},
				"previous": {},
				"search":   {},
				"send":     {},
			},
		},
		&AttributeTypeEnum{
			Name:        "hidden",
			Description: "All HTML elements may have the hidden content attribute set.",
			Allowed: map[string]struct{}{
				"hidden":      {},
				"until-found": {},
			},
		},
		&AttributeTypeString{
			Name:        "id",
			Description: "The id attribute specifies its element's unique identifier (ID).",
		},
		&AttributeTypeString{
			Name:        "slot",
			Description: "The slot attribute is used to assign a slot to an element: an element with a slot attribute is assigned to the slot created by the slot element whose name attribute's value matches that slot attribute's value — but only if that slot element finds itself in the shadow tree whose root's host has the corresponding slot attribute value.",
		},
	}
}

func GenerateHTMLSpec(closer io.ReadCloser) (*Spec, error) {
	p := NewSpecParser(HTML)

	// Add the defined global attributes
	p.Spec.Attributes = GlobalAttributes()

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
