package source

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/derekmwright/htemel/internal/generators/spec"
)

func BuildAttributes(e *spec.Element) func() (*template.Template, ImportSet) {
	return func() (*template.Template, ImportSet) {
		imports := ImportSet{}

		addTypes := &bytes.Buffer{}
		buf := &bytes.Buffer{}

		buf.WriteString(`
type ` + e.Tag + `Attrs map[string]any
`)

		for _, attr := range e.Attributes {
			switch a := attr.(type) {
			case *spec.AttributeTypeString:
				buf.WriteString(`
func (e *` + titleCase(e.Tag) + `Element) ` + titleCase(a.Name) + `(s string) *` + titleCase(e.Tag) + `Element {
	e.attributes["` + a.Name + `"] = s
	
	return e
}
`)
			case *spec.AttributeTypeChar:

			case *spec.AttributeTypeNumber:

			case *spec.AttributeTypeBool:

			case *spec.AttributeTypeEnum:
				typeName := titleCase(e.Tag) + titleCase(a.Name) + "AttrEnum"

				addTypes.WriteString("\ntype " + typeName + " string\n")
				addTypes.WriteString("\nconst (\n")
				for allowed, _ := range a.Allowed {
					fixed := strings.ReplaceAll(titleCase(allowed), "-", "")
					addTypes.WriteString("\t" + typeName + fixed + " " + typeName + " = \"" + allowed + "\"\n")
				}
				addTypes.WriteString(")\n")

				buf.WriteString(`
func (e *` + titleCase(e.Tag) + `Element) ` + titleCase(a.Name) + `(a ` + typeName + `) *` + titleCase(e.Tag) + `Element {
	e.attributes["` + a.Name + `"] = a
	
	return e
}
`)
			case *spec.AttributeTypeSST:
				buf.WriteString(`
func (e *` + titleCase(e.Tag) + `Element) ` + titleCase(a.Name) + `(s ...string) *` + titleCase(e.Tag) + `Element {
	e.attributes["` + a.Name + `"] = strings.Join(s, " ")
	
	return e
}
`)
			}
		}

		out := &bytes.Buffer{}

		tmpl := template.Must(template.New("BuildAttributes").Parse(addTypes.String()))
		if err := tmpl.Execute(out, e); err != nil {
			panic(err)
		}
		tmpl = template.Must(tmpl.Parse(buf.String()))
		if err := tmpl.Execute(out, e); err != nil {
			panic(err)
		}

		tmpl = template.Must(tmpl.Parse(out.String()))

		return tmpl, imports.Add("strings")
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
