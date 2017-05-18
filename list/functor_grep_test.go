package list

import (
	"errors"
	"testing"
)

func even(a interface{}) (bool, error) {
	_a, ok := a.(int)
	if ok != true {
		return false, errors.New("Failed type insertion on a")
	}

	if _a%2 == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func TestAnyTrue(t *testing.T) {
	list := newEvenList()

	if any, _ := list.Any(even); any != true {
		t.Error("Error Any even true")
	}
}

func TestAnyFalse(t *testing.T) {
	list := newOddList()

	if any, _ := list.Any(even); any != false {
		t.Error("Error Any even false")
	}
}

func TestAllTrue(t *testing.T) {
	list := newEvenList()

	if all, _ := list.All(even); all != true {
		t.Error("Error All even true")
	}
}

func TestAllFalse(t *testing.T) {
	list := newOddList()

	if all, _ := list.All(even); all != false {
		t.Error("Error All even false")
	}
}

func TestSelect(t *testing.T) {
	list := new1To10List()

	evenList, _ := list.Select(even)
	for e := range evenList.Iter() {
		if e.(int)%2 != 0 {
			t.Error("Error Select result")
		}
	}
}
