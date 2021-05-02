package formatter

import (
	"github.com/krok-o/krok/pkg/models"
)

// FormatVaultSecret formats a single object of vault secret and displays it with the requested
// format option.
func FormatVaultSecret(secret *models.VaultSetting, opt string) string {
	d := []kv{
		{"name", secret.Key},
		{"value", secret.Value},
	}
	formatter := NewFormatter(opt)
	return formatter.FormatObject(d)
}

// FormatVaultSecrets formats a list of vault secrets and displays it with the requested
// format option.
func FormatVaultSecrets(secrets []string, opt string) string {
	var d [][]kv
	for _, r := range secrets {
		data := []kv{
			{"keys", r},
		}
		d = append(d, data)
	}
	formatter := NewFormatter(opt)
	return formatter.FormatList(d)
}
