package types

import (
	"fmt"
	"time"

	"github.com/mattismoel/icalendar/util"
)

// Base struct for an ICalendar event
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

	startDate, err := util.TimeToICalTimestamp(&e.StartDate)
	endDate, err := util.TimeToICalTimestamp(&e.EndDate)
	if err != nil {
		return "", err
	}
	eventStr := fmt.Sprintf(layout, e.UID, startDate, startDate, endDate, e.Summary, e.Location)
	return eventStr, nil
}
