package time

import (
	"github.com/pkg/errors"
	"strconv"
	"time"
)

// FormatTime ... Format time to string
func FormatTime(t time.Time, format string) string {
	return t.Format(format)
}

// ParseTime formats the input time string to time.Time object.
func ParseTime(inputTime string, format string) (time.Time, error) {
	t, err := time.ParseInLocation(format, inputTime, time.Local)
	if err != nil {
		return t, errors.Wrap(err, "ParseTime: not in format "+format)
	}
	return t, nil
}

func GetRequiredTimeFormat(inputTime, requestFormat, respFormat string) string {
	respTime, err := ParseTime(inputTime, requestFormat)
	if err != nil {
		return ""
	}
	return FormatTime(respTime, respFormat)
}

func AddDays(inputTime int, requestFormat string, addDays int) (date int) {
	requestTime := strconv.Itoa(inputTime)
	respTime, _ := ParseTime(requestTime, requestFormat)
	respTime = respTime.Add(time.Duration(addDays) * time.Hour * 24)
	date, _ = strconv.Atoi(FormatTime(respTime, requestFormat))
	return date
}

func SubstractDays(inputTime int, requestFormat string, substractDays int) (date int) {
	requestTime := strconv.Itoa(inputTime)
	respTime, _ := ParseTime(requestTime, requestFormat)
	respTime = respTime.Add(-1 * time.Duration(substractDays) * time.Hour * 24)
	date, _ = strconv.Atoi(FormatTime(respTime, requestFormat))
	return
}
