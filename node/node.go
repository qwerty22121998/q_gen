package node

type Type string

const (
	Logical Type = "logical"
	Cmp     Type = "comparison"
	Math    Type = "math"
	Target  Type = "target"
	Var     Type = "var"
	Val     Type = "value"
	Many    Type = "many"
	Table   Type = "table"
)

type Node interface {
	T() Type
	String() string
}

type Group []Node

func (g Group) T() Type {
	return Many
}

func (g Group) String() string {
	return ""
}
