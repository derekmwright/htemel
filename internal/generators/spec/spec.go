package spec

import (
	"encoding/json"
)

type Spec struct {
	Name       string      `json:"name"`
	Elements   []*Element  `json:"elements"`
	Attributes []Attribute `json:"attributes,omitempty"`
}

func attrUnmarshal(in []json.RawMessage) ([]Attribute, error) {
	out := make([]Attribute, 0)

	var tmpAttr struct {
		Name          string `json:"name"`
		Description   string `json:"description"`
		AttributeType string `json:"attribute_type"`
	}

	for _, attr := range in {
		if err := json.Unmarshal(attr, &tmpAttr); err != nil {
			return nil, err
		}

		switch tmpAttr.AttributeType {
		case "AttributeTypeString":
			a := &AttributeTypeString{}
			if err := json.Unmarshal(attr, &a); err != nil {
				return nil, err
			}
			out = append(out, a)
		case "AttributeType":
			a := &AttributeTypeChar{}
			if err := json.Unmarshal(attr, &a); err != nil {
				return nil, err
			}
			out = append(out, a)
		case "AttributeTypeNumber":
			a := &AttributeTypeNumber{}
			if err := json.Unmarshal(attr, &a); err != nil {
				return nil, err
			}
			out = append(out, a)
		case "AttributeTypeFloat":
			a := &AttributeTypeFloat{}
			if err := json.Unmarshal(attr, &a); err != nil {
				return nil, err
			}
			out = append(out, a)
		case "AttributeTypeBool":
			a := &AttributeTypeBool{}
			if err := json.Unmarshal(attr, &a); err != nil {
				return nil, err
			}
			out = append(out, a)
		case "AttributeTypeEnum":
			a := &AttributeTypeEnum{}
			if err := json.Unmarshal(attr, &a); err != nil {
				return nil, err
			}
			out = append(out, a)
		case "AttributeTypeSST":
			a := &AttributeTypeSST{}
			if err := json.Unmarshal(attr, &a); err != nil {
				return nil, err
			}
			out = append(out, a)
		case "AttributeTypePrefixedCustom":
			a := &AttributeTypePrefixedCustom{}
			if err := json.Unmarshal(attr, &a); err != nil {
				return nil, err
			}
			out = append(out, a)
		}
	}

	return out, nil
}

func (sp *Spec) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Name       string            `json:"name"`
		Elements   []*Element        `json:"elements"`
		Attributes []json.RawMessage `json:"attributes,omitempty"`
	}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	sp.Name = tmp.Name
	sp.Elements = tmp.Elements
	attrs, err := attrUnmarshal(tmp.Attributes)
	if err != nil {
		return err
	}
	sp.Attributes = attrs

	return nil
}

type Element struct {
	Tag         string      `json:"tag"`
	Description string      `json:"description,omitempty"`
	Attributes  []Attribute `json:"attributes,omitempty"`

	// A Void element has no children
	Void bool `json:"void,omitempty"`
}

func (e *Element) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Tag         string            `json:"tag"`
		Description string            `json:"description,omitempty"`
		Attributes  []json.RawMessage `json:"attributes,omitempty"`
		Void        bool              `json:"void,omitempty"`
	}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	e.Tag = tmp.Tag
	e.Description = tmp.Description
	e.Void = tmp.Void
	attrs, err := attrUnmarshal(tmp.Attributes)
	if err != nil {
		return err
	}
	e.Attributes = attrs

	return nil
}

type Attribute interface {
	isAttr()
	GetName() string
}

type AttributeTypeString struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a AttributeTypeString) isAttr() {}

func (a AttributeTypeString) GetName() string {
	return a.Name
}

func (a AttributeTypeString) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name          string `json:"name"`
		Description   string `json:"description,omitempty"`
		AttributeType string `json:"attribute_type"`
	}{
		Name:          a.Name,
		Description:   a.Description,
		AttributeType: "AttributeTypeString",
	})
}

type AttributeTypeChar struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a AttributeTypeChar) isAttr() {}

func (a AttributeTypeChar) GetName() string {
	return a.Name
}

func (a AttributeTypeChar) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name          string `json:"name"`
		Description   string `json:"description,omitempty"`
		AttributeType string `json:"attribute_type"`
	}{
		Name:          a.Name,
		Description:   a.Description,
		AttributeType: "AttributeTypeChar",
	})
}

type AttributeTypeNumber struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a AttributeTypeNumber) isAttr() {}

func (a AttributeTypeNumber) GetName() string {
	return a.Name
}

func (a AttributeTypeNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name          string `json:"name"`
		Description   string `json:"description,omitempty"`
		AttributeType string `json:"attribute_type"`
	}{
		Name:          a.Name,
		Description:   a.Description,
		AttributeType: "AttributeTypeNumber",
	})
}

type AttributeTypeFloat struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a AttributeTypeFloat) isAttr() {}

func (a AttributeTypeFloat) GetName() string {
	return a.Name
}

func (a AttributeTypeFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name          string `json:"name"`
		Description   string `json:"description,omitempty"`
		AttributeType string `json:"attribute_type"`
	}{
		Name:          a.Name,
		Description:   a.Description,
		AttributeType: "AttributeTypeFloat",
	})
}

type AttributeTypeBool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a AttributeTypeBool) isAttr() {}

func (a AttributeTypeBool) GetName() string {
	return a.Name
}

func (a AttributeTypeBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name          string `json:"name"`
		Description   string `json:"description,omitempty"`
		AttributeType string `json:"attribute_type"`
	}{
		Name:          a.Name,
		Description:   a.Description,
		AttributeType: "AttributeTypeBool",
	})
}

type AttributeTypeEnum struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Allowed     map[string]struct{} `json:"allowed"`
	AllowCustom bool                `json:"allow_custom"`
	AllowEmpty  bool                `json:"allow_empty"`
}

func (a AttributeTypeEnum) isAttr() {}

func (a AttributeTypeEnum) GetName() string {
	return a.Name
}

func (a AttributeTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name          string              `json:"name"`
		Description   string              `json:"description,omitempty"`
		Allowed       map[string]struct{} `json:"allowed"`
		AllowEmpty    bool                `json:"allow_empty"`
		AllowCustom   bool                `json:"allow_custom"`
		AttributeType string              `json:"attribute_type"`
	}{
		Name:          a.Name,
		Description:   a.Description,
		Allowed:       a.Allowed,
		AllowEmpty:    a.AllowEmpty,
		AllowCustom:   a.AllowCustom,
		AttributeType: "AttributeTypeEnum",
	})
}

type AttributeTypeSST struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a AttributeTypeSST) isAttr() {}

func (a AttributeTypeSST) GetName() string {
	return a.Name
}

func (a AttributeTypeSST) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name          string `json:"name"`
		Description   string `json:"description,omitempty"`
		AttributeType string `json:"attribute_type"`
	}{
		Name:          a.Name,
		Description:   a.Description,
		AttributeType: "AttributeTypeSST",
	})
}

// AttributeTypePrefixedCustom provides support for attributes like `data-<user defined>`
type AttributeTypePrefixedCustom struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a AttributeTypePrefixedCustom) isAttr() {}

func (a AttributeTypePrefixedCustom) GetName() string {
	return a.Name
}

func (a AttributeTypePrefixedCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name          string `json:"name"`
		Description   string `json:"description,omitempty"`
		AttributeType string `json:"attribute_type"`
	}{
		Name:          a.Name,
		Description:   a.Description,
		AttributeType: "AttributeTypePrefixedCustom",
	})
}
