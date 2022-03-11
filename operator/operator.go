package operator

import (
	"fmt"
	"strings"
)

type Operator string

var (
	ErrOperatorNotSupported = func(op Operator) error {
		return fmt.Errorf("operator [%v] not supported", op)
	}
)

const (
	AND Operator = "AND"
	OR  Operator = "OR"
	NOT Operator = "NOT"
	XOR Operator = "XOR"
)

const ()

type Operand interface {
	String() string
}

func (op Operator) Generate(values ...string) string {
	switch op {
	case AND, OR, XOR:
		return gen("", op, values)
	case NOT:
		return gen(op, "", values)
	default:
		return gen("", op, values)
	}
}

func gen(prefix Operator, middle Operator, value []string) string {
	b := strings.Builder{}
	if prefix != "" {
		b.WriteString(string(prefix))
		b.WriteString(" ")
	}
	b.WriteString("(")
	for i := 0; i < len(value)-1; i++ {
		b.WriteString(value[i])
		b.WriteString(" ")
		b.WriteString(string(middle))
		b.WriteString(" ")
	}
	b.WriteString(value[len(value)-1])

	b.WriteString(")")

	return b.String()
}
