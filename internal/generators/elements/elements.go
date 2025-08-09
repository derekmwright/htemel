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
						List: []*ast.Field{
							&ast.Field{
								Names: []*ast.Ident{
									ast.NewIdent("children"),
								},
								Type: ast.NewIdent("[]htemel.Node"),
							},
						},
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
							Elt: ast.NewIdent("htemel.Node"),
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
				&ast.ReturnStmt{
					Results: []ast.Expr{
						ast.NewIdent("nil"),
					},
				},
			},
		},
	}

	return fn
}

func RenderFunc(e *spec.Element) (*ast.FuncDecl, []ast.Spec) {
	var imports []ast.Spec

	imports = append(imports, &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: `"io"`,
		},
	})

	fn := &ast.FuncDecl{
		Name: ast.NewIdent("Render"),
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						ast.NewIdent("e"),
					},
					Type: ast.NewIdent(`*` + titleCase.String(e.Tag) + `Element`),
				},
			},
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							ast.NewIdent("writer"),
						},
						Type: ast.NewIdent("io.Writer"),
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent("error"),
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
				&ast.ReturnStmt{
					Results: []ast.Expr{
						ast.NewIdent("nil"),
					},
				},
			},
		},
	}

	return fn, imports
}
