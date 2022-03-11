package node

import "fmt"

type Variable struct {
	Name  string
	Value interface{}
}

func (Variable) T() Type {
	return Var
}

func (n Variable) String() string {
	if n.Value == nil {
		return "@" + n.Name
	}
	switch n.Value.(type) {
	case string:
		return fmt.Sprintf(`"%v"`, n.Value)
	default:
		return fmt.Sprint(n.Value)
	}
}

func (n Variable) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"@%v"`, n.Name)), nil
}
