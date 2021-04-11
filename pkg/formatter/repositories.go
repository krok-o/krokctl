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
	return formatter.FormatObject(d)
}

// FormatRepositories formats a list of repositories and displays it with the requested
// format option.
func FormatRepositories(repos []*models.Repository, opt string) string {
	d := [][]kv{}
	for _, r := range repos {
		data := []kv{
			{"name", r.Name},
		}
		d = append(d, data)
	}
	formatter := NewFormatter(opt)
	return formatter.FormatList(d)
}
