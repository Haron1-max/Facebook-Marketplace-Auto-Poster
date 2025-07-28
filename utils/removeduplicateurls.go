package utils

func RemoveDuplicateURLs(elements []string) []string {
	// Use a map to keep track of seen elements
	seen := make(map[string]bool)
	var result []string

	for _, element := range elements {
		if !seen[element] {
			seen[element] = true
			result = append(result, element)
		}
	}

	return result
}
