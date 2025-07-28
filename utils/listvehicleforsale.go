package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-rod/rod"
)

type Vehicle struct {
	Page             *rod.Page
	R                *rand.Rand
	ImageFiles       []string
	VehicleType      string
	Year             string
	Make             string
	Model            string
	Mileage          string
	Price            string
	BodyStyle        string
	HasCleanTitle    string
	VehicleCondition string
	FuelType         string
	Transmission     string
	Description      string
}

func NewVehicle(page *rod.Page, r *rand.Rand, imageFiles []string, vehicleType, year, make, model, mileage, price, bodyStyle, hasCleanTitle, vehicleCondition, fuelType, transmission, description string) *Vehicle {
	return &Vehicle{
		Page:             page,
		R:                r,
		VehicleType:      vehicleType,
		ImageFiles:       imageFiles,
		Year:             year,
		Make:             make,
		Model:            model,
		Mileage:          mileage,
		Price:            price,
		BodyStyle:        bodyStyle,
		HasCleanTitle:    hasCleanTitle,
		VehicleCondition: vehicleCondition,
		FuelType:         fuelType,
		Transmission:     transmission,
		Description:      description,
	}
}

func (v Vehicle) ListVehicleForSale() {
	page := v.Page

	page.MustScreenshot("home.png")

	page.MustElement(`input[type="file"][accept="image/*,image/heif,image/heic"]`).MustSetFiles(v.ImageFiles...)

	time.Sleep(5 * time.Second)

	var options rod.Elements

	page.MustElement(`label[aria-label="Vehicle type"]`).MustClick()

	options = page.MustElements(`div[role="option"]`)

	for _, o := range options {
		if strings.Contains(o.MustText(), v.VehicleType) {
			o.MustClick()
			break
		}
	}

	page.MustElement(`label[aria-label="Year"]`).MustClick()

	options = page.MustElements(`div[role="option"]`)

	for _, o := range options {
		if o.MustText() == v.Year {
			o.MustClick()
			break
		}
	}

	page.MustElement(`label[aria-label="Make"]`).MustInput(v.Make)

	page.MustElement(`label[aria-label="Model"]`).MustInput(v.Model)

	page.MustElement(`label[aria-label="Mileage"]`).MustInput(v.Mileage)

	page.MustElement(`label[aria-label="Price"]`).MustInput(v.Price)

	page.MustElement(`label[aria-label="Body style"]`).MustClick()

	options = page.MustElements(`div[role="option"]`)

	for _, o := range options {
		if o.MustText() == v.BodyStyle {
			o.MustClick()
			break
		}
	}

	if v.HasCleanTitle == "Yes" {
		page.MustElement(`input[aria-label="This vehicle has a clean title."]`).MustClick()
	}

	page.MustElement(`label[aria-label="Vehicle condition"]`).MustClick()

	options = page.MustElements(`div[role="option"]`)

	for _, o := range options {
		if o.MustText() == v.VehicleCondition {
			o.MustClick()
			break
		}
	}

	page.MustElement(`label[aria-label="Fuel type"]`).MustClick()

	options = page.MustElements(`div[role="option"]`)

	for _, o := range options {
		if o.MustText() == v.FuelType {
			o.MustClick()
			break
		}
	}

	page.MustElement(`label[aria-label="Transmission"]`).MustClick()

	for _, o := range options {
		if o.MustText() == v.Transmission {
			o.MustClick()
			break
		}
	}

	page.MustElement(`label[aria-label="Description"]`).MustInput(v.Description)

	page.MustElement(`div[aria-label="Next"]`).MustClick()

	time.Sleep(3 * time.Second)

	allGroups := page.MustElements(`div.x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xurb0ha.x1sxyh0.x1anpbxc.xyorhqc.xzboxd6.x14l7nz5 > div.x1n2onr6.x1ja2u2z.x9f619.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xz9dl7a`)

	selectedGroups := 0

	if len(allGroups) == 2 {
		joinedGroups := allGroups[0]

		groups := joinedGroups.MustElements(`div[data-visualcompletion="ignore-dynamic"]`)

		totalSuggestedGroups := len(groups)

		fmt.Println("Total Suggested Groups:", totalSuggestedGroups)

		r := v.R

		switch {
		case totalSuggestedGroups <= 20:
			for _, group := range groups {
				group.MustClick().MustScreenshot("group.png")
				time.Sleep(500 * time.Millisecond)
			}

			selectedGroups = totalSuggestedGroups
		default:
			filteredGroups, totalFilteredGroups := FilterGroupsByMemberCount(groups, r)

			r.Shuffle(totalFilteredGroups, func(i, j int) { filteredGroups[i], filteredGroups[j] = filteredGroups[j], filteredGroups[i] })

			// tracked total selected groups
			totalGroupsSelected := 0

			// Click random groups up to the limit
			for i := 0; i < totalFilteredGroups; i++ {
				// Click the group
				filteredGroups[i].MustClick().MustScreenshot("group.png")

				// increase total groups selected
				totalGroupsSelected++

				if totalGroupsSelected == 20 {
					break
				}

				// Add a small delay between clicks to simulate human interaction (optional)
				time.Sleep(500 * time.Millisecond)
			}

			selectedGroups = totalGroupsSelected
		}
	}

	fmt.Printf("%d groups selected\n", selectedGroups)

	page.MustElement(`div[aria-label="Publish"]`).MustClick()

	time.Sleep(30 * time.Second)

	page.MustScreenshot("marketplace.png")

	fmt.Printf("Listed %s 🔥\n", v.Make)
}
