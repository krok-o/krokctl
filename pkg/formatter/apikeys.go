package formatter

import (
	"github.com/krok-o/krok/pkg/models"
)

// FormatCreatedApiKey formats an api key and displays it with the request
// format option including the generated random secret.
func FormatCreatedApiKey(key *models.APIKey, opt string) string {
	d := []kv{
		{"id", key.ID},
		{"name", key.Name},
		{"api-key-id", key.APIKeyID},
		{"api-key-secret", key.APIKeySecret},
		{"ttl", key.TTL},
		{"created_at", key.CreateAt.String()},
	}
	formatter := NewFormatter(opt)
	return formatter.FormatObject(d)
}

// FormatApiKey formats an api key and displays it with the request
// format option.
func FormatApiKey(key *models.APIKey, opt string) string {
	d := []kv{
		{"id", key.ID},
		{"name", key.Name},
		{"api-key-id", key.APIKeyID},
		{"ttl", key.TTL},
		{"created_at", key.CreateAt.String()},
	}
	formatter := NewFormatter(opt)
	return formatter.FormatObject(d)
}

// FormatApiKeys formats a list of api keys and displays it with the requested
// format option.
func FormatApiKeys(keys []*models.APIKey, opt string) string {
	var d [][]kv
	for _, key := range keys {
		data := []kv{
			{"id", key.ID},
			{"name", key.Name},
			{"api-key-id", key.APIKeyID},
			{"ttl", key.TTL},
			{"created_at", key.CreateAt.String()},
		}
		d = append(d, data)
	}
	formatter := NewFormatter(opt)
	return formatter.FormatList(d)
}
