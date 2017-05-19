package vector

import (
	"errors"
	"testing"
)

func TestGetAt(t *testing.T) {
	t.Parallel()

	v := New(1, 2, 3, 4, 5)
	n := v.GetAt(2)
	if n.(int) != 3 {
		t.Error("Wrong value")
	}
}

func TestSetAt(t *testing.T) {
	t.Parallel()

	v := New(1, 2, 3, 4, 5)
	v.SetAt(3, 99)

	n := v.GetAt(3)
	if n.(int) != 99 {
		t.Error("Wrong value")
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)
	v, _ := v1.Add(v2)

	n0 := v.GetAt(0)
	if n0.(int) != 3 {
		t.Error("Wrong value")
	}

	n1 := v.GetAt(1)
	if n1.(int) != 5 {
		t.Error("Wrong value")
	}

	n2 := v.GetAt(2)
	if n2.(int) != 7 {
		t.Error("Wrong value")
	}
}

func TestCalcByAdd(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)
	v, _ := v1.CalcBy(v2, intAdd)

	n0 := v.GetAt(0)
	if n0.(int) != 3 {
		t.Error("Wrong value")
	}

	n1 := v.GetAt(1)
	if n1.(int) != 5 {
		t.Error("Wrong value")
	}

	n2 := v.GetAt(2)
	if n2.(int) != 7 {
		t.Error("Wrong value")
	}
}

func intAdd(a interface{}, b interface{}) (interface{}, error) {
	na, ok := a.(int)
	if !ok {
		return 0, errors.New("Wrong value a")
	}

	nb, ok := b.(int)
	if !ok {
		return 0, errors.New("Wrong value b")
	}

	return na + nb, nil
}
