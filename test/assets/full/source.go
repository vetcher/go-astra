// This is a file documentation.
package full

// This is block comment for imports.
import (
	"context"
	"fmt"

	// This is documentation comment for package import
	thisisstubalias "github.com/vetcher/go-astra/test/assets/full/thisisstubpackage" // This is inline comment for package import
)

// This is a comment for string constant
const ConstString = "this is string const"

const ConstInt = 5 // This is inline comment.

// This is a block comment.
const (
	ConstBlock1 uint32  = 7
	ConstBlock2 float32 = 1.0
)

const (
	Iota1 = iota + 1
	Iota2
	Iota3
)

var VarA string

var VarB = "var b"

var VarC string = "var c"

// Block comment of variables.
var (
	BlockVarA = func(string) string { return "" }
	BlockVarB chan error
	BlockVarC func(string) string = BlockVarA
)

type StructOne struct {
	ExportedField string
	privateField  int
	FieldWithTags int `json:"field_with_tags" sometag:"param1,param2,param3"`
	// Documentation of complex field.
	ComplexField chan<- *[]**map[interface {
		InterfaceMethod(uint, ...complex64)
	}]func(int, string, [7]byte) (complex64, error) // Inline comment of complex field.
}

type (
	StructTwo struct {
		FieldOne   thisisstubalias.ThisIsStubStructure
		FieldTwo   []thisisstubalias.ThisIsStubStructure
		FieldThree *StructTwo
		FieldFour  []StructTwo
	}

	StructThree struct {
		StructTwo
		ExtendingField string
	}

	InterfaceOne interface {
		InterfaceMethod(uint, ...complex64)
	}
)

func FunctionOne(a string, b interface{}, c map[string]interface{}) (ctx context.Context, err error) {
	return nil, nil
}

func FunctionTwo(f func(string, func() error)) {
	return
}

func (m StructTwo) MethodOne() string {
	return fmt.Sprint(m)
}

type (
	X int
	Y string
)
