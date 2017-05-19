package list

import (
	"errors"
	"testing"
)

func TestMap(t *testing.T) {
	list := newOddList()

	evenList, _ := list.Map(mapper)
	for e := range evenList.Iter() {
		if e.(int)%2 != 0 {
			t.Error("Error Map result")
		}
	}
}

func mapper(a interface{}) (interface{}, error) {
	_a, ok := a.(int)
	if ok != true {
		return nil, errors.New("Failed type assertion on a")
	}

	return _a * 2, nil
}
