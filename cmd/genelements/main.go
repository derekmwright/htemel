package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/derekmwright/htemel/internal/generators/elements"
	"github.com/derekmwright/htemel/internal/generators/spec"
)

func main() {
	e := &spec.Element{
		Tag:         "html",
		Description: "Blah blah",
	}

	path, err := filepath.Abs("html/html.go")
	if err != nil {
		panic(err)
	}

	fmt.Println(path)

	buf := &bytes.Buffer{}
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	if err = elements.SourceHeader(
		buf,
		"html",
		e,
		elements.BaseStruct,
		elements.BaseFunc,
		elements.BaseCondFunc,
		elements.RenderFunc,
	); err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
}
