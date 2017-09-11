package vector

import (
	"math"
	"sync"
)

type IVector interface {
	Len() int
	GetAt(int) float64
	SetAt(int, float64)
	Sort() IVector
	Map(func(float64) float64) IVector
}

type Vector struct {
	sync.RWMutex
	vec []float64
}

func New(args ...float64) IVector {
	v := new(Vector)
	v.vec = make([]float64, len(args))

	for i, e := range args {
		v.SetAt(i, e)

	}

	return v
}

var mutex = &sync.Mutex{}

func WithSize(size int) IVector {
	v := new(Vector)
	v.vec = make([]float64, size)

	for i := 0; i < size; i++ {
		v.SetAt(i, 0.0)
	}

	return v
}

func FromArray(arr []float64) IVector {
	v := new(Vector)
	v.vec = make([]float64, len(arr))

	for i, e := range arr {
		v.SetAt(i, e)
	}

	return v
}

// The length of the vector
func (v *Vector) Len() int {
	return len(v.vec)
}

// Getter
func (v *Vector) GetAt(i int) float64 {
	if i < 0 || i >= v.Len() {
		panic("Index out of range")
	}

	return v.vec[i]
}

// Setter
func (v *Vector) SetAt(i int, data float64) {
	if i < 0 || i >= v.Len() {
		panic("Index out of range")
	}

	v.Lock()
	v.vec[i] = data
	v.Unlock()
}

func (v *Vector) Sort() IVector {
	if v.Len() == 0 {
		return v
	}

	arr := make([]float64, 1)
	arr[0] = v.GetAt(0)

	for i := 1; i < v.Len(); i++ {
		inserted := false

		for j := 0; j < len(arr); j++ {
			if v.GetAt(i) < arr[j] {
				if j == 0 {
					arr = append([]float64{v.GetAt(i)}, arr...)
				} else {
					arr = append(arr[:j], append([]float64{v.GetAt(i)}, arr[j:]...)...)
				}

				inserted = true
				break
			}
		}

		if !inserted {
			arr = append(arr, v.GetAt(i))
		}
	}

	return FromArray(arr)
}

// Vector addition
func Add(v1 IVector, v2 IVector) IVector {
	_len := v1.Len()

	if !(_len == v2.Len()) {
		panic("Unequal vector size")
	}

	out := WithSize(_len)

	var wg sync.WaitGroup

	for i := 0; i < _len; i++ {
		wg.Add(1)

		go func(v1 IVector, v2 IVector, out IVector, i int) {
			defer wg.Done()

			a := v1.GetAt(i)
			b := v2.GetAt(i)

			out.SetAt(i, a+b)

		}(v1, v2, out, i)
	}

	wg.Wait()

	return out
}

// Vector substraction
func Sub(v1 IVector, v2 IVector) IVector {
	_len := v1.Len()

	if !(_len == v2.Len()) {
		panic("Unequal vector size")
	}

	out := WithSize(_len)

	var wg sync.WaitGroup

	for i := 0; i < _len; i++ {
		wg.Add(1)

		go func(v1 IVector, v2 IVector, out IVector, i int) {
			defer wg.Done()

			a := v1.GetAt(i)
			b := v2.GetAt(i)

			out.SetAt(i, a-b)

		}(v1, v2, out, i)
	}

	wg.Wait()

	return out
}

// Vector multiplication
func Mul(v1 IVector, v2 IVector) IVector {
	_len := v1.Len()

	if !(_len == v2.Len()) {
		panic("Unequal vector size")
	}

	out := WithSize(_len)

	var wg sync.WaitGroup

	for i := 0; i < _len; i++ {
		wg.Add(1)

		go func(v1 IVector, v2 IVector, out IVector, i int) {
			defer wg.Done()

			a := v1.GetAt(i)
			b := v2.GetAt(i)

			out.SetAt(i, a*b)

		}(v1, v2, out, i)
	}

	wg.Wait()

	return out
}

// Vector division
func Div(v1 IVector, v2 IVector) IVector {
	_len := v1.Len()

	if !(_len == v2.Len()) {
		panic("Unequal vector size")
	}

	out := WithSize(_len)

	var wg sync.WaitGroup

	for i := 0; i < _len; i++ {
		wg.Add(1)

		go func(v1 IVector, v2 IVector, out IVector, i int) {
			defer wg.Done()

			a := v1.GetAt(i)
			b := v2.GetAt(i)

			out.SetAt(i, a/b)

		}(v1, v2, out, i)
	}

	wg.Wait()

	return out
}

// Vector power
func Pow(v1 IVector, v2 IVector) IVector {
	_len := v1.Len()

	if !(_len == v2.Len()) {
		panic("Unequal vector size")
	}

	out := WithSize(_len)

	var wg sync.WaitGroup

	for i := 0; i < _len; i++ {
		wg.Add(1)

		go func(v1 IVector, v2 IVector, out IVector, i int) {
			defer wg.Done()

			a := v1.GetAt(i)
			b := v2.GetAt(i)

			out.SetAt(i, math.Pow(a, b))

		}(v1, v2, out, i)
	}

	wg.Wait()

	return out
}

// Vector transformation delegating to function object.
// This method delegates vector transformation to function object set by users.
func (v *Vector) Map(f func(float64) float64) IVector {
	_len := v.Len()

	out := WithSize(_len)

	var wg sync.WaitGroup

	for i := 0; i < _len; i++ {
		wg.Add(1)

		go func(v1 IVector, v2 IVector, f func(float64) float64, i int) {
			defer wg.Done()

			out.SetAt(i, f(v.GetAt(i)))
		}(v, out, f, i)
	}

	wg.Wait()

	return out
}

// Vector algebra delegating to function object.
// This method delegates vector algebra to function object set by users, making
// it faster then these methods relying on reflection.
func Apply(v1 IVector, v2 IVector, f func(float64, float64) float64) IVector {
	_len := v1.Len()

	if !(_len == v2.Len()) {
		panic("Unequal vector size")
	}

	out := WithSize(_len)

	var wg sync.WaitGroup

	for i := 0; i < _len; i++ {
		wg.Add(1)

		go func(v1 IVector, v2 IVector, out IVector, f func(float64, float64) float64, i int) {
			defer wg.Done()

			out.SetAt(i, f(v1.GetAt(i), v2.GetAt(i)))
		}(v1, v2, out, f, i)
	}

	wg.Wait()

	return out
}
