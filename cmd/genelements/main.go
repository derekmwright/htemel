package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/derekmwright/htemel/internal/generators/elements"
	"github.com/derekmwright/htemel/internal/generators/spec"
)

func generate(pkg string, e *spec.Element) error {
	path, err := filepath.Abs(filepath.Join(pkg, e.Tag+".go"))
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}

	if err = elements.SourceHeader(
		buf,
		"html",
		e,
		elements.BaseStruct,
		elements.BaseFunc,
		elements.BaseCondFunc,
		elements.RenderFunc,
	); err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func main() {
	sp := spec.Spec{}

	specFile, err := os.ReadFile("spec/html.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(specFile, &sp); err != nil {
		panic(err)
	}

	for _, e := range sp.Elements {
		if err = generate(strings.ToLower(sp.Name), e); err != nil {
			panic(err)
		}
	}
}
