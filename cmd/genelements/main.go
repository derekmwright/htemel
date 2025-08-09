package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"path/filepath"

	"github.com/derekmwright/dwhtml/internal/generators/elements"
	"github.com/derekmwright/dwhtml/internal/generators/spec"
)

func main() {
	file := ast.File{
		Name:  ast.NewIdent("html"),
		Decls: []ast.Decl{},
	}

	importDecls := &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"fmt"`,
				},
			},
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/derekmwright/htemel"`,
				},
			},
		},
	}

	e := &spec.Element{Tag: "html"}
	structs := elements.BaseStruct(e)
	baseFunc := elements.BaseFunc(e)
	renderFunc, renderImports := elements.RenderFunc(e)
	importDecls.Specs = append(importDecls.Specs, renderImports...)
	file.Decls = append(file.Decls, importDecls, structs, baseFunc, renderFunc)

	fset := token.NewFileSet()

	path, err := filepath.Abs("html/html.go")
	if err != nil {
		panic(err)
	}

	fmt.Println(path)

	buf := &bytes.Buffer{}
	f, err := os.Create(path)
	if err = format.Node(buf, fset, &file); err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
}
