package spec

import (
	"encoding/json"
)

type Spec struct {
	Name       string      `json:"name"`
	Elements   []*Element  `json:"elements"`
	Attributes []Attribute `json:"attributes,omitempty"`
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

	var tmpAttr struct {
		Name          string `json:"name"`
		Description   string `json:"description"`
		AttributeType string `json:"attribute_type"`
	}

	for _, attr := range tmp.Attributes {
		if err := json.Unmarshal(attr, &tmpAttr); err != nil {
			return err
		}

		switch tmpAttr.AttributeType {
		case "AttributeTypeString":
			a := &AttributeTypeString{}
			if err := json.Unmarshal(attr, &a); err != nil {
				return err
			}
			sp.Attributes = append(sp.Attributes, a)
		case "AttributeType":
			a := &AttributeTypeChar{}
			if err := json.Unmarshal(attr, &a); err != nil {
				return err
			}
			sp.Attributes = append(sp.Attributes, a)
		case "AttributeTypeNumber":
			a := &AttributeTypeNumber{}
			if err := json.Unmarshal(attr, &a); err != nil {
				return err
			}
			sp.Attributes = append(sp.Attributes, a)
		case "AttributeTypeBool":
			a := &AttributeTypeBool{}
			if err := json.Unmarshal(attr, &a); err != nil {
				return err
			}
			sp.Attributes = append(sp.Attributes, a)
		case "AttributeTypeEnum":
			a := &AttributeTypeEnum{}
			if err := json.Unmarshal(attr, &a); err != nil {
				return err
			}
			sp.Attributes = append(sp.Attributes, a)
		case "AttributeTypeSST":
			a := &AttributeTypeSST{}
			if err := json.Unmarshal(attr, &a); err != nil {
				return err
			}
			sp.Attributes = append(sp.Attributes, a)
		}
	}

	return nil
}

type Element struct {
	Tag         string      `json:"tag"`
	Description string      `json:"description,omitempty"`
	Attributes  []Attribute `json:"attributes,omitempty"`

	// A Void element has no children
	Void bool `json:"void,omitempty"`
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
		AttributeType string              `json:"attribute_type"`
	}{
		Name:          a.Name,
		Description:   a.Description,
		Allowed:       a.Allowed,
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
