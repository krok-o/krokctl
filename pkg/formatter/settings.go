package formatter

import (
	"strconv"

	"github.com/krok-o/krok/pkg/models"
)

// FormatSetting formats a setting and displays it with the request
// format option.
func FormatSetting(setting *models.CommandSetting, opt string) string {
	d := []kv{
		{"id", setting.ID},
		{"command-id", setting.CommandID},
		{"key", setting.Key},
		{"value", setting.Value},
		{"in-vault", strconv.FormatBool(setting.InVault)},
	}
	formatter := NewFormatter(opt)
	return formatter.FormatObject(d)
}

// FormatSettings formats a list of settings and displays it with the request
// format option.
func FormatSettings(settings []*models.CommandSetting, opt string) string {
	var d [][]kv
	for _, setting := range settings {
		data := []kv{
			{"id", setting.ID},
			{"command-id", setting.CommandID},
			{"key", setting.Key},
			{"value", setting.Value},
			{"in-vault", strconv.FormatBool(setting.InVault)},
		}
		d = append(d, data)
	}
	formatter := NewFormatter(opt)
	return formatter.FormatList(d)
}
