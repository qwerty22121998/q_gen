package node

import "fmt"

type Field string

func (Field) T() Type {
	return Target
}

func (n Field) String() string {
	return fmt.Sprintf("`%v`", string(n))
}
