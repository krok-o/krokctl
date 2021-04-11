package formatter

import (
	"github.com/krok-o/krok/pkg/models"
)

// FormatRepository formats a repository and displays it with the request
// format option.
func FormatRepository(repo *models.Repository, opt string) string {
	d := []kv{
		{"name", repo.Name},
		{"url", repo.URL},
	}
	formatter := NewFormatter(opt)
	return formatter.Format(d...)
}
