package matrix

import "testing"

func TestMatrixNew(t *testing.T) {
	m := New(
		[][]float64{
			[]float64{1.0, 2.0, 3.0},
			[]float64{4.0, 5.0, 6.0},
		},
	)

	if !(m.GetAt(0, 0) == 1.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(0, 1) == 2.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(0, 2) == 3.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 0) == 4.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 1) == 5.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 2) == 6.0) {
		t.Error("Wrong value")
	}
}

func TestMatrixAdd(t *testing.T) {
	m1 := New(
		[][]float64{
			[]float64{1.0, 2.0, 3.0},
			[]float64{4.0, 5.0, 6.0},
		},
	)

	m2 := New(
		[][]float64{
			[]float64{2.0, 3.0, 4.0},
			[]float64{5.0, 6.0, 7.0},
		},
	)

	m := Add(m1, m2)

	if !(m.GetAt(0, 0) == 3.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(0, 1) == 5.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(0, 2) == 7.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 0) == 9.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 1) == 11.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 2) == 13.0) {
		t.Error("Wrong value")
	}
}
