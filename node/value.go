package node

import "fmt"

type Value struct {
	Value interface{}
}

func (Value) T() Type {
	return Val
}

func (n Value) String() string {
	val, _ := n.MarshalJSON()
	return string(val)
}

func (n Value) MarshalJSON() ([]byte, error) {
	switch n.Value.(type) {
	case string:
		return []byte(fmt.Sprintf(`"%v"`, n.Value)), nil
	default:
		return []byte(fmt.Sprint(n.Value)), nil
	}
}
