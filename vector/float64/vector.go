package vector

import (
	"math"
)

// IVector is the interface of Vector class.
type IVector interface {
	Len() int
	GetAt(int) float64
	SetAt(int, float64)
}

// Vector class implements IVector interface.
// Internally, it stores data into a float64 slice.
type Vector struct {
	vec []float64
}

// New creates a vector based on arguments.
func New(args ...float64) IVector {
	v := new(Vector)
	v.vec = make([]float64, len(args))

	for i, e := range args {
		v.SetAt(i, e)

	}

	return v
}

// WithSize creates a zero-based vector with specific size.
func WithSize(size int) IVector {
	v := new(Vector)
	v.vec = make([]float64, size)

	for i := 0; i < size; i++ {
		v.SetAt(i, 0.0)
	}

	return v
}

// FromArray creates a vector from float64 slice.
func FromArray(arr []float64) IVector {
	v := new(Vector)
	v.vec = make([]float64, len(arr))

	for i, e := range arr {
		v.SetAt(i, e)
	}

	return v
}

// Len returns the length of the vector. Compare it with Magnitude.
func (v *Vector) Len() int {
	return len(v.vec)
}

// GetAt returns the value of the vector at index i.
func (v *Vector) GetAt(i int) float64 {
	if i < 0 || i >= v.Len() {
		panic("Index out of range")
	}

	return v.vec[i]
}

// SetAt sets the value of the vector at index i.
func (v *Vector) SetAt(i int, data float64) {
	if i < 0 || i >= v.Len() {
		panic("Index out of range")
	}

	v.vec[i] = data
}

// Magnitude returns the magnitude of a vector.
func Magnitude(v IVector) float64 {
	temp := Map(v, func(n float64) float64 { return math.Pow(n, 2) })
	sum := Reduce(temp, func(a float64, b float64) float64 { return a + b })
	return math.Sqrt(sum)
}

// Sort sorts a vector.
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

// Add performs element-wise addition between two vectors.
func Add(v1 IVector, v2 IVector) IVector {
	return Apply(v1, v2, func(a float64, b float64) float64 { return a + b })
}

// ScalarAdd adds a vector and a scalar.
func ScalarAdd(v IVector, s float64) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return a + b })
}

// Sub performs element-wise substraction between two vectors.
func Sub(v1 IVector, v2 IVector) IVector {
	return Apply(v1, v2, func(a float64, b float64) float64 { return a - b })
}

// ScalarSubFirst substracts a scalar to a vector.
func ScalarSubFirst(s float64, v IVector) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return b - a })
}

// ScalarSubSecond substracts a vector to a scalar.
func ScalarSubSecond(v IVector, s float64) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return a - b })
}

// Mul performs element-wise multiplication between two vectors.
func Mul(v1 IVector, v2 IVector) IVector {
	return Apply(v1, v2, func(a float64, b float64) float64 { return a * b })
}

// ScalarMul performs element-wise multiplication between a vector and a scalar.
func ScalarMul(v IVector, s float64) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return a * b })
}

// Div perfomrms element-wise division.
func Div(v1 IVector, v2 IVector) IVector {
	return Apply(v1, v2, func(a float64, b float64) float64 { return a / b })
}

// ScalarDivFirst divides a scalar numerator and a vector denominator.
func ScalarDivFirst(s float64, v IVector) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return b / a })
}

// ScalarDivSecond divides a vector numerator and a scalar denominator.
func ScalarDivSecond(v IVector, s float64) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return a / b })
}

// Pow performs element-wise power.
func Pow(v1 IVector, v2 IVector) IVector {
	return Apply(v1, v2, func(a float64, b float64) float64 { return math.Pow(a, b) })
}

// ScalarPowBase return the power of a scalar base against a vector exponent.
func ScalarPowBase(s float64, v IVector) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return math.Pow(b, a) })
}

// ScalarPowExp return the power of a vector base against a scalar exponent.
func ScalarPowExp(v IVector, s float64) IVector {
	return ScalarApply(v, s, func(a float64, b float64) float64 { return math.Pow(a, b) })
}

// Dot returns the dot product between two vectors.
func Dot(v1 IVector, v2 IVector) float64 {
	return Reduce(Mul(v1, v2), func(a float64, b float64) float64 { return a + b })
}

// Cos returns the cosine value between two vectors.
func Cos(v1 IVector, v2 IVector) float64 {
	return Dot(v1, v2) / (Magnitude(v1) * Magnitude(v2))
}

// Map transform a vector into another vector by an user-defined function.
func Map(v IVector, f func(float64) float64) IVector {
	_len := v.Len()

	out := WithSize(_len)

	for i := 0; i < _len; i++ {
		out.SetAt(i, f(v.GetAt(i)))
	}

	return out
}

// Apply applies an user-defined function to two vectors.
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

// ScalarApply applies an user-defined function to a vector and a scalar.
func ScalarApply(v IVector, s float64, f func(float64, float64) float64) IVector {
	_len := v.Len()

	out := WithSize(_len)

	for i := 0; i < _len; i++ {
		out.SetAt(i, f(v.GetAt(i), s))
	}

	return out
}

// Reduce tranforms a vector to a scalar by an user-defined function.
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
