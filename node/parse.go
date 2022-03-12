package node

import (
	"encoding/json"
	"fmt"
	"q_gen/operator"
	"reflect"
)

var (
	ErrInvalidOperator = func(op interface{}) error { return fmt.Errorf("invalid operator %v", op) }
	ErrInvalidOperand  = func(op interface{}) error { return fmt.Errorf("invalid operands %v", op) }
	ErrInvalidType     = func(op interface{}) error { return fmt.Errorf("invalid type %v", op) }
)

func Parse(raw string) (*Tree, error) {
	t := &Tree{}
	if err := t.Parse(raw); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Tree) Parse(raw string) error {
	data := make(map[string]interface{})
	if err := json.Unmarshal([]byte(raw), &data); err != nil {
		return err
	}
	head, err := t.parse(data)
	if err != nil {
		return err
	}
	t.Head = head
	return nil
}

func (t *Tree) parse(v interface{}) (Node, error) {
	kind := reflect.TypeOf(v).Kind()
	switch kind {
	case reflect.String:
		vString, _ := v.(string)
		if len(vString) > 0 && vString[0] == '@' {

			return VAR(vString[1:], nil), nil
		}
		return VAL(vString), nil
	case reflect.Slice:
		arr := reflect.ValueOf(v)
		group := make(Group, 0, arr.Len())
		for i := 0; i < arr.Len(); i++ {
			e := arr.Index(i).Interface()
			eVal, err := t.parse(e)
			if err != nil {
				return nil, err
			}
			group = append(group, eVal)
		}
		return group, nil
	case reflect.Map:
		mp, _ := v.(map[string]interface{})
		op, ok := mp["operator"].(string)
		if !ok {
			return nil, ErrInvalidOperator(mp["operator"])
		}
		data, ok := mp["operands"]
		if !ok {
			return nil, ErrInvalidOperand(mp["operands"])
		}
		nType, ok := mp["type"].(string)
		if !ok {
			return nil, ErrInvalidType(mp["operands"])
		}
		if reflect.TypeOf(data).Kind() != reflect.Slice {
			return nil, ErrInvalidOperand(data)
		}
		sub, err := t.parse(data)
		if err != nil {
			return nil, err
		}
		return Operation{
			Type:     Type(nType),
			Operator: operator.Operator(op),
			Operands: sub.(Group),
		}, nil
	default:
		return VAL(v), nil
	}
}
