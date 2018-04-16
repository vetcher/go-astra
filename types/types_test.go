package types

import (
	"fmt"
	"testing"
)

type stringerTests struct {
	Name     string
	Stringer fmt.Stringer
	Result   string
}

func TestStringMethods(test *testing.T) {
	tt := []stringerTests{
		{
			Name: "Type name",
			Stringer: TName{
				TypeName: "NameTest",
			},
			Result: "NameTest",
		},
		{
			Name: "Type map",
			Stringer: TMap{
				Key: TName{
					TypeName: "key",
				},
				Value: TName{
					TypeName: "value",
				},
			},
			Result: "map[key]value",
		},
		{
			Name: "Type pointer 1",
			Stringer: TPointer{
				NumberOfPointers: 4,
				Next: TMap{
					Key: TName{
						TypeName: "key",
					},
					Value: TName{
						TypeName: "value",
					},
				},
			},
			Result: "****map[key]value",
		},
		{
			Name: "Type import",
			Stringer: Import{
				Base: Base{
					Name: "alias",
				},
				Package: "blabla.blo/foo/bar",
			},
			Result: "alias \"blabla.blo/foo/bar\"",
		},
		{
			Name: "Type array",
			Stringer: TArray{
				ArrayLen: 10,
				Next: TName{
					TypeName: "string",
				},
			},
			Result: "[10]string",
		},
		{
			Name: "Type array slice",
			Stringer: TArray{
				IsSlice: true,
				Next: TName{
					TypeName: "string",
				},
			},
			Result: "[]string",
		},
		{
			Name: "Type array ellipsis",
			Stringer: TArray{
				IsEllipsis: true,
				Next: TName{
					TypeName: "string",
				},
			},
			Result: "...string",
		},
		{
			Name:     "Type empty struct",
			Stringer: Struct{},
			Result:   " struct {}",
		},
		{
			Name: "Type struct",
			Stringer: Struct{
				Fields: []StructField{
					{
						Variable: Variable{
							Base: Base{
								Name: "PublicField",
							},
							Type: TName{
								TypeName: "string",
							},
						},
						RawTags: `json:"public_field,omitempty"`,
					},
					{
						Variable: Variable{
							Base: Base{
								Name: "privateField",
							},
							Type: TName{
								TypeName: "string",
							},
						},
						RawTags: `sql:"-"`,
					},
				},
			},
			Result: " struct {\n" +
				"PublicField string `json:\"public_field,omitempty\"`\n" +
				"privateField string `sql:\"-\"`\n" +
				"}",
		},
	}
	for _, t := range tt {
		test.Run(t.Name, func(test *testing.T) {
			s := t.Stringer.String()
			if t.Result != s {
				test.Error("has", s, "want", t.Result)
			}
		})
	}
}
