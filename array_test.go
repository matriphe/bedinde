package bedinde

import "testing"

var (
	a        = []string{"matriphe", "muhammad", "zamroni"}
	b        = []uint64{1, 23, 436, 778, 8463, 3, 0, 123}
	k string = "matriphe"
	l string = "zam"
	m uint64 = 0
	n uint64 = 2
)

func TestInArrayFoundKeyInString(t *testing.T) {
	e, i := InArray(k, a)

	if !e {
		t.Errorf("InArray cannot find %s on array %v, got %t, want %t", k, a, e, true)
	}

	if i != 0 {
		t.Errorf("InArray cannot find %s on array %v, got %d, want %d", k, a, i, 0)
	}
}

func TestInArrayNotFoundKeyInString(t *testing.T) {
	e, i := InArray(l, a)

	if e {
		t.Errorf("InArray find %s on array %v, got %t, want %t", l, a, e, false)
	}

	if i >= 0 {
		t.Errorf("InArray find %s on array %v, got %d, want %d", l, a, i, -1)
	}
}

func TestInArrayFoundKeyInIntegers(t *testing.T) {
	e, i := InArray(m, b)

	if !e {
		t.Errorf("InArray cannot find %d on array %v, got %t, want %t", m, b, e, true)
	}

	if i != 6 {
		t.Errorf("InArray cannot find %d on array %v, got %d, want %d", m, b, i, 6)
	}
}

func TestInArrayNotFoundKeyInIntegers(t *testing.T) {
	e, i := InArray(n, b)

	if e {
		t.Errorf("InArray find %d on array %v, got %t, want %t", n, b, e, false)
	}

	if i >= 0 {
		t.Errorf("InArray find %d on array %v, got %d, want %d", n, b, i, -1)
	}
}
