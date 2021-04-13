package formatter

import (
	"strconv"
	"strings"

	"github.com/krok-o/krok/pkg/models"
)

// FormatCommand formats a command and displays it with the request
// format option.
func FormatCommand(command *models.Command, opt string) string {
	listOfRepoNames := make([]string, 0)
	for _, r := range command.Repositories {
		listOfRepoNames = append(listOfRepoNames, r.Name)
	}
	d := []kv{
		{"id", strconv.FormatInt(int64(command.ID), 10)},
		{"name", command.Name},
		{"hash", command.Hash},
		{"location", command.Location},
		{"filename", command.Filename},
		{"schedule", command.Schedule},
		{"enabled", strconv.FormatBool(command.Enabled)},
		{"repositories", strings.Join(listOfRepoNames, ",")},
	}
	formatter := NewFormatter(opt)
	return formatter.FormatObject(d)
}

// FormatCommands formats a list of commands and displays it with the requested
// format option.
func FormatCommands(commands []*models.Command, opt string) string {
	d := [][]kv{}
	for _, command := range commands {
		data := []kv{
			{"id", strconv.FormatInt(int64(command.ID), 10)},
			{"name", command.Name},
			{"hash", command.Hash},
			{"location", command.Location},
			{"filename", command.Filename},
			{"schedule", command.Schedule},
			{"enabled", strconv.FormatBool(command.Enabled)},
		}
		d = append(d, data)
	}
	formatter := NewFormatter(opt)
	return formatter.FormatList(d)
}