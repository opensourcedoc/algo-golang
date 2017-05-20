package list

import (
	"errors"
	"testing"
)

func TestMap(t *testing.T) {
	t.Parallel()

	list := newOddList()
	evenList, _ := list.Map(mapper)

	for e := range evenList.Iter() {
		if e.(int)%2 != 0 {
			t.Error("Error Map result")
		}
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := newOddList()
		_, _ = list.Map(mapper)
	}
}

func mapper(a interface{}) (interface{}, error) {
	_a, ok := a.(int)
	if ok != true {
		return nil, errors.New("Failed type assertion on a")
	}

	return _a * 2, nil
}
