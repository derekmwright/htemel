package spec

import (
	"errors"
	"io"
	"slices"
	"strconv"
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
			AllowEmpty:  true,
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
			AllowEmpty:  true,
			Allowed: map[string]struct{}{
				"true":           {},
				"false":          {},
				"plaintext-only": {},
			},
		},
		&AttributeTypePrefixedCustom{
			Name:        "data",
			Description: "A custom data attribute is an attribute in no namespace whose name starts with the string \"data-\", has at least one character after the hyphen, is a valid attribute local name, and contains no ASCII upper alphas.",
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
			AllowEmpty:  true,
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
			AllowEmpty:  true,
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
			AllowEmpty:  true,
			Allowed: map[string]struct{}{
				"yes": {},
				"no":  {},
			},
		},
		&AttributeTypeEnum{
			Name:        "writingsuggestions",
			Description: "User agents offer writing suggestions as users type into editable regions, either in form controls (e.g., the textarea element) or in elements in an editing host.",
			AllowEmpty:  true,
			Allowed: map[string]struct{}{
				"true":  {},
				"false": {},
			},
		},
	}
}

// More statically defined things because parsing this from the whatwg spec is painful.
// Ordering them as per the way the reference site orders them for quick reference/upkeep.

var fetchPriority = &AttributeTypeEnum{
	Name:        "fetchpriority",
	Description: "Sets the priority for fetches initiated by the element",
	Allowed: map[string]struct{}{
		"high": {},
		"low":  {},
		"auto": {},
	},
}
var referrerPolicy = &AttributeTypeString{
	Name:        "referrerpolicy",
	Description: "Referrer policy for fetches initiated by the element",
}
var width = &AttributeTypeNumber{
	Name:        "width",
	Description: "Horizontal dimension",
}
var height = &AttributeTypeNumber{
	Name:        "height",
	Description: "Vertical dimension",
}
var loading = &AttributeTypeEnum{
	Name:        "loading",
	Description: "Used when determining loading deferral",
	Allowed: map[string]struct{}{
		"lazy":  {},
		"eager": {},
	},
}
var src = &AttributeTypeString{
	Name:        "src",
	Description: "Address of the resource",
}
var typ = &AttributeTypeString{
	Name:        "type",
	Description: "Type of embedded resource",
}
var crossorigin = &AttributeTypeEnum{
	Name:        "crossorigin",
	Description: "How the element handles crossorigin requests",
	AllowEmpty:  true,
	Allowed: map[string]struct{}{
		"anonymous":       {},
		"use-credentials": {},
	},
}
var href = &AttributeTypeString{
	Name:        "href",
	Description: "Address of the hyperlink",
}
var colspan = &AttributeTypeNumber{
	Name:        "span",
	Description: "Number of columns spanned by the element where the number is > 0 && <= 1000",
}
var rowspan = &AttributeTypeNumber{
	Name:        "rowspan",
	Description: "Number of rows that the cell is to span where the number is > 0 && <= 65534",
}
var headers = &AttributeTypeSST{
	Name:        "headers",
	Description: "The header cells for this cell",
}

func htmlAttr() []Attribute {
	return make([]Attribute, 0)
}

func headAttr() []Attribute {
	return make([]Attribute, 0)
}

func titleAttr() []Attribute {
	return make([]Attribute, 0)
}

func baseAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "href",
			Description: "Document base URL",
		},
		&AttributeTypeString{
			Name:        "target",
			Description: "Default navigable for hyperlink navigation and form submission",
		},
	}
}

func linkAttr() []Attribute {
	return []Attribute{
		href,
		crossorigin,
		&AttributeTypeSST{
			Name:        "rel",
			Description: "Relationship between the document containing the hyperlink and the destination resource",
		},
		&AttributeTypeString{
			Name:        "media",
			Description: "Applicable media",
		},
		&AttributeTypeString{
			Name:        "integrity",
			Description: "Integrity metadata used in Subresource Integrity checks [SRI]",
		},
		&AttributeTypeString{
			Name:        "hreflang",
			Description: "Language of the linked resource",
		},
		&AttributeTypeString{
			Name:        "type",
			Description: "Hint for the type of the referenced resource",
		},
		referrerPolicy,
		&AttributeTypeSST{
			Name:        "sizes",
			Description: "Sizes of the icons (for rel=\"icon\")",
		},
		&AttributeTypeString{
			Name:        "imagesrcset",
			Description: "Images to use in different situations, e.g., high-resolution displays, small monitors, etc. (for rel=\"preload\")",
		},
		&AttributeTypeString{
			Name:        "imagesizes",
			Description: "Image sizes for different page layouts (for rel=\"preload\")",
		},
		&AttributeTypeString{
			Name:        "as",
			Description: "Potential destination for a preload request (for rel=\"preload\" and rel=\"modulepreload\")",
		},
		&AttributeTypeEnum{
			Name:        "blocking",
			Description: "Whether the element is potentially render-blocking",
			Allowed: map[string]struct{}{
				"render": {},
			},
		},
		&AttributeTypeString{
			Name:        "color",
			Description: "Color to use when customizing a site's icon (for rel=\"mask-icon\")",
		},
		&AttributeTypeBool{
			Name:        "disabled",
			Description: "Whether the link is disabled",
		},
		fetchPriority,
	}
}

func metaAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "name",
			Description: "Metadata name",
		},
		&AttributeTypeEnum{
			Name:        "http-equiv",
			Description: "Pragma directive",
			Allowed: map[string]struct{}{
				"content-language":        {},
				"content-type":            {},
				"default-style":           {},
				"refresh":                 {},
				"set-cookie":              {},
				"x-ua-compatible":         {},
				"content-security-policy": {},
			},
		},
		&AttributeTypeString{
			Name:        "content",
			Description: "Value of the element",
		},
		&AttributeTypeString{
			Name:        "charset",
			Description: "Character encoding declaration",
		},
		&AttributeTypeString{
			Name:        "media",
			Description: "Applicable media",
		},
	}
}

func styleAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "media",
			Description: "Applicable media",
		},
		&AttributeTypeString{
			Name:        "blocking",
			Description: "Whether the element is potentially render-blocking",
		},
	}
}

func bodyAttr() []Attribute {
	attrs := []string{
		"onafterprint",
		"onbeforeprint",
		"onbeforeunload",
		"onhashchange",
		"onlanguagechange",
		"onmessage",
		"onmessageerror",
		"onoffline",
		"ononline",
		"onpageswap",
		"onpagehide",
		"onpagereveal",
		"onpageshow",
		"onpopstate",
		"onrejectionhandled",
		"onstorage",
		"onunhandledrejection",
		"onunload",
	}

	var out []Attribute
	for _, attr := range attrs {
		out = append(out, &AttributeTypeString{
			Name: attr,
		})
	}

	return out
}

func blockQuoteAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "cite",
			Description: "Content inside a blockquote must be quoted from another source, whose address, if it has one, may be cited in the cite attribute.",
		},
	}
}

func olAttr() []Attribute {
	return []Attribute{
		&AttributeTypeBool{
			Name:        "reversed",
			Description: "Number the list backwards",
		},
		&AttributeTypeNumber{
			Name:        "start",
			Description: "The start attribute, if present, must be a valid integer. It is used to determine the starting value of the list.",
		},
		&AttributeTypeChar{
			Name:        "type",
			Description: "The type attribute can be used to specify the kind of marker to use in the list, in the cases where that matters (e.g. because items are to be referenced by their number/letter).",
		},
	}
}

func liAttr() []Attribute {
	return []Attribute{
		&AttributeTypeNumber{
			Name:        "value",
			Description: "If the element is not a child of an ul or menu element: value — Ordinal value of the list item",
		},
	}
}

func aAttr() []Attribute {
	return []Attribute{
		href,
		&AttributeTypeString{
			Name:        "target",
			Description: "Navigable for hyperlink navigation",
		},
		&AttributeTypeBool{
			Name:        "download",
			Description: "Whether to download the resource instead of navigating to it, and its filename if so",
		},
		&AttributeTypeSST{
			Name:        "ping",
			Description: "URLs to ping",
		},
		&AttributeTypeSST{
			Name:        "rel",
			Description: "Relationship between the location in the document containing the hyperlink and the destination resource",
		},
		&AttributeTypeString{
			Name:        "hreflang",
			Description: "Language of the linked resource",
		},
		&AttributeTypeString{
			Name:        "type",
			Description: "Hint for the type of the referenced resource",
		},
		referrerPolicy,
	}
}

func qAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "cite",
			Description: "Link to the source of the quotation or more information about the edit",
		},
	}
}

func dataAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "value",
			Description: "Machine-readable value",
		},
	}
}

func timeAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "datetime",
			Description: "Machine-readable value",
		},
	}
}

func insDelAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "cite",
			Description: "Link to the source of the quotation or more information about the edit",
		},
		&AttributeTypeString{
			Name:        "datetime",
			Description: "Date and (optionally) time of the change",
		},
	}
}

func sourceAttr() []Attribute {
	return []Attribute{
		typ,
		&AttributeTypeString{
			Name:        "media",
			Description: "Applicable media",
		},
		src,
		&AttributeTypeString{
			Name:        "srcset",
			Description: "Images to use in different situations, e.g., high-resolution displays, small monitors, etc.",
		},
		&AttributeTypeString{
			Name:        "sizes",
			Description: "Image sizes for different page layouts",
		},
		width,
		height,
	}
}

func imgAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "alt",
			Description: "Replacement text for use when images are not available",
		},
		src,
		&AttributeTypeString{
			Name:        "srcset",
			Description: "Images to use in different situations, e.g., high-resolution displays, small monitors, etc.",
		},
		&AttributeTypeString{
			Name:        "sizes",
			Description: "Image sizes for different page layouts",
		},
		crossorigin,
		&AttributeTypeString{
			Name:        "usemap",
			Description: "Name of image map to use",
		},
		&AttributeTypeBool{
			Name:        "ismap",
			Description: "Whether the image is a server-side image map",
		},
		width,
		height,
		referrerPolicy,
		&AttributeTypeEnum{
			Name:        "decoding",
			Description: "Decoding hint to use when processing this image for presentation",
			Allowed: map[string]struct{}{
				"sync":  {},
				"async": {},
				"auto":  {},
			},
		},
		loading,
		fetchPriority,
	}
}

func iframeAttr() []Attribute {
	return []Attribute{
		src,
		&AttributeTypeString{
			Name:        "srcdoc",
			Description: "A document to render in the iframe",
		},
		&AttributeTypeString{
			Name:        "name",
			Description: "Name of content navigable",
		},
		&AttributeTypeSST{
			Name:        "sandbox",
			Description: "Security rules for nested content",
		},
		&AttributeTypeString{
			Name:        "allow",
			Description: "Permissions policy to be applied to the iframe's contents",
		},
		&AttributeTypeBool{
			Name:        "allowfullscreen",
			Description: "Whether to allow the iframe's contents to use requestFullscreen()",
		},
		width,
		height,
		referrerPolicy,
		loading,
	}
}

func embedAttr() []Attribute {
	return []Attribute{
		src,
		typ,
		width,
		height,
	}
}

func objectAttr() []Attribute {
	return []Attribute{
		// TODO: Fix this collision with `data-`
		&AttributeTypeString{
			Name:        "data",
			Description: "Address of the resource",
		},
		typ,
		&AttributeTypeString{
			Name:        "name",
			Description: "Name of content navigable",
		},
		&AttributeTypeString{
			Name:        "form",
			Description: "Associates the element with a form element",
		},
		width,
		height,
	}
}

func videoAttr() []Attribute {
	return []Attribute{
		src,
		crossorigin,
		&AttributeTypeString{
			Name:        "poster",
			Description: "Poster frame to show prior to video playback",
		},
		&AttributeTypeEnum{
			Name:        "preload",
			Description: "Hints how much buffering the media resource will likely need",
			AllowEmpty:  true,
			Allowed: map[string]struct{}{
				"auto":     {},
				"none":     {},
				"metadata": {},
			},
		},
		&AttributeTypeBool{
			Name:        "autoplay",
			Description: "Hint that the media resource can be started automatically when the page is loaded",
		},
		&AttributeTypeBool{
			Name:        "playsinline",
			Description: "Encourage the user agent to display video content within the element's playback area",
		},
		&AttributeTypeBool{
			Name:        "loop",
			Description: "Whether to loop the media resource",
		},
		&AttributeTypeBool{
			Name:        "muted",
			Description: "Whether to mute the media resource by default",
		},
		&AttributeTypeBool{
			Name:        "controls",
			Description: "Show user agent controls",
		},
		width,
		height,
	}
}

func audioAttr() []Attribute {
	return []Attribute{
		src,
		crossorigin,
		&AttributeTypeEnum{
			Name:        "preload",
			Description: "Hints how much buffering the media resource will likely need",
			AllowEmpty:  true,
			Allowed: map[string]struct{}{
				"auto":     {},
				"none":     {},
				"metadata": {},
			},
		},
		&AttributeTypeBool{
			Name:        "autoplay",
			Description: "Hint that the media resource can be started automatically when the page is loaded",
		},
		&AttributeTypeBool{
			Name:        "playsinline",
			Description: "Encourage the user agent to display video content within the element's playback area",
		},
		&AttributeTypeBool{
			Name:        "loop",
			Description: "Whether to loop the media resource",
		},
		&AttributeTypeBool{
			Name:        "muted",
			Description: "Whether to mute the media resource by default",
		},
		&AttributeTypeBool{
			Name:        "controls",
			Description: "Show user agent controls",
		},
	}
}

func trackAttr() []Attribute {
	return []Attribute{
		&AttributeTypeEnum{
			Name:        "kind",
			Description: "The type of text track",
			Allowed: map[string]struct{}{
				"subtitles":    {},
				"captions":     {},
				"descriptions": {},
				"chapters":     {},
				"metadata":     {},
			},
		},
		src,
		&AttributeTypeString{
			Name:        "srclang",
			Description: "Language of the text track",
		},
		&AttributeTypeString{
			Name:        "label",
			Description: "User-visible label",
		},
		&AttributeTypeBool{
			Name:        "default",
			Description: "Enable the track if no other text track is more suitable",
		},
	}
}

func mapAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "name",
			Description: "Name of image map to reference from the usemap attribute",
		},
	}
}

func areaAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "alt",
			Description: "Replacement text for use when images are not available",
		},
		&AttributeTypeString{
			Name:        "coords",
			Description: "Coordinates for the shape to be created in an image map",
		},
		&AttributeTypeEnum{
			Name:        "shape",
			Description: "The kind of shape to be created in an image map",
			Allowed: map[string]struct{}{
				"circle":  {},
				"default": {},
				"poly":    {},
				"rect":    {},
			},
		},
		href,
		&AttributeTypeString{
			Name:        "target",
			Description: "Navigable for hyperlink navigation",
		},
		&AttributeTypeBool{
			Name:        "download",
			Description: "Whether to download the resource instead of navigating to it, and its filename if so",
		},
		&AttributeTypeSST{
			Name:        "ping",
			Description: "URLs to ping",
		},
		&AttributeTypeSST{
			Name:        "rel",
			Description: "Relationship between the location in the document containing the hyperlink and the destination resource",
		},
		referrerPolicy,
	}
}

func colgroupAttr() []Attribute {
	return []Attribute{
		colspan,
	}
}

func colAttr() []Attribute {
	return []Attribute{
		colspan,
	}
}

func tdAttr() []Attribute {
	return []Attribute{
		colspan,
		rowspan,
		headers,
	}
}

func thAttr() []Attribute {
	return []Attribute{
		colspan,
		rowspan,
		headers,
		&AttributeTypeString{
			Name:        "scope",
			Description: "Specifies which cells the header cell applies to",
		},
		&AttributeTypeString{
			Name:        "abbr",
			Description: "Alternative label to use for the header cell when referencing the cell in other contexts",
		},
	}
}

func formAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "accept-charset",
			Description: "Character encodings to use for form submission",
		},
		&AttributeTypeString{
			Name:        "action",
			Description: "URL to use for form submission",
		},
		&AttributeTypeEnum{
			Name:        "autocomplete",
			Description: "Default setting for autofill feature for controls in the form",
			Allowed: map[string]struct{}{
				"on":  {},
				"off": {},
			},
		},
		&AttributeTypeString{
			Name:        "enctype",
			Description: "Entry list encoding type to use for form submission",
		},
		&AttributeTypeEnum{
			Name:        "method",
			Description: "Variant to use for form submission",
			Allowed: map[string]struct{}{
				"get":    {},
				"post":   {},
				"dialog": {},
			},
		},
		&AttributeTypeString{
			Name:        "name",
			Description: "Name of form to use in the document.forms API",
		},
		&AttributeTypeBool{
			Name:        "novalidate",
			Description: "Bypass form control validation for form submission",
		},
		&AttributeTypeString{
			Name:        "target",
			Description: "Navigable for form submission",
		},
		&AttributeTypeSST{
			Name:        "rel",
			Description: "The rel attribute on form elements controls what kinds of links the elements create.",
		},
	}
}

func labelAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "for",
			Description: "Associate the label with form control",
		},
	}
}

func inputAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "accept",
			Description: "Hint for expected file type in file upload controls",
		},
		&AttributeTypeBool{
			Name:        "alpha",
			Description: "Allow the color's alpha component to be set",
		},
		&AttributeTypeString{
			Name:        "alt",
			Description: "Replacement text for use when images are not available",
		},
		&AttributeTypeSST{
			Name:        "autocomplete",
			Description: "Hint for form autofill feature",
		},
		&AttributeTypeBool{
			Name:        "checked",
			Description: "Hint for form autofill feature",
		},
		&AttributeTypeEnum{
			Name:        "colorspace",
			Description: "The color space of the serialized color",
			Allowed: map[string]struct{}{
				"limited-srgb": {},
				"display-p3":   {},
			},
		},
		&AttributeTypeString{
			Name:        "dirname",
			Description: "Name of form control to use for sending the element's directionality in form submission",
		},
		&AttributeTypeBool{
			Name:        "disabled",
			Description: "Whether the form control is disabled",
		},
		&AttributeTypeString{
			Name:        "form",
			Description: "Associates the element with a form element",
		},
		&AttributeTypeString{
			Name:        "formaction",
			Description: "URL to use for form submission",
		},
		&AttributeTypeString{
			Name:        "formenctype",
			Description: "Entry list encoding type to use for form submission",
		},
		&AttributeTypeEnum{
			Name:        "formmethod",
			Description: "Variant to use for form submission",
			Allowed: map[string]struct{}{
				"get":    {},
				"post":   {},
				"dialog": {},
			},
		},
		&AttributeTypeBool{
			Name:        "formnovalidate",
			Description: "Bypass form control validation for form submission",
		},
		&AttributeTypeString{
			Name:        "formtarget",
			Description: "Navigable for form submission",
		},
		height,
		&AttributeTypeString{
			Name:        "list",
			Description: "List of autocomplete options",
		},
		&AttributeTypeNumber{
			Name:        "max",
			Description: "Maximum value",
		},
		&AttributeTypeNumber{
			Name:        "maxlength",
			Description: "Maximum length of value",
		},
		&AttributeTypeNumber{
			Name:        "min",
			Description: "Minimum value",
		},
		&AttributeTypeNumber{
			Name:        "minlength",
			Description: "Minimum length of value",
		},
		&AttributeTypeBool{
			Name:        "multiple",
			Description: "Whether to allow multiple values",
		},
		&AttributeTypeString{
			Name:        "name",
			Description: "Name of the element to use for form submission and in the form.elements API",
		},
		&AttributeTypeString{
			Name:        "pattern",
			Description: "Pattern to be matched by the form control's value",
		},
		&AttributeTypeString{
			Name:        "placeholder",
			Description: "User-visible label to be placed within the form control",
		},
		&AttributeTypeString{
			Name:        "popovertarget",
			Description: "Targets a popover element to toggle, show, or hide",
		},
		&AttributeTypeEnum{
			Name:        "popovertargetaction",
			Description: "Indicates whether a targeted popover element is to be toggled, shown, or hidden",
			Allowed: map[string]struct{}{
				"toggle": {},
				"show":   {},
				"hide":   {},
			},
		},
		&AttributeTypeBool{
			Name:        "readonly",
			Description: "Whether to allow the value to be edited by the user",
		},
		&AttributeTypeBool{
			Name:        "required",
			Description: "Whether the control is required for form submission",
		},
		&AttributeTypeNumber{
			Name:        "size",
			Description: "Size of the control",
		},
		src,
		&AttributeTypeString{
			Name:        "step",
			Description: "Granularity to be matched by the form control's value",
		},
		&AttributeTypeEnum{
			Name:        "type",
			Description: "Type of form control",
			Allowed: map[string]struct{}{
				"hidden":         {},
				"text":           {},
				"search":         {},
				"tel":            {},
				"url":            {},
				"email":          {},
				"password":       {},
				"date":           {},
				"month":          {},
				"week":           {},
				"time":           {},
				"datetime-local": {},
				"number":         {},
				"range":          {},
				"color":          {},
				"checkbox":       {},
				"radio":          {},
				"file":           {},
				"submit":         {},
				"image":          {},
				"reset":          {},
				"button":         {},
			},
		},
		&AttributeTypeString{
			Name:        "value",
			Description: "Value of the form control",
		},
		width,
	}
}

func buttonAttr() []Attribute {
	return []Attribute{
		&AttributeTypeEnum{
			Name:        "command",
			Description: "Indicates to the targeted element which action to take.",
			AllowCustom: true,
			Allowed: map[string]struct{}{
				"toggle-popover": {},
				"show-popover":   {},
				"hide-popover":   {},
				"close":          {},
				"request-close":  {},
				"show-modal":     {},
			},
		},
		&AttributeTypeString{
			Name:        "commandfor",
			Description: "Targets another element to be invoked.",
		},
		&AttributeTypeBool{
			Name:        "disabled",
			Description: "Whether the form control is disabled",
		},
		&AttributeTypeString{
			Name:        "form",
			Description: "Associates the element with a form element",
		},
		&AttributeTypeString{
			Name:        "formaction",
			Description: "URL to use for form submission",
		},
		&AttributeTypeString{
			Name:        "formenctype",
			Description: "Entry list encoding type to use for form submission",
		},
		&AttributeTypeEnum{
			Name:        "formmethod",
			Description: "Variant to use for form submission",
			Allowed: map[string]struct{}{
				"get":    {},
				"post":   {},
				"dialog": {},
			},
		},
		&AttributeTypeBool{
			Name:        "formnovalidate",
			Description: "Bypass form control validation for form submission",
		},
		&AttributeTypeString{
			Name:        "formtarget",
			Description: "Navigable for form submission",
		},
		&AttributeTypeString{
			Name:        "name",
			Description: "Name of the element to use for form submission and in the form.elements API",
		},
		&AttributeTypeString{
			Name:        "popovertarget",
			Description: "Targets a popover element to toggle, show, or hide",
		},
		&AttributeTypeEnum{
			Name:        "popovertargetaction",
			Description: "Indicates whether a targeted popover element is to be toggled, shown, or hidden",
			Allowed: map[string]struct{}{
				"toggle": {},
				"show":   {},
				"hide":   {},
			},
		},
		&AttributeTypeEnum{
			Name:        "type",
			Description: "Type of button",
			Allowed: map[string]struct{}{
				"submit": {},
				"reset":  {},
				"button": {},
			},
		},
		&AttributeTypeString{
			Name:        "value",
			Description: "Value to be used for form submission",
		},
	}
}

func selectAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "autocomplete",
			Description: "Hint for form autofill feature",
		},
		&AttributeTypeBool{
			Name:        "disabled",
			Description: "Whether the form control is disabled",
		},
		&AttributeTypeString{
			Name:        "form",
			Description: "Associates the element with a form element",
		},
		&AttributeTypeBool{
			Name:        "multiple",
			Description: "Whether to allow multiple values",
		},
		&AttributeTypeString{
			Name:        "name",
			Description: "Name of the element to use for form submission and in the form.elements API",
		},
		&AttributeTypeBool{
			Name:        "required",
			Description: "Whether the control is required for form submission",
		},
		&AttributeTypeNumber{
			Name:        "size",
			Description: "Size of the control",
		},
	}
}

func optgroupAttr() []Attribute {
	return []Attribute{
		&AttributeTypeBool{
			Name:        "disabled",
			Description: "Whether the form control is disabled",
		},
		&AttributeTypeString{
			Name:        "label",
			Description: "User-visible label",
		},
	}
}

func optionAttr() []Attribute {
	return []Attribute{
		&AttributeTypeBool{
			Name:        "disabled",
			Description: "Whether the form control is disabled",
		},
		&AttributeTypeString{
			Name:        "label",
			Description: "User-visible label",
		},
		&AttributeTypeBool{
			Name:        "selected",
			Description: "Whether the option is selected by default",
		},
		&AttributeTypeString{
			Name:        "value",
			Description: "Value to be used for form submission",
		},
	}
}

func textareaAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "autocomplete",
			Description: "Hint for form autofill feature",
		},
		&AttributeTypeNumber{
			Name:        "cols",
			Description: "Maximum number of characters per line",
		},
		&AttributeTypeString{
			Name:        "dirname",
			Description: "Name of form control to use for sending the element's directionality in form submission",
		},
		&AttributeTypeBool{
			Name:        "disabled",
			Description: "Whether the form control is disabled",
		},
		&AttributeTypeString{
			Name:        "form",
			Description: "Associates the element with a form element",
		},
		&AttributeTypeNumber{
			Name:        "maxlength",
			Description: "Maximum length of value",
		},
		&AttributeTypeNumber{
			Name:        "minlength",
			Description: "Minimum length of value",
		},
		&AttributeTypeString{
			Name:        "name",
			Description: "Name of the element to use for form submission and in the form.elements API",
		},
		&AttributeTypeString{
			Name:        "placeholder",
			Description: "User-visible label to be placed within the form control",
		},
		&AttributeTypeBool{
			Name:        "readonly",
			Description: "Whether to allow the value to be edited by the user",
		},
		&AttributeTypeBool{
			Name:        "required",
			Description: "Whether the control is required for form submission",
		},
		&AttributeTypeString{
			Name:        "rows",
			Description: "Number of lines to show",
		},
		&AttributeTypeEnum{
			Name:        "wrap",
			Description: "How the value of the form control is to be wrapped for form submission",
			Allowed: map[string]struct{}{
				"soft": {},
				"hard": {},
			},
		},
	}
}

func outputAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "for",
			Description: "Specifies controls from which the output was calculated",
		},
		&AttributeTypeString{
			Name:        "form",
			Description: "Associates the element with a form element",
		},
		&AttributeTypeString{
			Name:        "name",
			Description: "Name of the element to use in the form.elements API.",
		},
	}
}

func progressAttr() []Attribute {
	return []Attribute{
		&AttributeTypeFloat{
			Name:        "value",
			Description: "Current value of the element",
		},
		&AttributeTypeFloat{
			Name:        "max",
			Description: "Upper bound of range",
		},
	}
}

func meterAttr() []Attribute {
	return []Attribute{
		&AttributeTypeFloat{
			Name:        "value",
			Description: "Current value of the element",
		},
		&AttributeTypeFloat{
			Name:        "min",
			Description: "Lower bound of range",
		},
		&AttributeTypeFloat{
			Name:        "max",
			Description: "Upper bound of range",
		},
		&AttributeTypeFloat{
			Name:        "low",
			Description: "High limit of low range",
		},
		&AttributeTypeFloat{
			Name:        "high",
			Description: "Low limit of high range",
		},
		&AttributeTypeFloat{
			Name:        "optimum",
			Description: "Optimum value in gauge",
		},
	}
}

func fieldsetAttr() []Attribute {
	return []Attribute{
		&AttributeTypeBool{
			Name:        "disabled",
			Description: "Whether the descendant form controls, except any inside legend, are disabled",
		},
		&AttributeTypeString{
			Name:        "form",
			Description: "Associates the element with a form element",
		},
		&AttributeTypeString{
			Name:        "name",
			Description: "Name of the element to use in the form.elements API.",
		},
	}
}

func detailsAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "name",
			Description: "Name of group of mutually-exclusive details elements",
		},
		&AttributeTypeBool{
			Name:        "open",
			Description: "Whether the details are visible",
		},
	}
}

func dialogAttr() []Attribute {
	return []Attribute{
		&AttributeTypeEnum{
			Name:        "closedby",
			Description: "Which user actions will close the dialog",
			Allowed: map[string]struct{}{
				"any":          {},
				"closerequest": {},
				"none":         {},
			},
		},
		&AttributeTypeBool{
			Name:        "open",
			Description: "Whether the dialog box is showing",
		},
	}
}

func scriptAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "type",
			Description: "Type of script",
		},
		src,
		&AttributeTypeBool{
			Name:        "nomodule",
			Description: "Prevents execution in user agents that support module scripts",
		},
		&AttributeTypeBool{
			Name:        "async",
			Description: "Execute script when available, without blocking while fetching",
		},
		&AttributeTypeBool{
			Name:        "defer",
			Description: "Defer script execution",
		},
		&AttributeTypeEnum{
			Name:        "blocking",
			Description: "Whether the element is potentially render-blocking",
			Allowed: map[string]struct{}{
				"render": {},
			},
		},
		crossorigin,
		referrerPolicy,
		&AttributeTypeString{
			Name:        "integrity",
			Description: "Integrity metadata used in Subresource Integrity checks [SRI]",
		},
		fetchPriority,
	}
}

func templateAttr() []Attribute {
	return []Attribute{
		&AttributeTypeEnum{
			Name:        "shadowrootmode",
			Description: "Enables streaming declarative shadow roots",
			Allowed: map[string]struct{}{
				"open":   {},
				"closed": {},
			},
		},
		&AttributeTypeBool{
			Name:        "shadowrootdelegatefocus",
			Description: "Sets delegates focus on a declarative shadow root",
		},
		&AttributeTypeBool{
			Name:        "shadowrootclonable",
			Description: "Sets clonable on a declarative shadow root",
		},
		&AttributeTypeBool{
			Name:        "shadowrootserializable",
			Description: "Sets serializable on a declarative shadow root",
		},
		&AttributeTypeBool{
			Name:        "shadowrootcustomelementregistry",
			Description: "Enables declarative shadow roots to indicate they will use a custom element registry\n",
		},
	}
}

func slotAttr() []Attribute {
	return []Attribute{
		&AttributeTypeString{
			Name:        "name",
			Description: "Name of shadow tree slot",
		},
	}
}

func canvasAttr() []Attribute {
	return []Attribute{
		width,
		height,
	}
}

var attrFuncs = map[string]func() []Attribute{
	"html":       htmlAttr,
	"head":       headAttr,
	"title":      titleAttr,
	"base":       baseAttr,
	"link":       linkAttr,
	"meta":       metaAttr,
	"style":      styleAttr,
	"body":       bodyAttr,
	"blockquote": blockQuoteAttr,
	"ol":         olAttr,
	"li":         liAttr,
	"a":          aAttr,
	"q":          qAttr,
	"data":       dataAttr,
	"time":       timeAttr,
	"ins":        insDelAttr,
	"del":        insDelAttr,
	"source":     sourceAttr,
	"img":        imgAttr,
	"iframe":     iframeAttr,
	"embed":      embedAttr,
	"object":     objectAttr,
	"video":      videoAttr,
	"audio":      audioAttr,
	"track":      trackAttr,
	"map":        mapAttr,
	"area":       areaAttr,
	"colgroup":   colgroupAttr,
	"col":        colAttr,
	"td":         tdAttr,
	"th":         thAttr,
	"form":       formAttr,
	"label":      labelAttr,
	"input":      inputAttr,
	"button":     buttonAttr,
	"select":     selectAttr,
	"optgroup":   optgroupAttr,
	"option":     optionAttr,
	"textarea":   textareaAttr,
	"output":     outputAttr,
	"progress":   progressAttr,
	"meter":      meterAttr,
	"fieldset":   fieldsetAttr,
	"details":    detailsAttr,
	"dialog":     dialogAttr,
	"script":     scriptAttr,
	"template":   templateAttr,
	"slot":       slotAttr,
	"canvas":     canvasAttr,
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

			if _, ok = getIDIndex(child.Attr, "id", "semantics"); ok {
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

	// Void Elements
	isVoid := []string{
		"area",
		"base",
		"br",
		"col",
		"embed",
		"hr",
		"img",
		"input",
		"link",
		"meta",
		"source",
		"track",
		"wbr",
	}

	for _, e := range p.Spec.Elements {
		if fn, ok := attrFuncs[e.Tag]; ok {
			e.Attributes = append(e.Attributes, fn()...)
		}
		if slices.Contains(isVoid, e.Tag) {
			e.Void = true
		}
	}

	// Manually add h2-h6
	for i := 2; i < 7; i++ {
		e := &Element{
			Tag:         "h" + strconv.Itoa(i),
			Description: "These elements represent headings for their sections.",
		}

		p.Spec.Elements = append(p.Spec.Elements, e)
	}

	return p.Spec, nil
}
