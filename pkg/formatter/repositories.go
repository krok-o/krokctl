package formatter

import (
	"strconv"
	"strings"

	"github.com/krok-o/krok/pkg/models"
)

// FormatRepository formats a repository and displays it with the request
// format option.
func FormatRepository(repo *models.Repository, opt string) string {
	listOfCommandNames := make([]string, 0)
	for _, c := range repo.Commands {
		listOfCommandNames = append(listOfCommandNames, c.Name)
	}
	d := []kv{
		{"id", strconv.Itoa(repo.ID)},
		{"name", repo.Name},
		{"url", repo.URL},
		{"vcs", strconv.Itoa(repo.VCS)},
		{"callback-url", repo.UniqueURL},
		{"attached-commands", strings.Join(listOfCommandNames, ",")},
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
			{"id", strconv.Itoa(r.ID)},
			{"name", r.Name},
			{"url", r.URL},
			{"vcs", strconv.Itoa(r.VCS)},
		}
		d = append(d, data)
	}
	formatter := NewFormatter(opt)
	return formatter.FormatList(d)
}
