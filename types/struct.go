package types

type StructField struct {
	Variable
	Tags    map[string][]string `json:"tags,omitempty"`
	RawTags string              `json:"raw,omitempty"` // Raw string from source.
}

type Struct struct {
	Base
	Fields  []StructField `json:"fields,omitempty"`
	Methods []*Method     `json:"methods,omitempty"`
}
