package types

import (
	"strconv"
	"strings"
)

//go:generate stringer -type=Kind

type Kind int32

const (
	KindName Kind = iota
	KindPointer
	KindArray
	KindMap
	KindInterface
	KindImport
	KindEllipsis
	KindChan
	KindFunc
	KindStruct
)

type Type interface {
	TypeOf() Kind
	String() string
}

type LinearType interface {
	NextType() Type
}

type TInterface struct {
	Kind      Kind       `json:"kind"`
	Interface *Interface `json:"interface,omitempty"`
}

func (i TInterface) TypeOf() Kind {
	return i.Kind
}

func (i TInterface) String() string {
	if i.Interface != nil {
		return i.Interface.String()
	}
	return ""
}

type TMap struct {
	Kind  Kind `json:"kind"`
	Key   Type `json:"key,omitempty"`
	Value Type `json:"value,omitempty"`
}

func (m TMap) TypeOf() Kind {
	return m.Kind
}

func (m TMap) String() string {
	return "map[" + m.Key.String() + "]" + m.Value.String()
}

type TName struct {
	Kind     Kind   `json:"kind"`
	TypeName string `json:"type_name,omitempty"`
}

func (i TName) TypeOf() Kind {
	return i.Kind
}

func (i TName) String() string {
	return i.TypeName
}

func (i TName) NextType() Type {
	return nil
}

type TPointer struct {
	Kind             Kind `json:"kind"`
	NumberOfPointers int  `json:"number_of_pointers,omitempty"`
	Next             Type `json:"next,omitempty"`
}

func (i TPointer) TypeOf() Kind {
	return i.Kind
}

func (i TPointer) String() string {
	str := strings.Repeat("*", i.NumberOfPointers)
	if i.Next != nil {
		str += i.Next.String()
	}
	return str
}

func (i TPointer) NextType() Type {
	return i.Next
}

type TArray struct {
	Kind       Kind `json:"kind"`
	ArrayLen   int  `json:"array_len,omitempty"`
	IsSlice    bool `json:"is_slice,omitempty"` // [] declaration
	IsEllipsis bool `json:"is_ellipsis,omitempty"`
	Next       Type `json:"next,omitempty"`
}

func (i TArray) TypeOf() Kind {
	return i.Kind
}

func (i TArray) String() string {
	str := ""
	if i.IsEllipsis {
		str += "..."
	} else if i.IsSlice {
		str += "[]"
	} else {
		str += "[" + strconv.Itoa(i.ArrayLen) + "]"
	}
	if i.Next != nil {
		str += i.Next.String()
	}
	return str
}

func (i TArray) NextType() Type {
	return i.Next
}

type TImport struct {
	Kind   Kind    `json:"kind"`
	Import *Import `json:"import,omitempty"`
	Next   Type    `json:"next,omitempty"`
}

func (i TImport) TypeOf() Kind {
	return i.Kind
}

func (i TImport) String() string {
	str := ""
	if i.Import != nil {
		str += i.Import.Name + "."
	}
	if i.Next != nil {
		str += i.Next.String()
	}
	return str
}

func (i TImport) NextType() Type {
	return i.Next
}

// TEllipsis used only for function params in declarations like `strs ...string`
type TEllipsis struct {
	Kind Kind `json:"kind"`
	Next Type `json:"next,omitempty"`
}

func (i TEllipsis) TypeOf() Kind {
	return i.Kind
}

func (i TEllipsis) String() string {
	str := "..."
	if i.Next != nil {
		return str + i.Next.String()
	}
	return str
}

func (i TEllipsis) NextType() Type {
	return i.Next
}

const (
	ChanDirSend = 1
	ChanDirRecv = 2
	ChanDirAny  = ChanDirSend | ChanDirRecv
)

type TChan struct {
	Kind      Kind `json:"kind"`
	Direction int  `json:"direction"`
	Next      Type `json:"next"`
}

func (c TChan) TypeOf() Kind {
	return c.Kind
}

func (c TChan) NextType() Type {
	return c.Next
}

var strForChan = map[int]string{
	ChanDirSend: "chan<-",
	ChanDirRecv: "<-chan",
	ChanDirAny:  "chan",
}

func (c TChan) String() string {
	str := strForChan[c.Direction]
	if c.Next != nil {
		return str + " " + c.Next.String()
	}
	return str
}
