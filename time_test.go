package bedinde

import (
	"testing"
	"time"
)

type tt struct {
	s string
	d TimeDiff
}

var td = []tt{
	{"5s", TimeDiff{0, 0, 0, 5}},
	{"63s", TimeDiff{0, 0, 1, 3}},
	{"2m7s", TimeDiff{0, 0, 2, 7}},
	{"2h5m3s", TimeDiff{0, 2, 5, 3}},
	{"64m3s", TimeDiff{0, 1, 4, 3}},
	{"51h4m5s", TimeDiff{2, 3, 4, 5}},
	{"48h5s", TimeDiff{2, 0, 0, 5}},
	{"48h5m", TimeDiff{2, 0, 5, 0}},
	{"50h75m", TimeDiff{2, 3, 15, 0}},
}

func TestFormatTimeDiff(t *testing.T) {
	for _, tt := range td {
		d, _ := time.ParseDuration(tt.s)
		r := FormatTimeDiff(d)

		if r != tt.d {
			t.Errorf("FormatTimeDiff is wrong formatting %v, got %v, want %v", d, r, tt.d)
		}
	}
}
