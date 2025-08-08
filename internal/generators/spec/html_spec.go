package spec

import (
	"errors"
	"io"

	"golang.org/x/net/html"
)

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

	body := findTag(doc, "body")
	if body == nil {
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
			if child.Data == "h4" {
				tagNode := findTag(child, "code")
				if tagNode != nil {
					p.Activate(tagNode.FirstChild.Data)
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
