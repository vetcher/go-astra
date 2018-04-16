package types

import (
	"fmt"
	"strings"
)

type StructField struct {
	Variable
	Tags    map[string][]string `json:"tags,omitempty"`
	RawTags string              `json:"raw,omitempty"` // Raw string from source.
}

func (f StructField) String() string {
	return fmt.Sprintf("%s `%s`", f.Variable.String(), f.RawTags)
}

type Struct struct {
	Base
	Fields  []StructField `json:"fields,omitempty"`
	Methods []*Method     `json:"methods,omitempty"`
	Kind    Kind          `json:"kind"`
}

func (f Struct) TypeOf() Kind {
	return f.Kind
}

func stringFields(fields []StructField) string {
	var str strings.Builder
	if len(fields) == 0 {
		return ""
	}
	for i := range fields {
		str.WriteString("\n")
		str.WriteString(fields[i].String())
	}
	str.WriteString("\n")
	return str.String()
}

func (f Struct) String() string {
	return fmt.Sprintf("%s struct {%s}", f.Name, stringFields(f.Fields))
}
