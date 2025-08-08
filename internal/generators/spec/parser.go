package spec

import (
	"slices"
	"strings"

	"golang.org/x/net/html"
)

type Parser struct {
	active      bool
	currElement *Element
	descParsed  bool
	Elements    []*Element
}

func NewSpecParser() *Parser {
	return &Parser{}
}

func (p *Parser) Activate(element string) {
	p.active = true
	p.currElement = &Element{
		Tag: element,
	}
}

func (p *Parser) Reset() {
	p.Elements = append(p.Elements, p.currElement)
	p.active = false
	p.currElement = nil
	p.descParsed = false
}

func findTag(doc *html.Node, tag string) *html.Node {
	if doc == nil {
		return nil
	}

	if doc.Type == html.ElementNode && doc.Data == tag {
		return doc
	}

	for child := range doc.ChildNodes() {
		if result := findTag(child, tag); result != nil {
			return result
		}
	}

	return nil
}

func getIDIndex(attrs []html.Attribute, key, value string) (int, bool) {
	idx := slices.IndexFunc(attrs, func(attr html.Attribute) bool {
		return attr.Key == key
	})

	if idx != -1 {
		if attrs[idx].Val == value {
			return idx, true
		}
	}

	return -1, false
}

func gatherText(node *html.Node, builder *strings.Builder) string {
	if builder == nil {
		builder = &strings.Builder{}
	}

	if node.Type == html.TextNode {
		// TODO: We could probably do better here
		cleaned := strings.ReplaceAll(node.Data, "\n   ", "")
		cleaned = strings.ReplaceAll(cleaned, "\n ", "")
		builder.WriteString(cleaned)
	} else {
		for child := range node.ChildNodes() {
			gatherText(child, builder)
		}
	}

	return builder.String()
}
