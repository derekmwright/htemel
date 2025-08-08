package spec

type Element struct {
	Tag         string            `json:"tag"`
	Description string            `json:"description,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
}
