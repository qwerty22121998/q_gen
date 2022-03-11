package node

type Type string

const (
	Logical Type = "logical"
	Cmp     Type = "comparison"
	Math    Type = "math"
	Target  Type = "target"
	Var     Type = "var"
	Val     Type = "value"
)

type Node interface {
	T() Type
	String() string
}
