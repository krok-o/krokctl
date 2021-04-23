package formatter

import (
	"fmt"
	"strings"

	"github.com/krok-o/krok/pkg/models"
)

// FormatEvent formats a event and displays it with the request
// format option.
func FormatEvent(event *models.Event, opt string) string {
	var commandRuns []string
	for _, c := range event.CommandRuns {
		commandRuns = append(commandRuns, fmt.Sprintf("%d; %s; %s", c.ID, c.CommandName, c.Status))
	}
	d := []kv{
		{"id", event.ID},
		{"event-id", event.EventID},
		{"repository-id", event.RepositoryID},
		{"vcs", event.VCS},
		{"payload", event.Payload},
		{"created-at", event.CreateAt.String()},
		{"command-runs", strings.Join(commandRuns, "|")},
	}
	formatter := NewFormatter(opt)
	return formatter.FormatObject(d)
}

// FormatEvents formats a list of events and displays it with the requested
// format option.
func FormatEvents(events []*models.Event, opt string) string {
	var d [][]kv
	for _, event := range events {
		var commandRuns []string
		for _, c := range event.CommandRuns {
			commandRuns = append(commandRuns, fmt.Sprintf("%d; %s; %s", c.ID, c.CommandName, c.Status))
		}
		data := []kv{
			{"id", event.ID},
			{"event-id", event.EventID},
			{"repository-id", event.RepositoryID},
			{"vcs", event.VCS},
			{"created-at", event.CreateAt.String()},
			{"command-runs", strings.Join(commandRuns, "|")},
		}
		d = append(d, data)
	}
	formatter := NewFormatter(opt)
	return formatter.FormatList(d)
}
