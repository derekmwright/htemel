package elements

import (
	"go/ast"
	"go/token"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/derekmwright/dwhtml/internal/generators/spec"
)

var titleCase = cases.Title(language.English)

func BaseStruct(e *spec.Element) *ast.GenDecl {
	st := &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(titleCase.String(e.Tag) + `Element`),
				Type: &ast.StructType{
					Fields: &ast.FieldList{
						List: []*ast.Field{},
					},
				},
			},
		},
	}

	return st
}

func BaseFunc(e *spec.Element) *ast.FuncDecl {
	fn := &ast.FuncDecl{
		Name: ast.NewIdent(titleCase.String(e.Tag)),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							ast.NewIdent("children"),
						},
						Type: &ast.Ellipsis{
							Elt: ast.NewIdent("Node"),
						},
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent(`*` + titleCase.String(e.Tag) + `Element`),
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("fmt"),     // Package fmt
							Sel: ast.NewIdent("Println"), // Function Println
						},
						Args: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: `"Hello, World!"`,
							},
						},
					},
				},
			},
		},
	}

	return fn
}
