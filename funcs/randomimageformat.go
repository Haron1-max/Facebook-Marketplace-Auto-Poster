package funcs

import (
	"time"

	"math/rand"
)

// Function to generate a random image format
func GetRandomImageFormat() string {
	imageFormats := []string{
		"jpg",
		"png",
		"tiff",
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Seed the random number generator with current time

	// Generate a random index within the range of imageFormats slice
	randomIndex := r.Intn(len(imageFormats))

	return imageFormats[randomIndex]
}
