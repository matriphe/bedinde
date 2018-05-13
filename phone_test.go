package bedinde

import "testing"

type p struct {
	s string
	r string
}

var (
	id = []p{
		{"", ""},
		{"081823456789", "+6281823456789"},
		{"0818 2345 6789", "+6281823456789"},
		{"0818-2345-6789", "+6281823456789"},
	}
	sg = []p{
		{"", ""},
		{"081823456789", "+65081823456789"},
		{"0818 2345 6789", "+65081823456789"},
		{"0818-2345-6789", "+65081823456789"},
	}
)

func TestFormatPhoneIndonesia(t *testing.T) {
	for _, p := range id {
		r := FormatPhone(p.s, "id")

		if r != p.r {
			t.Errorf("FormatPhone is wrong formatting %s, got %v, want %v", p.s, r, p.r)
		}
	}
}

func TestFormatPhoneSingapore(t *testing.T) {
	for _, p := range sg {
		r := FormatPhone(p.s, "sg")

		if r != p.r {
			t.Errorf("FormatPhone is wrong formatting %s, got %v, want %v", p.s, r, p.r)
		}
	}
}
