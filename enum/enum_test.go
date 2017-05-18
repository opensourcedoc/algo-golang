package enum

import "testing"

func TestExistingValue(t *testing.T) {
	t.Parallel()

	enum := New("Red", "Green", "Blue")

	if !enum.Has("Green") {
		t.Error("Green should be there")
	}
}

func TestNonExistingValue(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	enum := New("Red", "Green", "Blue")
	if enum.Has("Yellow") {
		t.Error("Yellow should not be there")
	}
}

func TestGetValue(t *testing.T) {
	t.Parallel()

	enum := New("Red", "Green", "Blue")
	green := enum.Get("Green")

	g, ok := green.(string)
	if !ok {
		t.Error("Green should be a string")
	}

	if !(g == "Green") {
		t.Error("Green should be there")
	}
}

func TestEquality(t *testing.T) {
	t.Parallel()

	enum := New("Red", "Green", "Blue")
	green := enum.Get("Green")

	if !(enum.Eq(green, "Green")) {
		t.Error("It should be equal")
	}

	if enum.Eq(green, "Blue") {
		t.Error("It should not be equal")
	}
}

func TestGetNonExistingValue(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	enum := New("Red", "Green", "Blue")
	_ = enum.Get("Yellow")
}

func TestExistingValueFromArray(t *testing.T) {
	t.Parallel()

	enum := NewFromArray([]interface{}{"Red", "Green", "Blue"})

	if !enum.Has("Green") {
		t.Error("Green should be there")
	}
}

func TestNonExistingValueFromArray(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	enum := NewFromArray([]interface{}{"Red", "Green", "Blue"})
	if enum.Has("Yellow") {
		t.Error("Yellow should not be there")
	}
}

func TestGetValueFromArray(t *testing.T) {
	t.Parallel()

	enum := NewFromArray([]interface{}{"Red", "Green", "Blue"})
	green := enum.Get("Green")

	g, ok := green.(string)
	if !ok {
		t.Error("Green should be a string")
	}

	if !(g == "Green") {
		t.Error("Green should be there")
	}
}

func TestGetNonExistingValueFromArray(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	enum := NewFromArray([]interface{}{"Red", "Green", "Blue"})
	_ = enum.Get("Yellow")
}
