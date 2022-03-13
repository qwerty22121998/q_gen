package node

import (
	"q_gen/utils"
	"strings"
)

type Join struct {
	Alias string
	Type  string
}

type Tree struct {
	head    Node
	Values  []*Variable
	joins   map[string]bool
	joinMap map[string]Join
}

type ValueMap map[string]interface{}

func (v ValueMap) Get(path string) (interface{}, error) {
	paths := strings.Split(path, ".")
	if len(paths) == 0 {
		return v[path], nil
	}
	return utils.ExtractPath(v[path], paths[1:])
}

func (t *Tree) SetJoinMap(joinMap map[string]Join) {
	t.joinMap = joinMap
}

func (t *Tree) FillData(data ValueMap) error {
	for _, key := range t.Values {
		value, err := data.Get(key.Name)
		if err != nil {
			return err
		}
		key.Value = value
	}
	return nil
}

func (t *Tree) String() string {
	b := strings.Builder{}
	for name := range t.joins {
		joinOp := t.joinMap[name]
		b.WriteString(joinOp.Type)
		b.WriteString(" JOIN `")
		b.WriteString(joinOp.Alias)
		b.WriteString("` as `")
		b.WriteString(name)
		b.WriteString("` ")
	}

	return b.String() + t.head.String()
}
