package formatter

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
