package vector

import (
	"math"
	"testing"
)

func TestVectorSort(t *testing.T) {
	t.Parallel()

	v := New(4, 2, 3, 5, 1)
	sorted := Sort(v)

	for i := 0; i < 5; i++ {
		if !(sorted.GetAt(i) == float64(i+1)) {
			t.Error("Wrong value")
		}
	}
}

func TestVectorMagnitude(t *testing.T) {
	t.Parallel()

	v := New(1, 2, 3)
	n := Magnitude(v)

	if math.Abs(n-3.741657) > 1.0/100000 {
		t.Error("Wrong number")
	}
}

func TestVectorAddition(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)

	v := Add(v1, v2)

	if !(v.GetAt(0) == 3.0) {
		t.Log(v.GetAt(0))
		t.Error("Wrong value")
	}

	if !(v.GetAt(1) == 5.0) {
		t.Log(v.GetAt(1))
		t.Error("Wrong value")
	}

	if !(v.GetAt(2) == 7.0) {
		t.Log(v.GetAt(2))
		t.Error("Wrong value")
	}
}

func TestVectorScalarAdd(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)

	v := ScalarAdd(v1, 3)

	if !(v.GetAt(0) == 4) {
		t.Log(v.GetAt(0))
		t.Error("Wrong value")
	}

	if !(v.GetAt(1) == 5) {
		t.Log(v.GetAt(1))
		t.Error("Wrong value")
	}

	if !(v.GetAt(2) == 6) {
		t.Log(v.GetAt(2))
		t.Error("Wrong value")
	}
}

func TestVectorSubstration(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)

	v := Sub(v1, v2)

	if !(v.GetAt(0) == -1) {
		t.Log(v.GetAt(0))
		t.Error("Wrong value")
	}

	if !(v.GetAt(1) == -1) {
		t.Log(v.GetAt(1))
		t.Error("Wrong value")
	}

	if !(v.GetAt(2) == -1) {
		t.Log(v.GetAt(2))
		t.Error("Wrong value")
	}
}

func TestVectorScalarSubFirst(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)

	v := ScalarSubFirst(3, v1)

	if !(v.GetAt(0) == 2) {
		t.Log(v.GetAt(0))
		t.Error("Wrong value")
	}

	if !(v.GetAt(1) == 1) {
		t.Log(v.GetAt(1))
		t.Error("Wrong value")
	}

	if !(v.GetAt(2) == 0) {
		t.Log(v.GetAt(2))
		t.Error("Wrong value")
	}
}

func TestVectorScalarSubSecond(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)

	v := ScalarSubSecond(v1, 3)

	if !(v.GetAt(0) == -2) {
		t.Log(v.GetAt(0))
		t.Error("Wrong value")
	}

	if !(v.GetAt(1) == -1) {
		t.Log(v.GetAt(1))
		t.Error("Wrong value")
	}

	if !(v.GetAt(2) == 0) {
		t.Log(v.GetAt(2))
		t.Error("Wrong value")
	}
}

func TestVectorMultiplication(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)

	v := Mul(v1, v2)

	if !(v.GetAt(0) == 2) {
		t.Log(v.GetAt(0))
		t.Error("Wrong value")
	}

	if !(v.GetAt(1) == 6) {
		t.Log(v.GetAt(1))
		t.Error("Wrong value")
	}

	if !(v.GetAt(2) == 12) {
		t.Log(v.GetAt(2))
		t.Error("Wrong value")
	}
}

func TestVectorScalarMul(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)

	v := ScalarMul(v1, 3)

	if !(v.GetAt(0) == 3) {
		t.Log(v.GetAt(0))
		t.Error("Wrong value")
	}

	if !(v.GetAt(1) == 6) {
		t.Log(v.GetAt(1))
		t.Error("Wrong value")
	}

	if !(v.GetAt(2) == 9) {
		t.Log(v.GetAt(2))
		t.Error("Wrong value")
	}
}

func TestVectorDivision(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)

	v := Div(v1, v2)

	if !(v.GetAt(0) == 0.5) {
		t.Log(v.GetAt(0))
		t.Error("Wrong value")
	}

	if !(v.GetAt(1) == 2.0/3.0) {
		t.Log(v.GetAt(1))
		t.Error("Wrong value")
	}

	if !(v.GetAt(2) == 0.75) {
		t.Log(v.GetAt(2))
		t.Error("Wrong value")
	}
}

func TestVectorPower(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)

	v := Pow(v1, v2)

	if !(v.GetAt(0) == 1) {
		t.Log(v.GetAt(0))
		t.Error("Wrong value")
	}

	if !(v.GetAt(1) == 8) {
		t.Log(v.GetAt(1))
		t.Error("Wrong value")
	}

	if !(v.GetAt(2) == 81) {
		t.Log(v.GetAt(2))
		t.Error("Wrong value")
	}
}

func TestVectorDot(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)

	n := Dot(v1, v2)

	if n != 20 {
		t.Error("Wrong value")
	}
}

func TestVectorCos(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)

	n := Cos(v1, v2)

	if math.Abs(n-0.9925833) > 1.0/1000000 {
		t.Error("Wrong value")
	}
}

func TestVectorMap(t *testing.T) {
	t.Parallel()

	v := New(1, 2, 3)
	out := Map(v, func(n float64) float64 { return n * n })

	if !(out.GetAt(0) == 1) {
		t.Log(v.GetAt(0))
		t.Error("Wrong value")
	}

	if !(out.GetAt(1) == 4) {
		t.Log(v.GetAt(1))
		t.Error("Wrong value")
	}

	if !(out.GetAt(2) == 9) {
		t.Log(v.GetAt(2))
		t.Error("Wrong value")
	}
}

func TestVectorApply(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)

	v := Apply(v1, v2, func(a float64, b float64) float64 { return a + b })

	if !(v.GetAt(0) == 3.0) {
		t.Log(v.GetAt(0))
		t.Error("Wrong value")
	}

	if !(v.GetAt(1) == 5.0) {
		t.Log(v.GetAt(1))
		t.Error("Wrong value")
	}

	if !(v.GetAt(2) == 7.0) {
		t.Log(v.GetAt(2))
		t.Error("Wrong value")
	}
}

func TestReduceByAdd(t *testing.T) {
	t.Parallel()

	v := New(1, 2, 3, 4, 5)
	n := Reduce(v, func(a float64, b float64) float64 { return a + b })
	if n != 15 {
		t.Log(n)
		t.Error("Wrong value")
	}
}
