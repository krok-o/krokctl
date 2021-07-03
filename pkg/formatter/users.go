package formatter

import (
	"fmt"
	"strings"

	"github.com/krok-o/krok/pkg/models"
)

// FormatUser formats a user and displays it with the request
// format option.
func FormatUser(user *models.User, opt string) string {
	var keys []string
	for _, key := range user.APIKeys {
		keys = append(keys, fmt.Sprintf("%d:%s", key.ID, key.APIKeyID))
	}
	d := []kv{
		{"id", user.ID},
		{"display_name", user.DisplayName},
		{"email", user.Email},
		{"last_login", user.LastLogin.String()},
		{"api_keys", strings.Join(keys, ", ")},
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
