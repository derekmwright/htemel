package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"os"

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
					Value: `"errors"`,
				},
			},
		},
	}

	e := &spec.Element{Tag: "html"}
	structs := elements.BaseStruct(e)
	mainFunc := elements.BaseFunc(e)
	file.Decls = append(file.Decls, importDecls, structs, mainFunc)

	fset := token.NewFileSet()
	var buf bytes.Buffer

	err := format.Node(&buf, fset, &file)
	if err != nil {
		panic(err)
	}

	_, err = os.Stdout.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
}
