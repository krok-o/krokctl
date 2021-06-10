package formatter

import (
	"strings"

	"github.com/krok-o/krok/pkg/models"
)

// FormatRepository formats a repository and displays it with the request
// format option.
func FormatRepository(repo *models.Repository, opt string) string {
	var listOfCommandNames []string
	for _, c := range repo.Commands {
		listOfCommandNames = append(listOfCommandNames, c.Name)
	}

	d := []kv{
		{"id", repo.ID},
		{"name", repo.Name},
		{"url", repo.URL},
		{"vcs", repo.VCS},
		{"callback-url", repo.UniqueURL},
		{"attached-commands", strings.Join(listOfCommandNames, ",")},
	}
	if repo.GitLab != nil {
		d = append(d, kv{
			"project-id", *repo.GitLab.GetProjectID(),
		})
	}
	formatter := NewFormatter(opt)
	return formatter.FormatObject(d)
}

// FormatRepositories formats a list of repositories and displays it with the requested
// format option.
func FormatRepositories(repos []*models.Repository, opt string) string {
	var d [][]kv
	for _, r := range repos {
		data := []kv{
			{"id", r.ID},
			{"name", r.Name},
			{"url", r.URL},
			{"vcs", r.VCS},
		}
		if r.GitLab != nil {
			data = append(data, kv{
				"project-id", *r.GitLab.GetProjectID(),
			})
		}
		d = append(d, data)
	}
	formatter := NewFormatter(opt)
	return formatter.FormatList(d)
}
