package types

import (
	"fmt"
	"os"
)

// Base struct for a ICalendar. When updated, it writes to the specified file path
type ICalendar struct {
	Coorperation string       // The author of the calendar
	ProductName  string       // The title of the product
	FilePath     string       // The file that the calendar should be updated to
	Events       []*ICalEvent // The events that the calendar bears
}

// Writes all events of an ICalendar to a specified file path. This function must be called everytime
// one wishes to update an ICalendar file with new events.
func (c *ICalendar) Update() error {
	// Open or create a file at the specified path
	f, err := os.OpenFile(c.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	// Initialises boilerplate
	lines := []string{
		"BEGIN:VCALENDAR\n",
		"VERSION:2.0\n",
		fmt.Sprintf("PRODID:-//%s/%s//NONSGML v1.0//EN\n", c.Coorperation, c.ProductName),
	}

	// Ranges over the events in the ICalendar, and creates ICal strings from them
	for _, event := range c.Events {
		eventString, err := event.ToString()
		if err != nil {
			return err
		}
		lines = append(lines, eventString)
	}

	// Writes the ICal strings to the file
	for _, line := range lines {
		f.WriteString(line)
	}

	// Ends Ical file with boilerplate
	f.WriteString("END:VCALENDAR")

	return nil
}
