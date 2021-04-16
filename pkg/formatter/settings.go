package formatter

import (
	"strconv"

	"github.com/krok-o/krok/pkg/models"
)

// FormatSettings formats a list of settings and displays it with the request
// format option.
func FormatSettings(settings []*models.CommandSetting, opt string) string {
	var d [][]kv
	for _, setting := range settings {
		data := []kv{
			{"id", strconv.Itoa(setting.ID)},
			{"command-id", strconv.Itoa(setting.CommandID)},
			{"key", setting.Key},
			{"value", setting.Value},
			{"in-vault", strconv.FormatBool(setting.InVault)},
		}
		d = append(d, data)
	}
	formatter := NewFormatter(opt)
	return formatter.FormatList(d)
}
