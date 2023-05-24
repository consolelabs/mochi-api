package vault

func validateVaultName(name string) bool {
	return len(name) <= 24
}
