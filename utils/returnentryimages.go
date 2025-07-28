package utils

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

func ReturnEntryImages(subDir string, r *rand.Rand) ([]string, error) {

	// Read files from the subdirectory
	subEntries, err := os.ReadDir(subDir)
	if err != nil {
		return []string{}, fmt.Errorf("error reading subdirectory: %v", err)
	}

	var imageFiles []string

	// Loop through files in the subdirectory
	for _, subEntry := range subEntries {
		if !subEntry.IsDir() {
			filePath := filepath.Join(subDir, subEntry.Name())
			ext := filepath.Ext(filePath)
			// Check if the file is not a .txt or .md file
			if ext != ".txt" && ext != ".md" {
				//fmt.Println("IMAGE:", filePath)
				// Collect file paths to attach
				imageFiles = append(imageFiles, filePath)
			}
		}
	}

	// Ensure the image list is shuffled before posting
	r.Shuffle(len(imageFiles), func(i, j int) {
		imageFiles[i], imageFiles[j] = imageFiles[j], imageFiles[i]
	})

	// Limit the number of images to a maximum of 20
	if len(imageFiles) > 20 {
		imageFiles = imageFiles[:20]
	}

	return imageFiles, nil
}
