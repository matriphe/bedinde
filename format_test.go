package bedinde

import "testing"

type fn struct {
	n uint64
	s string
}

var (
	dot = []fn{
		{1, "1"},
		{12, "12"},
		{124, "124"},
		{1256, "1.256"},
		{12569, "12.569"},
		{125698, "125.698"},
		{1256982, "1.256.982"},
	}
	comma = []fn{
		{1, "1"},
		{12, "12"},
		{124, "124"},
		{1256, "1,256"},
		{12569, "12,569"},
		{125698, "125,698"},
		{1256982, "1,256,982"},
	}
)

func TestFormatNumberIndonesian(t *testing.T) {
	for _, n := range dot {
		f := FormatNumber(n.n, "id")

		if f != n.s {
			t.Errorf("FormatNumber is wrong formatting %d, got %s, want %s", n.n, f, n.s)
		}
	}
}

func TestFormatNumberEnglish(t *testing.T) {
	for _, n := range comma {
		f := FormatNumber(n.n, "en")

		if f != n.s {
			t.Errorf("FormatNumber is wrong formatting %d, got %s, want %s", n.n, f, n.s)
		}
	}
}

func TestFormatNumberGerman(t *testing.T) {
	for _, n := range dot {
		f := FormatNumber(n.n, "de")

		if f != n.s {
			t.Errorf("FormatNumber is wrong formatting %d, got %s, want %s", n.n, f, n.s)
		}
	}
}

func TestFormatNumberSingaporean(t *testing.T) {
	for _, n := range dot {
		f := FormatNumber(n.n, "sg")

		if f != n.s {
			t.Errorf("FormatNumber is wrong formatting %d, got %s, want %s", n.n, f, n.s)
		}
	}
}
