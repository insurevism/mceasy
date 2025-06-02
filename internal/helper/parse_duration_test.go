package helper

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {

	criteria := []struct {
		name   string
		param  string
		result time.Duration
		err    error
	}{
		{
			name:   "1ms",
			param:  "1ms",
			result: time.Millisecond,
		},
		{
			name:   "millisecond",
			param:  "2 millisecond",
			result: time.Millisecond * 2,
		},
		{
			name:   "milliseconds",
			param:  "3 milliseconds",
			result: time.Millisecond * 3,
		},
		{
			name:   "s",
			param:  "1s",
			result: time.Second,
		},
		{
			name:   "second",
			param:  "2 second",
			result: time.Second * 2,
		},
		{
			name:   "seconds",
			param:  "3 seconds",
			result: time.Second * 3,
		},
		{
			name:   "m",
			param:  "1m",
			result: time.Minute,
		},
		{
			name:   "minute",
			param:  "2 minute",
			result: time.Minute * 2,
		},
		{
			name:   "minutes",
			param:  "3 minutes",
			result: time.Minute * 3,
		},
		{
			name:   "h",
			param:  "1h",
			result: time.Hour,
		},
		{
			name:   "hour",
			param:  "2 hour",
			result: time.Hour * 2,
		},
		{
			name:   "hours",
			param:  "3 hours",
			result: time.Hour * 3,
		},
		{
			name:   "d",
			param:  "1d",
			result: time.Hour * 24,
		},
		{
			name:   "day",
			param:  "2 day",
			result: time.Hour * 24 * 2,
		},
		{
			name:   "days",
			param:  "3 days",
			result: time.Hour * 24 * 3,
		},
		{
			name:   "float minute",
			param:  "0.5m",
			result: time.Second * 30,
		},
		{
			name:   "float hour",
			param:  "0.25 hours",
			result: time.Minute * 15,
		},
		{
			name:   "float day",
			param:  "0.25d",
			result: time.Hour * 6,
		},
		{
			name:   "negative ms",
			param:  "-1ms",
			result: -time.Millisecond,
		},
		{
			name:   "negative ms",
			param:  "-10ms",
			result: -time.Millisecond * 10,
		},
		{
			name:   "value parse error",
			param:  "a1ms",
			result: 0,
			err:    &strconv.NumError{},
		},
		{
			name:   "unit parse error",
			param:  "5month",
			result: 0,
			err:    errors.New("unit parse error"),
		},
	}

	for _, crit := range criteria {

		t.Run(crit.name, func(t *testing.T) {

			result, err := ParseDuration(crit.param)
			if crit.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, crit.result, result)
			} else {
				assert.Error(t, err)
				assert.IsType(t, crit.err, err)
			}
		})
	}

}
