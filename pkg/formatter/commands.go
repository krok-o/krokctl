package formatter

import (
	"strconv"
	"strings"

	"github.com/krok-o/krok/pkg/models"
)

// FormatCommand formats a command and displays it with the request
// format option.
func FormatCommand(command *models.Command, opt string) string {
	var listOfRepoNames []string
	for _, r := range command.Repositories {
		listOfRepoNames = append(listOfRepoNames, r.Name)
	}
	var listOfPlatforms []string
	for _, r := range command.Platforms {
		listOfPlatforms = append(listOfPlatforms, r.Name)
	}
	d := []kv{
		{"id", command.ID},
		{"name", command.Name},
		{"schedule", command.Schedule},
		{"image", command.Image},
		{"enabled", strconv.FormatBool(command.Enabled)},
		{"repositories", strings.Join(listOfRepoNames, ",")},
		{"platforms", strings.Join(listOfPlatforms, ",")},
	}
	formatter := NewFormatter(opt)
	return formatter.FormatObject(d)
}

// FormatCommands formats a list of commands and displays it with the requested
// format option.
func FormatCommands(commands []*models.Command, opt string) string {
	var d [][]kv
	for _, command := range commands {
		data := []kv{
			{"id", command.ID},
			{"name", command.Name},
			{"schedule", command.Schedule},
			{"image", command.Image},
			{"enabled", strconv.FormatBool(command.Enabled)},
		}
		d = append(d, data)
	}
	formatter := NewFormatter(opt)
	return formatter.FormatList(d)
}
