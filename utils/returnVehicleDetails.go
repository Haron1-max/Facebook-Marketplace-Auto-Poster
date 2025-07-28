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

func ReturnVehicleDetails(subDir string, page *rod.Page, r *rand.Rand, imageFiles []string) (*Vehicle, error) {
	// Open the details.txt file within the subdirectory
	detailsFile := filepath.Join(subDir, "details.txt")
	file, err := os.Open(detailsFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	defer file.Close()

	// Initialize variables to hold the extracted fields
	var vehicleType, year, make, model, mileage, price, bodyStyle, hasCleanTitle, vehicleCondition, fuelType, transmission, description string

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "vehicle type:"):
			vehicleType = strings.TrimSpace(line[len("vehicle type:"):])
		case strings.HasPrefix(line, "year:"):
			year = strings.TrimSpace(line[len("year:"):])
		case strings.HasPrefix(line, "make:"):
			make = strings.TrimSpace(line[len("make:"):])
		case strings.HasPrefix(line, "model"):
			model = strings.ToLower(strings.TrimSpace(line[len("model:"):]))
		case strings.HasPrefix(line, "mileage"):
			mileage = strings.ToLower(strings.TrimSpace(line[len("mileage:"):]))
		case strings.HasPrefix(line, "price:"):
			price = strings.TrimSpace(line[len("price:"):])
		case strings.HasPrefix(line, "body style:"):
			bodyStyle = strings.TrimSpace(line[len("body style:"):])
		case strings.HasPrefix(line, "has clean title:"):
			hasCleanTitle = strings.TrimSpace(line[len("has clean title:"):])
		case strings.HasPrefix(line, "vehicle condition:"):
			vehicleCondition = strings.TrimSpace(line[len("vehicle condition:"):])
		case strings.HasPrefix(line, "fuel type:"):
			fuelType = strings.TrimSpace(line[len("fuel type:"):])
		case strings.HasPrefix(line, "transmission:"):
			transmission = strings.TrimSpace(line[len("transmission:"):])
		case strings.HasPrefix(line, "description:"):
			description = strings.TrimSpace(line[len("description:"):])
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	description = FormatDescription(description)

	// print extracted fields
	// fmt.Println("\nVehicle Type:", vehicleType)
	// fmt.Println("\nYear:", year)
	// fmt.Println("\nMake:", make)
	// fmt.Println("\nModel:", model)
	// fmt.Println("\nMileage:", mileage)
	// fmt.Println("\nPrice:", price)
	// fmt.Println("\nBody Style:", bodyStyle)
	// fmt.Println("\nHas Clean Title:", hasCleanTitle)
	// fmt.Println("\nVehicle Condition:", vehicleCondition)
	// fmt.Println("\nFuel Type:", fuelType)
	// fmt.Println("\nTransmission:", transmission)
	fmt.Println("\nDescription:", description)

	vehicle := NewVehicle(page, r, imageFiles, vehicleType, year, make, model, mileage, price, bodyStyle, hasCleanTitle, vehicleCondition, fuelType, transmission, description)

	return vehicle, nil
}
