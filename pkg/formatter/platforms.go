package formatter

import (
	"github.com/krok-o/krok/pkg/models"
)

// FormatPlatforms formats a list of platforms and displays it with the requested
// format option.
func FormatPlatforms(platforms []models.Platform, opt string) string {
	var d [][]kv
	for _, platform := range platforms {
		data := []kv{
			{"id", platform.ID},
			{"name", platform.Name},
		}
		d = append(d, data)
	}
	formatter := NewFormatter(opt)
	return formatter.FormatList(d)
}
