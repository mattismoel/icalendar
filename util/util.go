package util

import (
	"fmt"
	"strconv"
	"time"
)

// Returns a ICalTimestamp string by an input date time
func TimeToICalTimestamp(t *time.Time) (string, error) {
	year := PadInt(t.Year(), 2)
	month := PadInt(int(t.Month()), 2)
	day := PadInt(t.Day(), 2)
	hour := PadInt(t.Hour(), 2)
	minute := PadInt(t.Minute(), 2)

	str := fmt.Sprintf("%s%s%sT%s%s00Z", year, month, day, hour, minute)
	return str, nil
}

// Returns a string representation of an integer with given amount of padding zeroes
func PadInt(i int, count int) string {
	layout := fmt.Sprintf("%%0%dd", count)
	return fmt.Sprintf(layout, i)
}

// Returns a date time given an input ICalendar timestamp
func ICalTimestampToTime(stamp string) (*time.Time, error) {
	year, err := strconv.Atoi(stamp[:4])
	if err != nil {
		return &time.Time{}, err
	}
	month, err := strconv.Atoi(stamp[4:6])
	if err != nil {
		return &time.Time{}, err
	}
	day, err := strconv.Atoi(stamp[6:8])
	if err != nil {
		return &time.Time{}, err
	}
	hour, err := strconv.Atoi(stamp[9:11])
	if err != nil {
		return &time.Time{}, err
	}
	minute, err := strconv.Atoi(stamp[11:13])
	if err != nil {
		return &time.Time{}, err
	}
	second, err := strconv.Atoi(stamp[13:15])
	if err != nil {
		return &time.Time{}, err
	}

	location, err := time.LoadLocation("Europe/Copenhagen")
	date := time.Date(year, time.Month(month), day, hour, minute, second, 0, location)
	return &date, err
}
