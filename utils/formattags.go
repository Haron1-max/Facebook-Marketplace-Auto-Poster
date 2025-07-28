package utils

import "strings"

// format tags
func FormatTags(text string) []string {
	tags := strings.Split(text, ",")

	for i := len(tags) - 1; i >= 0; i-- {
		if tags[i] == "" {
			tags = tags[:i]
		} else {
			break
		}
	}

	return tags
}
