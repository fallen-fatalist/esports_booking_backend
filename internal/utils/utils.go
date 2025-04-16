package utils

func IsAnyEmpty(strs []string) bool {
	for _, str := range strs {
		if str == "" {
			return true
		}
	}
	return false
}
