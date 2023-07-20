package connectc

import (
	"errors"
	"fmt"
)

type VariableRange struct {
	// range is inclusive
	startIndex int
	endIndex int

	Name string
}

func (vr VariableRange) Len() int {
	if vr.endIndex < vr.startIndex {
		panic("endIndex is lower than startIndex")
	}
	return vr.endIndex - vr.startIndex + 1;
}

func (vr VariableRange) Var(index int) (int, error) {
	if index < 0 || index >= vr.Len() {
		return 0, errors.New(fmt.Sprintf("Index %d is out of range.", index))
	}
	return vr.startIndex + index, nil
}

type VariablePool struct {
	lastIndex int
	ranges map[string]VariableRange
}

func (v *VariablePool) AddRange(name string, size int) (VariableRange, error) {
	if v.ranges == nil {
		v.ranges = make(map[string]VariableRange)
	} else {
		if _, ok := v.ranges[name]; ok {
			return VariableRange{}, errors.New(fmt.Sprintf("VariableRange with name %s already exists.", name))
		}
	}

	r := VariableRange {
		Name: name,
		startIndex: v.lastIndex,
		endIndex: v.lastIndex + size - 1,
	}

	v.ranges[name] = r
	v.lastIndex += size

	return r, nil
}

func (v *VariablePool) GetRange(name string) (VariableRange, error) {
	vr, ok := v.ranges[name];

	if ok {
		return vr, nil
	} else {
		return VariableRange{}, errors.New(fmt.Sprintf("Variable range with name %s does not exist in this Pool.", name))
	}
}

