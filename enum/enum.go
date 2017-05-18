package enum

type Enum struct {
	set map[interface{}]bool
}

func New(args ...interface{}) *Enum {
	enum := new(Enum)
	enum.set = make(map[interface{}]bool)

	for _, e := range args {
		enum.set[e] = true
	}

	return enum
}

func NewFromArray(arr []interface{}) *Enum {
	enum := new(Enum)
	enum.set = make(map[interface{}]bool)

	for _, e := range arr {
		enum.set[e] = true
	}

	return enum
}

func (e *Enum) Has(v interface{}) bool {
	_, ok := e.set[v]

	if !ok {
		panic("Unknown value")
	}

	return true
}

func (e *Enum) Get(v interface{}) interface{} {
	// Check v exists; otherwise, emit panic event
	e.Has(v)

	return v
}

func (e *Enum) Eq(a interface{}, b interface{}) bool {
	// Check a and b exist; otherwise, emit panic event
	e.Has(a)
	e.Has(b)

	if a == b {
		return true
	}

	return false
}
