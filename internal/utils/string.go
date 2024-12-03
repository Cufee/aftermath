package utils

func StringOr(s ...string) string {
	for _, s := range s {
		if s != "" {
			return s
		}
	}
	return ""
}
