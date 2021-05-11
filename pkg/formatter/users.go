package formatter

import (
	"github.com/krok-o/krok/pkg/models"
)

// FormatUser formats a user and displays it with the request
// format option.
func FormatUser(user *models.User, opt string) string {
	d := []kv{
		{"id", user.ID},
		{"display_name", user.DisplayName},
		{"email", user.Email},
		{"last_login", user.LastLogin.String()},
	}
	formatter := NewFormatter(opt)
	return formatter.FormatObject(d)
}

// FormatUsers formats a list of users and displays it with the requested
// format option.
func FormatUsers(users []*models.User, opt string) string {
	var d [][]kv
	for _, user := range users {
		data := []kv{
			{"id", user.ID},
			{"display_name", user.DisplayName},
			{"email", user.Email},
			{"last_login", user.LastLogin.String()},
		}
		d = append(d, data)
	}
	formatter := NewFormatter(opt)
	return formatter.FormatList(d)
}
