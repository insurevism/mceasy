package helper

import (
	"fmt"
	"time"
)

const (
	// DateTimeLayout is a const to represent [time.DateTime] value.
	// This is for backward compatibility, we should use the std const once we update our Go version.
	DateTimeLayout = "2006-01-02 15:04:05"

	DateISO8601Layout = "2006-01-02T15:04:05Z"

	NetSuiteERPDateLayout = "02/01/2006"
	DateLayout     = "2006-01-02"
)

func ParseStringToTime(date string) (time.Time, error) {
	// Parse the date-time string into a time.Time object
	parsedTime, err := time.Parse(DateTimeLayout, date)
	if err != nil {
		fmt.Println("Error parsing date-time:", err)
		return parsedTime, err
	}

	// Print the parsed time
	fmt.Println("Parsed time:", parsedTime)
	return parsedTime, err
}

func TimeNowToString() string {
	currentTime := time.Now()
	return currentTime.Format("02-01-2006_15_04")
}

func FormatDateYYYYMMddHHmmss() string {
	return DateTimeLayout
}

func FormatDateYYYYMMdd() string {
	return DateLayout
}

func ParseStringToUnix(date string) int {
	if date == "" {
		return 0
	}

	asTime, err := time.Parse(DateTimeLayout, date)
	if err != nil {
		return 0
	}

	return int(asTime.Unix())
}

func ParseUnixToTime(unixTimestamp int) time.Time {
	if unixTimestamp == 0 {
		return time.Time{}
	}

	return time.Unix(int64(unixTimestamp), 0)
}

func MustParseStringToTime(date string) time.Time {
	parsedTime, err := time.Parse(DateTimeLayout, date)
	if err != nil {
		return time.Time{}
	}

	return parsedTime
}

func TimeToISO8601(requestTime time.Time) string {
	loc, _ := time.LoadLocation("Asia/Bangkok") // Zona waktu GMT+7
	now := requestTime.In(loc).Format(time.RFC3339)
	return now
}
