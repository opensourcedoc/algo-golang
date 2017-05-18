package list

import (
	"errors"
	"testing"
)

func cmp(a interface{}, b interface{}) (int, error) {
	na, ok := a.(int)
	if !ok {
		return 0, errors.New("Wrong type on a")
	}

	nb, ok := b.(int)
	if !ok {
		return 0, errors.New("Wrong type on b")
	}

	if na > nb {
		return 1, nil
	} else if na < nb {
		return -1, nil
	} else {
		return 0, nil
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

func TestClone(t *testing.T) {
	list := new1To10List()

	cloned := list.Clone()

	if b, _ := list.Equal(cloned, cmp); b != true {
		t.Error("Error Clone")
	}
}

func TestEqualFalse(t *testing.T) {
	list1 := newEvenList()
	list2 := newOddList()

	if b, _ := list1.Equal(list2, cmp); b != false {
		t.Error("Error Equal false")
	}
}

func TestFind(t *testing.T) {
	list := New()

	if _, err := list.Find(9999, cmp); err == nil {
		t.Error("No not found error")
	}

	list.Push(100)
	list.Push(200)
	list.Push(300)

	if index, _ := list.Find(100, cmp); index != 0 {
		t.Error("Error Find at 0")
	}

	if index, _ := list.Find(200, cmp); index != 1 {
		t.Error("Error Find at 1")
	}

	if index, _ := list.Find(300, cmp); index != 2 {
		t.Error("Error Find at 2")
	}

	if _, err := list.Find(9999, cmp); err == nil {
		t.Error("No not found error")
	}
}

func TestRemove(t *testing.T) {
	list := New()

	if err := list.Remove(9999, cmp); err == nil {
		t.Error("No index out of range error")
	}

	list.Push(100)
	list.Push(200)
	list.Push(300)
	list.Push(400)
	list.Push(500)

	if err := list.Remove(100, cmp); err != nil {
		t.Error("Error Remove 100")
	}

	if data, _ := list.At(0); data != 200 {
		t.Error("Error At 0 after Remove")
	}

	if err := list.Remove(500, cmp); err != nil {
		t.Error("Error Remove 500")
	}

	if data, _ := list.At(list.Len() - 1); data != 400 {
		t.Error("Error At Len() - 1 after Remove")
	}

	if err := list.Remove(300, cmp); err != nil {
		t.Error("Error Remove 300")
	}

	if err := list.Remove(9999, cmp); err == nil {
		t.Error("Error Remove 9999")
	}

	if list.Len() != 2 {
		t.Error("Error Len after Remove")
	}
}

func TestSort(t *testing.T) {
	list := New()

	list.Push(4)
	list.Push(2)
	list.Push(1)
	list.Push(3)

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
	list := New()

	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)

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
