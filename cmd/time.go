package cmd

import (
	"time"
)

// ParseTime will parse the string representation of time.
func ParseTime(date string) (time.Time, error) {
	return time.Parse(time.RFC3339, date)
}
