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
				buf.WriteString(`
func (e *` + titleCase(e.Tag) + `Element) ` + titleCase(a.Name) + `(r rune) *` + titleCase(e.Tag) + `Element {
	e.attributes["` + a.Name + `"] = r
	
	return e
}
`)
			case *spec.AttributeTypeNumber:
				buf.WriteString(`
func (e *` + titleCase(e.Tag) + `Element) ` + titleCase(a.Name) + `(i int) *` + titleCase(e.Tag) + `Element {
	e.attributes["` + a.Name + `"] = i
	
	return e
}
`)
			case *spec.AttributeTypeBool:
				buf.WriteString(`
func (e *` + titleCase(e.Tag) + `Element) ` + titleCase(a.Name) + `(b bool) *` + titleCase(e.Tag) + `Element {
	e.attributes["` + a.Name + `"] = b
	
	return e
}
`)
			case *spec.AttributeTypeEnum:
				typeName := titleCase(e.Tag) + titleCase(a.Name) + "Enum"

				addTypes.WriteString("\ntype " + typeName + " string\n")
				addTypes.WriteString("\nconst (\n")
				for allowed := range a.Allowed {
					fixed := strings.ReplaceAll(titleCase(allowed), "-", "")
					addTypes.WriteString("\t" + typeName + fixed + " " + typeName + " = \"" + allowed + "\"\n")
				}
				if a.AllowEmpty {
					addTypes.WriteString("\t" + typeName + "Empty " + typeName + " = \"\"\n")
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
			case *spec.AttributeTypePrefixedCustom:
				buf.WriteString(`
func (e *` + titleCase(e.Tag) + `Element) ` + titleCase(a.Name) + `Unsafe(name string, s string) *` + titleCase(e.Tag) + `Element {
	tag := strings.ToLower("` + a.Name + `-" + name)
	
	e.attributes[tag] = s
	
	return e
}

func (e *` + titleCase(e.Tag) + `Element) ` + titleCase(a.Name) + `(name string, s string) *` + titleCase(e.Tag) + `Element {
	return e.` + titleCase(a.Name) + `Unsafe(name, html.EscapeString(s))
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
