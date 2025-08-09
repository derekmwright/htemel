package spec

type Spec struct {
	Name       string
	Elements   []*Element  `json:"elements"`
	Attributes []Attribute `json:"attributes,omitempty"`
}

type Element struct {
	Tag         string       `json:"tag"`
	Description string       `json:"description,omitempty"`
	Attributes  []*Attribute `json:"attributes,omitempty"`

	// A Void element has no children
	Void bool `json:"void,omitempty"`
}

type Attribute interface {
	isAttr()
}

type AttributeTypeString struct {
	Name        string
	Description string
}

func (a AttributeTypeString) isAttr() {}

type AttributeTypeChar struct {
	Name        string
	Description string
}

func (a AttributeTypeChar) isAttr() {}

type AttributeTypeNumber struct {
	Name        string
	Description string
}

func (a AttributeTypeNumber) isAttr() {}

type AttributeTypeBool struct {
	Name        string
	Description string
}

func (a AttributeTypeBool) isAttr() {}

type AttributeTypeEnum struct {
	Name        string
	Description string
	Allowed     map[string]struct{}
}

func (a AttributeTypeEnum) isAttr() {}
