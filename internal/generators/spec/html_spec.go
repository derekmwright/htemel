package spec

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

func GenerateHTMLSpec(url string, w io.Writer) error {
	p := NewSpecParser()

	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	doc, err := html.Parse(req.Body)
	if err != nil {
		panic(err)
	}

	body := findTag(doc, "body")
	if body == nil {
		return errors.New("could not find body")
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

	spec, err := json.MarshalIndent(p.Elements, "", "  ")
	if err != nil {
		return err
	}

	_, err = w.Write(spec)

	return err
}
