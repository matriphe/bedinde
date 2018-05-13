package bedinde

import (
	"math"
	"time"
)

type TimeDiff struct {
	Day    uint64
	Hour   uint64
	Minute uint64
	Second uint64
}

func FormatTimeDiff(diff time.Duration) TimeDiff {
	res := TimeDiff{
		Day:    0,
		Hour:   0,
		Minute: 0,
		Second: 0,
	}

	seconds := uint64(diff.Seconds())
	if seconds < 60 {
		res.Second = seconds
		return res
	}

	minutes := uint64(diff.Minutes())
	seconds = uint64(math.Mod(float64(seconds), 60))

	if minutes < 60 {
		res.Second = seconds
		res.Minute = minutes
		return res
	}

	hours := uint64(diff.Hours())
	minutes = uint64(math.Mod(float64(minutes), 60))

	if hours < 24 {
		res.Second = seconds
		res.Minute = minutes
		res.Hour = hours
		return res
	}

	days := uint64(diff.Hours() / 24)
	hours = uint64(math.Mod(float64(hours), 24))

	res.Day = days
	res.Second = seconds
	res.Minute = minutes
	res.Hour = hours
	return res
}
