package list

import (
	"errors"
	"testing"
)

func TestClone(t *testing.T) {
	list := new1To10List()

	cloned := list.Clone()

	if b, _ := list.Equal(cloned, isEqual); b != true {
		t.Error("Error Clone")
	}
}

func TestEqualFalse(t *testing.T) {
	list1 := newEvenList()
	list2 := newOddList()

	if b, _ := list1.Equal(list2, isEqual); b != false {
		t.Error("Error Equal false")
	}
}

func TestFind(t *testing.T) {
	list := FromArgs(100, 200, 300)

	if index, _ := list.Find(100, isEqual); index != 0 {
		t.Error("Error Find at 0")
	}

	if index, _ := list.Find(200, isEqual); index != 1 {
		t.Error("Error Find at 1")
	}

	if index, _ := list.Find(300, isEqual); index != 2 {
		t.Error("Error Find at 2")
	}

	if _, err := list.Find(9999, isEqual); err == nil {
		t.Error("No not found error")
	}
}

func TestRemove(t *testing.T) {
	list := FromArgs(100, 200, 300, 400, 500)

	if err := list.Remove(100, isEqual); err != nil {
		t.Error("Error Remove 100")
	}

	if data, _ := list.At(0); data != 200 {
		t.Error("Error At 0 after Remove")
	}

	if err := list.Remove(500, isEqual); err != nil {
		t.Error("Error Remove 500")
	}

	if data, _ := list.At(list.Len() - 1); data != 400 {
		t.Error("Error At Len() - 1 after Remove")
	}

	if err := list.Remove(300, isEqual); err != nil {
		t.Error("Error Remove 300")
	}

	if err := list.Remove(9999, isEqual); err == nil {
		t.Error("Error Remove 9999")
	}

	if list.Len() != 2 {
		t.Error("Error Len after Remove")
	}
}

func isEqual(a interface{}, b interface{}) (bool, error) {
	na, ok := a.(int)
	if !ok {
		return false, errors.New("Wrong type on a")
	}

	nb, ok := b.(int)
	if !ok {
		return false, errors.New("Wrong type on b")
	}

	return na == nb, nil
}

func TestSort(t *testing.T) {
	list := FromArgs(4, 2, 1, 3)

	// (1, 2, 3, 4)
	sorted, _ := list.Sort(isSmaller)

	count := 1
	for e := range sorted.Iter() {
		if e != count {
			t.Error("Error Sort")
		}

		count++
	}

	if sorted.Len() != 4 {
		t.Error("Error Len after Sort")
	}
}

func TestSortSorted(t *testing.T) {
	list := FromArgs(1, 2, 3, 4)

	// (1, 2, 3, 4)
	sorted, _ := list.Sort(isSmaller)

	count := 1
	for e := range sorted.Iter() {
		if e != count {
			t.Error("Error Sort")
		}

		count++
	}

	if sorted.Len() != 4 {
		t.Error("Error Len after Sort")
	}
}

func isSmaller(a interface{}, b interface{}) (bool, error) {
	na, ok := a.(int)
	if !ok {
		return false, errors.New("Wrong type on a")
	}

	nb, ok := b.(int)
	if !ok {
		return false, errors.New("Wrong type on b")
	}

	return na < nb, nil
}
