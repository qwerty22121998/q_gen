package node

import (
	"q_gen/utils"
	"strings"
)

type joinTable struct {
	Alias string
	Type  string
}

type Tree struct {
	head   Node
	Values []*Variable
	Joins  map[string]joinTable
}

type ValueMap map[string]interface{}

func (v ValueMap) Get(path string) (interface{}, error) {
	paths := strings.Split(path, ".")
	if len(paths) == 0 {
		return v[path], nil
	}
	return utils.ExtractPath(v[path], paths[1:])
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
	return t.head.String()
}
