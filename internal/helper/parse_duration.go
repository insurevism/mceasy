package helper

import (
	"fmt"
	"github.com/gookit/goutil/arrutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var unitMap = []struct {
	unit     []string
	duration time.Duration
}{
	{
		unit:     []string{"ns", "nanosecond", "nanoseconds"},
		duration: time.Nanosecond,
	},
	{
		unit: []string{
			"us",
			"µs", // U+00B5
			"μs", // U+03BC
			"microsecond",
			"microseconds",
		},
		duration: time.Microsecond,
	},
	{
		unit:     []string{"ms", "millisecond", "milliseconds"},
		duration: time.Millisecond,
	},
	{
		unit:     []string{"s", "second", "seconds"},
		duration: time.Second,
	},
	{
		unit:     []string{"m", "minute", "minutes"},
		duration: time.Minute,
	},
	{
		unit:     []string{"h", "hour", "hours"},
		duration: time.Hour,
	},
	{
		unit:     []string{"d", "day", "days"},
		duration: time.Hour * 24,
	},
}

var parseDurationRegexp = regexp.MustCompile("[a-zA-Zµμ]+$")

// ParseDuration parses a duration string.
// A duration string is a possibly signed sequence of
// decimal numbers, each with optional fraction and a unit suffix,
// examples "300ms", "-1.5h" or "2h45m".
// Supported time units are :
// - ms, millisecond, milliseconds
// - s, second, seconds
// - m, minute, minutes
// - h, hour, hours
// - d, day, days
func ParseDuration(duration string) (time.Duration, error) {

	duration = strings.TrimSpace(duration)

	result, err := time.ParseDuration(duration)
	if err == nil {
		return result, err
	}

	unit := parseDurationRegexp.Find([]byte(duration))
	head := strings.TrimSpace(strings.Replace(duration, string(unit), "", 1))

	value, err := strconv.ParseFloat(head, 64)
	if err != nil {
		return 0, err
	}

	for _, table := range unitMap {

		if arrutil.In(strings.ToLower(string(unit)), table.unit) {

			return time.Duration(float64(table.duration) * value), nil
		}
	}

	return 0, fmt.Errorf("invalid unit duration: %s", unit)

}
