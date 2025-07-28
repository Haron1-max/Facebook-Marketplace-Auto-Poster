package utils

import "strings"

// format description
func FormatDescription(description string) string {
	// Split the text by "..."
	parts := strings.Split(description, "...")

	// Trim spaces from each part
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i]) // Trim spaces from each part
	}

	// Prepend 🔥 to the first line
	if len(parts) > 0 {
		parts[0] = "🔥 " + parts[0]
	}

	// Prepend ✅ to middle lines
	for i := 1; i < len(parts)-1; i++ {
		parts[i] = "✅ " + parts[i]
	}

	// Prepend 📞 to the last line
	if len(parts) > 1 {
		parts[len(parts)-1] = "📞 " + parts[len(parts)-1]
	}

	// Join the parts with "...\n"
	description = strings.Join(parts, "\n\n")

	return description
}
