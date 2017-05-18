package list

import (
	"testing"
)

func TestPushPop(t *testing.T) {
	list := New()

	if list.Len() != 0 {
		t.Error("Error Len on empty list")
	}

	list.Push(100)
	list.Push(200)
	list.Push(300)

	if list.Len() != 3 {
		t.Error("Error Len after Push")
	}

	data, _ := list.Pop()
	if data != 300 {
		t.Error("Error popped data")
	}

	if list.Len() != 2 {
		t.Error("Error Len after Pop")
	}

	list.Pop()
	list.Pop()
	if _, err := list.Pop(); err == nil {
		t.Error("No empty list error")
	}
}

func TestUnshiftShift(t *testing.T) {
	list := New()

	list.Unshift(100)
	list.Unshift(200)
	list.Unshift(300)

	if list.Len() != 3 {
		t.Error("Error Len after Unshift")
	}

	data, _ := list.Shift()
	if data != 300 {
		t.Error("Error shifted data")
	}

	list.Shift()
	list.Shift()

	if list.Len() != 0 {
		t.Error("Error Len after Shift")
	}

	if _, err := list.Shift(); err == nil {
		t.Error("No empty list error")
	}
}

func TestEach(t *testing.T) {
	list := New()

	list.Push(100)
	list.Push(200)
	list.Push(300)
	list.Push(400)
	list.Push(500)

	for e := range list.Each() {
		if e.(int)%100 != 0 {
			t.Error("Error Each")
		}
	}
}

func TestRemoveAt(t *testing.T) {
	list := New()

	list.Push(100)
	list.Push(200)
	list.Push(300)
	list.Push(400)
	list.Push(500)

	if data, _ := list.RemoveAt(0); data != 100 {
		t.Error("Error RemoveAt 0")
	}

	if data, _ := list.RemoveAt(list.Len() - 1); data != 500 {
		t.Error("Error Remove At Len() - 1")
	}

	if data, _ := list.RemoveAt(1); data != 300 {
		t.Error("Error RemoveAt 1")
	}

	if list.Len() != 2 {
		t.Error("Error Len after RemoveAt")
	}
}

func TestAt(t *testing.T) {
	list := New()

	if _, err := list.At(0); err == nil {
		t.Error("No index out of range")
	}

	list.Push(100)
	list.Push(200)
	list.Push(300)

	if data, _ := list.At(0); data.(int) != 100 {
		t.Error("Error at 0")
	}

	if data, _ := list.At(1); data.(int) != 200 {
		t.Error("Error at 1")
	}

	if data, _ := list.At(2); data.(int) != 300 {
		t.Error("Error at 2")
	}

	if _, err := list.At(99); err == nil {
		t.Error("No index out of range error")
	}
}
