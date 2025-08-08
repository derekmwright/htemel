package spec

type Spec struct {
	Name             string
	Elements         []*Element `json:"elements"`
	GlobalAttributes []any      `json:"globalAttributes"`
}

type Element struct {
	Tag         string            `json:"tag"`
	Description string            `json:"description,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
}
