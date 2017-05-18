package list

import (
	"errors"
	"testing"
)

func add(a interface{}, b interface{}) (interface{}, error) {
	_a, ok := a.(int)
	if ok != true {
		return nil, errors.New("Failed type assertion on a")
	}

	_b, ok := b.(int)
	if ok != true {
		return nil, errors.New("Failed type assertion on b")
	}

	return _a + _b, nil
}

func TestReduce(t *testing.T) {
	list := new1To10List()

	result, _ := list.Reduce(add)
	if result != 55 {
		t.Error("Error Reduce")
	}
}

func TestReduceSingle(t *testing.T) {
	list := New()
	list.Unshift(100)

	result, _ := list.Reduce(add)
	if result != 100 {
		t.Error("Error Reduce on single element")
	}
}
