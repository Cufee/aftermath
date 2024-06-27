package builder

func stringOr(options ...string) string {
	for _, value := range options {
		if value != "" {
			return value
		}
	}
	return ""
}
