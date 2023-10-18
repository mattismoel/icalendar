package types

import (
	"fmt"
	"strconv"
	"time"
)

type ICalendar struct {
	Coorperation string       // The author of the calendar
	ProductName  string       // The title of the product
	FilePath     string       // The file that the calendar should be updated to
	Events       []*ICalEvent // The events that the calendar bears
}

type ICalEvent struct {
	UID         string    // The unique ID of the calendar event
	StartDate   time.Time // The start of the event - Date time
	EndDate     time.Time // The end of the event - Date time
	Summary     string    // The title of the event
	Location    string    // The location of the event
	Description string    // The description of the event
}

// Returns the string representation of an ICalendar struct. This can be input into an .ics file.
func (e *ICalEvent) ToString() (string, error) {
	layout := `BEGIN:VEVENT
UID:%s
DTSTAMP:%s
DTSTART:%s
DTEND:%s
SUMMARY:%s 
LOCATION:%s
END:VEVENT
`

	startDate, err := TimeToICalTimestamp(&e.StartDate)
	endDate, err := TimeToICalTimestamp(&e.EndDate)
	if err != nil {
		return "", err
	}
	eventStr := fmt.Sprintf(layout, e.UID, startDate, startDate, endDate, e.Summary, e.Location)
	return eventStr, nil
}

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
	fmt.Println(layout)
	return fmt.Sprintf(layout, i)
}

func ICalTimestampToTime(stamp string) (*time.Time, error) {
	// "0 1 2 3 | 45 | 67 | 8 | 9 10 | 11 12 | 13 14 | 15"
	// "1 9 9 7 | 07 | 15 | T | 0 4  | 0  0  | 0  0  | Z"

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
