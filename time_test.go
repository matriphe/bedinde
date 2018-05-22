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

func TestDatetimeStringToTime(t *testing.T) {
	tz := "Asia/Jakarta"
	s := "2018-05-22 23:05:47"
	l, _ := time.LoadLocation(tz)
	td := time.Date(2018, 5, 22, 23, 05, 47, 0, l)
	r, _ := DatetimeStringToTime(s, tz)

	if td.String() != r.String() {
		t.Errorf("DatetimeStringToTime is incorrect, got %v, want %v", r, td)
	}
}

func TestDateStringToTime(t *testing.T) {
	tz := "Asia/Jakarta"
	s := "2018-05-22"
	l, _ := time.LoadLocation(tz)
	td := time.Date(2018, 5, 22, 0, 0, 0, 0, l)
	r, _ := DateStringToTime(s, tz)

	if td.String() != r.String() {
		t.Errorf("DateStringToTime is incorrect, got %v, want %v", r, td)
	}
}
