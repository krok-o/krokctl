package formatter

import (
	"strconv"

	"github.com/krok-o/krok/pkg/models"
)

// FormatEvent formats a event and displays it with the request
// format option.
func FormatEvent(event *models.Event, opt string) string {
	d := []kv{
		{"id", strconv.Itoa(event.ID)},
		{"event-id", event.EventID},
		{"repository-id", strconv.Itoa(event.RepositoryID)},
		{"vcs", strconv.Itoa(event.VCS)},
		{"payload", event.Payload},
		{"created-at", event.CreateAt.String()},
		//{"command-runs", event.CommandRuns},// TODO: Figure out how to display command runs.
	}
	formatter := NewFormatter(opt)
	return formatter.FormatObject(d)
}

// FormatEvents formats a list of events and displays it with the requested
// format option.
func FormatEvents(events []*models.Event, opt string) string {
	var d [][]kv
	for _, event := range events {
		data := []kv{
			{"id", strconv.Itoa(event.ID)},
			{"event-id", event.EventID},
			{"repository-id", strconv.Itoa(event.RepositoryID)},
			{"vcs", strconv.Itoa(event.VCS)},
			{"created-at", event.CreateAt.String()},
			// TODO: Add command run IDs so they can be looked up.
		}
		d = append(d, data)
	}
	formatter := NewFormatter(opt)
	return formatter.FormatList(d)
}
