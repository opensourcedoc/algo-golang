package matrix

import (
	"sync"
)

type IMatrix interface {
	Row() int
	Col() int
	GetAt(int, int) float64
	SetAt(int, int, float64)
}

type Matrix struct {
	sync.RWMutex
	row int
	col int
	mtx []float64
}

func New(mtx [][]float64) IMatrix {
	row := len(mtx)
	col := len(mtx[0])

	for i := 1; i < row; i++ {
		if len(mtx[i]) != col {
			panic("Unequal column size")
		}
	}

	m := WithSize(row, col)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m.SetAt(i, j, mtx[i][j])
		}
	}

	return m
}

func WithSize(row int, col int) IMatrix {
	if row < 0 {
		panic("Invalid row size")
	}

	if col < 0 {
		panic("Invalid column size")
	}

	m := new(Matrix)

	m.row = row
	m.col = col
	m.mtx = make([]float64, m.row*m.col)

	return m
}

// The row size of the matrix.
func (m *Matrix) Row() int {
	return m.row
}

// The column size of the matrix.
func (m *Matrix) Col() int {
	return m.col
}

// Getter
func (m *Matrix) GetAt(r int, c int) float64 {
	if r < 0 || r >= m.Row() {
		panic("Invalid row size")
	}

	if c < 0 || c >= m.Col() {
		panic("Invalid column size")
	}

	return m.mtx[r*m.Col()+c]
}

// Setter
func (m *Matrix) SetAt(r int, c int, data float64) {
	if r < 0 || r >= m.Row() {
		panic("Invalid row size")
	}

	if c < 0 || c >= m.Col() {
		panic("Invalid column size")
	}

	m.Lock()
	m.mtx[r*m.Col()+c] = data
	m.Unlock()
}

// Matrix element-wise addition
func Add(m1 IMatrix, m2 IMatrix) IMatrix {
	row := m1.Row()
	col := m1.Col()

	if row != m2.Row() {
		panic("Unequal row size")
	}

	if col != m2.Col() {
		panic("Unequal column size")
	}

	out := WithSize(row, col)

	var wg sync.WaitGroup

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			wg.Add(1)

			go func(m1 IMatrix, m2 IMatrix, out IMatrix, i int, j int) {
				defer wg.Done()

				a := m1.GetAt(i, j)
				b := m2.GetAt(i, j)

				out.SetAt(i, j, a+b)
			}(m1, m2, out, i, j)
		}
	}

	wg.Wait()

	return out
}
