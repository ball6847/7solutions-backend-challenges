package utils

// SliceContains checks if a slice contains a specific string
func SliceContains(slice []string, item string) bool {
	for _, str := range slice {
		if str == item {
			return true
		}
	}
	return false
}
