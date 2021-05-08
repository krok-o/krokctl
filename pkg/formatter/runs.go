package formatter

import (
	"github.com/krok-o/krok/pkg/models"
)

// FormatCommandRun formats a command run and displays it with the request
// format option.
func FormatCommandRun(run *models.CommandRun, opt string) string {
	d := []kv{
		{"id", run.ID},
		{"event-id", run.EventID},
		{"name", run.CommandName},
		{"status", run.Status},
		{"created-at", run.CreateAt.String()},
		{"outcome", run.Outcome},
	}
	formatter := NewFormatter(opt)
	return formatter.FormatObject(d)
}
