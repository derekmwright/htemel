package spec

type Spec struct {
	Name       string
	Elements   []*Element   `json:"elements"`
	Attributes []*Attribute `json:"attributes,omitempty"`
}

type Element struct {
	Tag         string       `json:"tag"`
	Description string       `json:"description,omitempty"`
	Attributes  []*Attribute `json:"attributes,omitempty"`
	
	// A Void element has no children
	Void bool `json:"void,omitempty"`
}

type AttributeType string

const (
	AttributeTypeString  AttributeType = "string"
	AttributeTypeNumber  AttributeType = "number"
	AttributeTypeBoolean AttributeType = "boolean"
	AttributeTypeEnum    AttributeType = "enum"
)

type Attribute struct {
	Name        string
	Description string
	Type        AttributeType
}
