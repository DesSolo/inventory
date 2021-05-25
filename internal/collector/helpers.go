package collector

func isUniq(newItem string, items []string) bool {
	for _, item := range items {
		if newItem == item {
			return false
		}
	}
	return true
}
