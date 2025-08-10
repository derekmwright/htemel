package source

import (
	"bytes"
	"text/template"

	"github.com/derekmwright/htemel/internal/generators/spec"
)

func BuildAttributes(e *spec.Element) func() (*template.Template, ImportSet) {
	return func() (*template.Template, ImportSet) {
		imports := make(ImportSet)

		buf := &bytes.Buffer{}

		for _, attr := range e.Attributes {
			data := struct {
				Tag       string
				Attribute spec.Attribute
			}{
				Tag:       e.Tag,
				Attribute: attr,
			}

			tmpl, imps := AttributeBaseFunc()
			if err := tmpl.Execute(buf, data); err != nil {
				panic(err)
			}

			imports.Merge(imps)
		}

		tmpl := template.Must(template.New("BuildAttributes").Parse(buf.String()))

		return tmpl, imports
	}
}

func AttributeEnumDecl(attribute spec.AttributeTypeEnum) string {
	buf := &bytes.Buffer{}
	tmpl := template.Must(template.New("AttributeEnumDecl").
		Funcs(template.FuncMap{}).
		Parse(`
type {{ .Tag | titleCase }}AttributeValue string

const (
	{{ range $key, $value := .Allowed }}{{ $key }} = "{{ $value }}"
{{ end -}}
)
`))

	if err := tmpl.Execute(buf, attribute); err != nil {
		panic(err)
	}

	return buf.String()
}

func AttributeBaseFunc() (*template.Template, ImportSet) {
	tmpl := template.New("AttributeBaseFunc").Funcs(
		template.FuncMap{
			"titleCase": titleCase,
		},
	)

	tmpl = template.Must(tmpl.Parse(`
func (e *{{ .Tag | titleCase }}Element) {{ .Attribute.Name | titleCase }}() *{{ .Tag | titleCase }}Element {
	return e
}
`))

	return tmpl, nil
}
