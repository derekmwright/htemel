package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/imports"

	"github.com/derekmwright/htemel/internal/generators/source"
	"github.com/derekmwright/htemel/internal/generators/spec"
)

func generate(pkg string, e *spec.Element) error {
	path, err := filepath.Abs(filepath.Join(pkg, e.Tag+".go"))
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}

	var funcs = []source.TemplateFunc{
		source.BaseStruct,
		source.BaseFunc,
		source.BaseCondFunc,
		source.BaseTernaryFunc,
		source.BuildAttributes(e),
		source.RenderFunc,
	}

	if err = source.SourceHeader(
		buf,
		"html",
		e,
		funcs...,
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

func fmtFiles(dir string) error {
	opt := &imports.Options{
		AllErrors:  true,
		Comments:   true,
		TabIndent:  true,
		TabWidth:   8,
		FormatOnly: false,
	}

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}

		f, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		fixed, err := imports.Process(path, f, opt)
		if err != nil {
			return err
		}

		return os.WriteFile(path, fixed, 0644)
	})
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
		// Merge global attributes into element
		e.Attributes = append(e.Attributes, sp.Attributes...)
		if err = generate(strings.ToLower(sp.Name), e); err != nil {
			panic(err)
		}
	}

	err = fmtFiles("html")
	if err != nil {
		panic(err)
	}
}
