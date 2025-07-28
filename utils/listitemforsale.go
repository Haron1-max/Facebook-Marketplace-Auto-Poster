package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

// var (
// 	duration = 20 * time.Microsecond
// )

type Item struct {
	Page        *rod.Page
	R           *rand.Rand
	ImageFiles  []string
	ListingType string
	Title       string
	Price       string
	Category    string
	Condition   string
	Description string
	Tags        []string
}

func NewItem(page *rod.Page, r *rand.Rand, imageFiles []string, listingType, title, price, category, condition, description string, tags []string) *Item {
	return &Item{
		Page:        page,
		R:           r,
		ImageFiles:  imageFiles,
		ListingType: listingType,
		Title:       title,
		Price:       price,
		Category:    category,
		Condition:   condition,
		Description: description,
		Tags:        tags,
	}
}

func (i Item) ListItemForSale() error {

	page := i.Page

	page.MustScreenshot("home.png")
	//os.Exit(1)

	// Locate the file input element on the page
	page.MustElement(`input[type="file"][accept="image/*,image/heif,image/heic"]`).MustSetFiles(i.ImageFiles...)

	// filesInput := page.MustElement(`input[type="file"][accept="image/*,image/heif,image/heic"]`)

	// // Upload each image file slowly
	// for _, file := range i.ImageFiles {
	// 	filesInput.MustSetFiles(file) // Set the current file
	// 	time.Sleep(2 * time.Second)   // Delay before the next upload
	// }

	fmt.Println("Images inserted")

	// get title input
	page.MustElement(`label[aria-label="Title"]`).MustInput(i.Title)

	// for _, char := range i.Title {
	// 	titleInput.MustInput(string(char)) // Type one character
	// 	time.Sleep(duration)               // Adjust delay as needed
	// }

	fmt.Println("Title inserted")

	// get price input
	page.MustElement(`label[aria-label="Price"]`).MustInput(i.Price)

	// for _, char := range i.Price {
	// 	priceInput.MustInput(string(char)) // Type one character
	// 	time.Sleep(duration)               // Adjust delay as needed
	// }

	fmt.Println("Price inserted")

	// get category input
	page.MustElement(`label[aria-label="Category"]`).MustClick()

	// Find all parent divs with data-visualcompletion="ignore-dynamic"
	cats := page.MustElements(`div[data-visualcompletion="ignore-dynamic"]`)

	for _, cat := range cats {
		if strings.ToLower(cat.MustText()) == i.Category {
			cat.MustClick()
			break
		}
	}

	fmt.Println("Category selected")

	// get condition input
	page.MustElement(`label[aria-label="Condition"]`).MustClick()

	// get all conditions
	options := i.Page.MustElements(`div[role="option"]`)

	fmt.Println("Conditions retrieved")

	for _, option := range options {
		if strings.ToLower(option.MustText()) == i.Condition {
			option.MustClick()
			break
		}
	}

	fmt.Println("Condition selected")

	// fmt.Println("Description inserted")
	page.MustElement(`label[aria-label="Description"]`).MustInput(i.Description)

	// for _, char := range i.Description {
	// 	descriptionInput.MustInput(string(char)) // Type one character
	// 	time.Sleep(duration * time.Millisecond)  // Adjust delay as needed
	// }

	fmt.Println("Description inserted")

	// get availability inputs
	// page.MustElement(`label[aria-label="Availability"]`).MustClick()

	// options = page.MustElements(`div[role="option"]`)

	// options[1].MustClick()

	allTags := len(i.Tags)

	if allTags > 0 {
		//get product tags textareas
		tagsInput := page.MustElement(`label[aria-label="Product tags"]`)

		for _, tag := range i.Tags {
			tag = strings.TrimSpace(tag)

			tagsInput.MustInput(tag)

			// for _, char := range tag {
			// 	tagsInput.MustInput(string(char))       // Type one character
			// 	time.Sleep(duration * time.Millisecond) // Adjust delay as needed
			// }

			err := i.Page.Keyboard.Press(input.Enter)
			if err != nil {
				return fmt.Errorf("error pressing enter key: %v", err)
			}
		}

	}

	fmt.Println("Tags inserted")

	// get "next" button
	page.MustElement(`div[aria-label="Next"]`).MustClick()

	fmt.Println("Next button clicked")

	// wait 10 seconds for the next page to load
	time.Sleep(3 * time.Second)

	// allGroups := page.MustElements(`div.x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xurb0ha.x1sxyh0.x1anpbxc.xyorhqc.xzboxd6.x14l7nz5 > div.x1n2onr6.x1ja2u2z.x9f619.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xz9dl7a`)

	// selectedGroups := 0

	groups := page.MustElements(`div[role="checkbox"]`)

	// if len(allGroups) == 2 {
	// 	joinedGroups := allGroups[0]

	// 	groups := joinedGroups.MustElements(`div[data-visualcompletion="ignore-dynamic"]`)

	// 	totalSuggestedGroups := len(groups)

	// 	fmt.Println("Total Suggested Groups:", totalSuggestedGroups)

	// 	r := i.R

	// 	switch {
	// 	case totalSuggestedGroups <= 20:
	// 		for _, group := range groups {
	// 			group.MustClick().MustScreenshot("group.png")
	// 			time.Sleep(500 * time.Millisecond)
	// 		}

	// 		selectedGroups = totalSuggestedGroups
	// 	default:
	// 		filteredGroups, totalFilteredGroups := FilterGroupsByMemberCount(groups, r)

	// 		r.Shuffle(totalFilteredGroups, func(i, j int) { filteredGroups[i], filteredGroups[j] = filteredGroups[j], filteredGroups[i] })

	// 		// tracked total selected groups
	// 		totalGroupsSelected := 0

	// 		// Click random groups up to the limit
	// 		for i := 0; i < totalFilteredGroups; i++ {
	// 			// Click the group
	// 			filteredGroups[i].MustClick().MustScreenshot("group.png")

	// 			// increase total groups selected
	// 			totalGroupsSelected++

	// 			if totalGroupsSelected == 20 {
	// 				break
	// 			}

	// 			// Add a small delay between clicks to simulate human interaction (optional)
	// 			time.Sleep(500 * time.Millisecond)
	// 		}

	// 		selectedGroups = totalGroupsSelected
	// 	}
	// }

	// fmt.Printf("%d groups selected\n", selectedGroups)

	// page.MustScreenshot("home.png")
	// os.Exit(1)

	fmt.Println("All Groups:", len(groups))

	for _, group := range groups {
		group.MustClick().MustScreenshot("group.png")
		time.Sleep(500 * time.Millisecond)
	}

	// get publish button
	page.MustElement(`div[aria-label="Publish"]`).MustClick()

	time.Sleep(30 * time.Second)

	fmt.Printf("Listed %s 🔥\n", i.Title)

	return nil
}
