package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReturnListingType(subDir string) (string, error) {
	// Open the details.txt file within the subdirectory
	detailsFile := filepath.Join(subDir, "details.txt")

	file, err := os.Open(detailsFile)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}

	defer file.Close()

	// Initialize variables to hold the extracted fields
	var listingType string

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "listing type:"):
			listingType = strings.TrimSpace(line[len("listing type:"):])
		}
	}

	return listingType, nil
}
