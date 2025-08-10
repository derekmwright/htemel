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
		&AttributeTypeBool{
			Name:        "inert",
			Description: "The inert attribute is a boolean attribute that indicates, by its presence, that the element and all its flat tree descendants which don't otherwise escape inertness (such as modal dialogs) are to be made inert by the user agent.",
		},
		&AttributeTypeEnum{
			Name:        "inputmode",
			Description: "User agents can support the inputmode attribute on form controls (such as the value of textarea elements), or in elements in an editing host (e.g., using contenteditable).",
			Allowed: map[string]struct{}{
				"none":    {},
				"text":    {},
				"tel":     {},
				"url":     {},
				"email":   {},
				"numeric": {},
				"decimal": {},
				"search":  {},
			},
		},
		&AttributeTypeString{
			Name:        "itemid",
			Description: "The itemid attribute, if specified, must have a value that is a valid URL potentially surrounded by spaces.",
		},
		&AttributeTypeSST{
			Name:        "itemprop",
			Description: "The itemprop attribute, if specified, must have a value that is an unordered set of unique space-separated tokens none of which are identical to another token, representing the names of the name-value pairs that it adds. The attribute's value must have at least one token.",
		},
		&AttributeTypeSST{
			Name:        "itemref",
			Description: "The itemref attribute, if specified, must have a value that is an unordered set of unique space-separated tokens none of which are identical to another token and consisting of IDs of elements in the same tree.",
		},
		&AttributeTypeBool{
			Name:        "itemscope",
			Description: "Every HTML element may have an itemscope attribute specified. The itemscope attribute is a boolean attribute.",
		},
		&AttributeTypeSST{
			Name:        "itemtype",
			Description: "The itemtype attribute, if specified, must have a value that is an unordered set of unique space-separated tokens, none of which are identical to another token and each of which is a valid URL string that is an absolute URL, and all of which are defined to use the same vocabulary. The attribute's value must have at least one token.",
		},
		&AttributeTypeString{
			Name:        "lang",
			Description: "The lang attribute (in no namespace) specifies the primary language for the element's contents and for any of the element's attributes that contain text. Its value must be a valid BCP 47 language tag, or the empty string. Setting the attribute to the empty string indicates that the primary language is unknown.",
		},
		&AttributeTypeString{
			Name:        "nonce",
			Description: "A nonce content attribute represents a cryptographic nonce (\"number used once\") which can be used by Content Security Policy to determine whether or not a given fetch will be allowed to proceed. The value is text.",
		},
		&AttributeTypeString{
			Name:        "popover",
			Description: "All HTML elements may have the popover content attribute set. When specified, the element won't be rendered until it becomes shown, at which point it will be rendered on top of other page content.",
		},
		&AttributeTypeString{
			Name:        "slot",
			Description: "The slot attribute is used to assign a slot to an element: an element with a slot attribute is assigned to the slot created by the slot element whose name attribute's value matches that slot attribute's value — but only if that slot element finds itself in the shadow tree whose root's host has the corresponding slot attribute value.",
		},
		&AttributeTypeEnum{
			Name:        "spellcheck",
			Description: "User agents can support the checking of spelling and grammar of editable text, either in form controls (such as the value of textarea elements), or in elements in an editing host (e.g. using contenteditable).",
			Allowed: map[string]struct{}{
				"true":  {},
				"false": {},
			},
		},
		&AttributeTypeString{
			Name:        "style",
			Description: "All HTML elements may have the style content attribute set. This is a style attribute as defined by CSS Style Attributes.",
		},
		&AttributeTypeNumber{
			Name:        "tabindex",
			Description: "The tabindex attribute, if specified, must have a value that is a valid integer. Positive numbers specify the relative position of the element's focusable areas in the sequential focus navigation order, and negative numbers indicate that the control is not sequentially focusable.",
		},
		&AttributeTypeString{
			Name:        "title",
			Description: "The title attribute represents advisory information for the element, such as would be appropriate for a tooltip. On a link, this could be the title or a description of the target resource; on an image, it could be the image credit or a description of the image; on a paragraph, it could be a footnote or commentary on the text; on a citation, it could be further information about the source; on interactive content, it could be a label for, or instructions for, use of the element; and so forth. The value is text.",
		},
		&AttributeTypeEnum{
			Name:        "translate",
			Description: "The translate attribute is used to specify whether an element's attribute values and the values of its Text node children are to be translated when the page is localized, or whether to leave them unchanged.",
			Allowed: map[string]struct{}{
				"yes": {},
				"no":  {},
			},
		},
		&AttributeTypeEnum{
			Name:        "writingsuggestions",
			Description: "User agents offer writing suggestions as users type into editable regions, either in form controls (e.g., the textarea element) or in elements in an editing host.",
			Allowed: map[string]struct{}{
				"true":  {},
				"false": {},
			},
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
