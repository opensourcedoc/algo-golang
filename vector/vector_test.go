package vector

import (
	"errors"
	"math"
	"math/big"
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

func TestAddFloatInt(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2.0, 3.0, 4.0)
	v, _ := v1.Add(v2)

	n0 := v.GetAt(0)
	if n0.(float64) != 3.0 {
		t.Error("Wrong value")
	}

	n1 := v.GetAt(1)
	if n1.(float64) != 5.0 {
		t.Error("Wrong value")
	}

	n2 := v.GetAt(2)
	if n2.(float64) != 7.0 {
		t.Error("Wrong value")
	}
}

func TestUncompatibleType(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(big.NewInt(2), big.NewInt(3), big.NewInt(4))
	_, err := v1.Add(v2)
	if err == nil {
		t.Error("The two vectors should be incompatible")
	}
}

func TestSubFloatInt(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2.0, 3.0, 4.0)
	v, _ := v1.Sub(v2)

	n0 := v.GetAt(0)
	if n0.(float64) != -1.0 {
		t.Error("Wrong value")
	}

	n1 := v.GetAt(1)
	if n1.(float64) != -1.0 {
		t.Error("Wrong value")
	}

	n2 := v.GetAt(2)
	if n2.(float64) != -1.0 {
		t.Error("Wrong value")
	}
}

func TestMulFloatInt(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2.0, 3.0, 4.0)
	v, _ := v1.Mul(v2)

	n0 := v.GetAt(0)
	if n0.(float64) != 2.0 {
		t.Error("Wrong value")
	}

	n1 := v.GetAt(1)
	if n1.(float64) != 6.0 {
		t.Error("Wrong value")
	}

	n2 := v.GetAt(2)
	if n2.(float64) != 12.0 {
		t.Error("Wrong value")
	}
}

func TestDivFloatInt(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2.0, 3.0, 4.0)
	v, _ := v1.Div(v2)

	n0 := v.GetAt(0)
	if n0.(float64) != 0.5 {
		t.Error("Wrong value")
	}

	n1 := v.GetAt(1)
	if !(math.Abs(n1.(float64)-2.0/3.0) < 1.0/1000000) {
		t.Error("Wrong value")
	}

	n2 := v.GetAt(2)
	if n2.(float64) != 0.75 {
		t.Error("Wrong value")
	}
}

func TestApplyAdd(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)
	v, _ := v1.Apply(v2, intAdd)

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
