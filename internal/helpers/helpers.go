package helpers

func SetUpdatedContent(old string, new *string) string {
	if new == nil {
		return old
	}
	if *new == "" {
		return old
	}
	if *new == "null" {
		return ""
	}
	return *new
}
