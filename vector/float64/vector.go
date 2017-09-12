package vector

import (
	"math"
)

type IVector interface {
	Len() int
	GetAt(int) float64
	SetAt(int, float64)
}

type Vector struct {
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

	v.vec[i] = data
}

func Magnitude(v IVector) float64 {
	temp := Map(v, func(n float64) float64 { return math.Pow(n, 2) })
	sum := Reduce(temp, func(a float64, b float64) float64 { return a + b })
	return math.Sqrt(sum)
}

func Sort(v IVector) IVector {
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
	return Apply(v1, v2, func(a float64, b float64) float64 { return a + b })
}

func ScalarAdd(v IVector, s float64) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return a + b })
}

// Vector substraction
func Sub(v1 IVector, v2 IVector) IVector {
	return Apply(v1, v2, func(a float64, b float64) float64 { return a - b })
}

func ScalarSubFirst(s float64, v IVector) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return b - a })
}

func ScalarSubSecond(v IVector, s float64) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return a - b })
}

// Vector multiplication
func Mul(v1 IVector, v2 IVector) IVector {
	return Apply(v1, v2, func(a float64, b float64) float64 { return a * b })
}

func ScalarMul(v IVector, s float64) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return a * b })
}

// Vector division
func Div(v1 IVector, v2 IVector) IVector {
	return Apply(v1, v2, func(a float64, b float64) float64 { return a / b })
}

func ScalarDivFirst(s float64, v IVector) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return b / a })
}

func ScalarDivSecond(v IVector, s float64) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return a / b })
}

// Vector power
func Pow(v1 IVector, v2 IVector) IVector {
	return Apply(v1, v2, func(a float64, b float64) float64 { return math.Pow(a, b) })
}

func ScalarPowBase(s float64, v IVector) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return math.Pow(b, a) })
}

func ScalarPowExp(v IVector, s float64) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return math.Pow(a, b) })
}

// Vector product
func Dot(v1 IVector, v2 IVector) float64 {
	return Reduce(Mul(v1, v2), func(a float64, b float64) float64 { return a + b })
}

func Cos(v1 IVector, v2 IVector) float64 {
	return Dot(v1, v2) / (Magnitude(v1) * Magnitude(v2))
}

// Vector transformation delegating to function object.
// This method delegates vector transformation to function object set by users.
func Map(v IVector, f func(float64) float64) IVector {
	_len := v.Len()

	out := WithSize(_len)

	for i := 0; i < _len; i++ {
		out.SetAt(i, f(v.GetAt(i)))
	}

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

	for i := 0; i < _len; i++ {
		out.SetAt(i, f(v1.GetAt(i), v2.GetAt(i)))
	}

	return out
}

func ScalarApply(v IVector, s float64, f func(float64, float64) float64) IVector {
	_len := v.Len()

	out := WithSize(_len)

	for i := 0; i < _len; i++ {
		out.SetAt(i, f(v.GetAt(i), s))
	}

	return out
}

// Vector to scalars reduction.
// This method delegates vector reduction to function object set by users.
func Reduce(v IVector, f func(float64, float64) float64) float64 {
	_len := v.Len()

	if _len <= 1 {
		return v.GetAt(0)
	}

	out := v.GetAt(0)

	for i := 1; i < _len; i++ {
		out = f(out, v.GetAt(i))
	}

	return out
}
