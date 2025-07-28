package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-rod/rod"
)

func ReturnItemDetails(subDir string, page *rod.Page, r *rand.Rand, imageFiles []string) (*Item, error) {
	// Open the details.txt file within the subdirectory
	detailsFile := filepath.Join(subDir, "details.txt")
	file, err := os.Open(detailsFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	defer file.Close()

	// Initialize variables to hold the extracted fields
	var listingType, title, price, category, condition, description, tagsString string

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "listing type:"):
			listingType = strings.TrimSpace(line[len("listing type:"):])
		case strings.HasPrefix(line, "title:"):
			title = strings.TrimSpace(line[len("title:"):])
		case strings.HasPrefix(line, "price:"):
			price = strings.TrimSpace(line[len("price:"):])
		case strings.HasPrefix(line, "category"):
			category = strings.ToLower(strings.TrimSpace(line[len("category:"):]))
		case strings.HasPrefix(line, "condition"):
			condition = strings.ToLower(strings.TrimSpace(line[len("condition:"):]))
		case strings.HasPrefix(line, "description:"):
			description = strings.TrimSpace(line[len("description:"):])
		case strings.HasPrefix(line, "tags:"):
			tagsString = strings.TrimSpace(line[len("tags:"):])
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// format description
	//description = FormatDescription(description)

	// Split the text by "..."
	parts := strings.Split(description, "...")

	// Trim spaces from each part
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	// Join the parts with "...\n"
	description = strings.Join(parts, "\n\n")

	// format tags
	tags := FormatTags(tagsString)

	// Print the extracted fields
	// fmt.Println("\nListing Type:", listingType)
	// fmt.Println("\nTitle:", title)
	// fmt.Println("\nPrice:", price)
	// fmt.Println("\nCategory:", category)
	// fmt.Println("\nCondition:", condition)
	// fmt.Println("\nDescription:\n" + description)
	// fmt.Println("\nTags:", tags)
	fmt.Println(description)

	item := NewItem(page, r, imageFiles, listingType, title, price, category, condition, description, tags)

	return item, nil
}
