package vector

import "testing"

func TestVectorSort(t *testing.T) {
	v := New(4, 2, 3, 5, 1)
	sorted := v.Sort()

	for i := 0; i < 5; i++ {
		if !(sorted.GetAt(i) == float64(i+1)) {
			t.Error("Wrong value")
		}
	}
}

func TestVectorAddition(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)

	v := Add(v1, v2)

	if !(v.GetAt(0) == 3.0) {
		t.Log(v.GetAt(0))
	}

	if !(v.GetAt(1) == 5.0) {
		t.Log(v.GetAt(1))
	}

	if !(v.GetAt(2) == 7.0) {
		t.Log(v.GetAt(2))
	}
}

func TestVectorSubstration(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)

	v := Sub(v1, v2)

	if !(v.GetAt(0) == -1) {
		t.Log(v.GetAt(0))
	}

	if !(v.GetAt(1) == -1) {
		t.Log(v.GetAt(1))
	}

	if !(v.GetAt(2) == -1) {
		t.Log(v.GetAt(2))
	}
}

func TestVectorMultiplication(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)

	v := Mul(v1, v2)

	if !(v.GetAt(0) == 2) {
		t.Log(v.GetAt(0))
	}

	if !(v.GetAt(1) == 6) {
		t.Log(v.GetAt(1))
	}

	if !(v.GetAt(2) == 12) {
		t.Log(v.GetAt(2))
	}
}

func TestVectorDivision(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)

	v := Div(v1, v2)

	if !(v.GetAt(0) == 0.5) {
		t.Log(v.GetAt(0))
	}

	if !(v.GetAt(1) == 2.0/3.0) {
		t.Log(v.GetAt(1))
	}

	if !(v.GetAt(2) == 0.75) {
		t.Log(v.GetAt(2))
	}
}

func TestVectorPower(t *testing.T) {
	t.Parallel()

	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)

	v := Pow(v1, v2)

	if !(v.GetAt(0) == 1) {
		t.Log(v.GetAt(0))
	}

	if !(v.GetAt(1) == 8) {
		t.Log(v.GetAt(1))
	}

	if !(v.GetAt(2) == 81) {
		t.Log(v.GetAt(2))
	}
}

func TestVectorMap(t *testing.T) {
	t.Parallel()

	v := New(1, 2, 3)
	out := v.Map(func(n float64) float64 { return n * n })

	if !(out.GetAt(0) == 1) {
		t.Log(v.GetAt(0))
	}

	if !(out.GetAt(1) == 4) {
		t.Log(v.GetAt(1))
	}

	if !(out.GetAt(2) == 9) {
		t.Log(v.GetAt(2))
	}
}
