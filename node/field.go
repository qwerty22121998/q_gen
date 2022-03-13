package node

import "fmt"

type Field struct {
	Type Type   `json:"type"`
	Name string `json:"name"`
}

func (Field) T() Type {
	return Target
}

func (n Field) String() string {
	return fmt.Sprintf("`%v`", n.Name)
}
