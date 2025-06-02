package helper_test

import (
	"mceasy/internal/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStringToTime(t *testing.T) {
	t.Parallel()

	input := "2006-01-02 15:04:05"
	res, err := helper.ParseStringToTime(input)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestTimeNowToString(t *testing.T) {
	t.Parallel()

	assert.NotEmpty(t, helper.TimeNowToString())
}

func TestFormatDateYYYYMMddHHmmss(t *testing.T) {
	t.Parallel()

	assert.Equal(t, helper.DateTimeLayout, helper.FormatDateYYYYMMddHHmmss())
}

func TestParseStringToUnix(t *testing.T) {
	t.Parallel()

	t.Run("empty date string always return 0", func(t *testing.T) {
		assert.Equal(t, 0, helper.ParseStringToUnix(""))
	})

	t.Run("invalid date string always return 0", func(t *testing.T) {
		assert.Equal(t, 0, helper.ParseStringToUnix("2023"))
	})

	t.Run("successfully parse to unix", func(t *testing.T) {
		assert.True(t, helper.ParseStringToUnix("2006-01-02 15:04:05") > 0)

	})
}

func TestParseUnixToTime(t *testing.T) {
	t.Parallel()

	t.Run("empty unix timestamp always return zero time", func(t *testing.T) {
		assert.True(t, helper.ParseUnixToTime(0).IsZero())
	})

	t.Run("successfully parse unix timestamp", func(t *testing.T) {
		assert.True(t, !helper.ParseUnixToTime(1725190946).IsZero())
	})
}

func TestMustParseStringToTime(t *testing.T) {
	t.Parallel()

	t.Run("fail to parse will return zero time", func(t *testing.T) {
		assert.True(t, helper.MustParseStringToTime("2023").IsZero())
	})

	t.Run("successfully parse time", func(t *testing.T) {
		assert.True(t, !helper.MustParseStringToTime("2006-01-02 15:04:05").IsZero())
	})
}
