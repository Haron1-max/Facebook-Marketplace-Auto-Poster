package funcs

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/haron1996/fb/config"
	"github.com/haron1996/fb/cookies"
)

// var (
// 	c_user   = "61560452168137"
// 	datr     = "mpH7ZpR8c0msYQfV5xS5gZDD"
// 	fr       = "1pw5gRyxhG4Q0hcbw.AWUl9sux4imO8v7-YOCH-4mI0D4.Bm_DEe..AAA.0.0.Bm_DTw.AWVh7VWpMMI"
// 	presence = "C%7B%22t3%22%3A%5B%5D%2C%22utc3%22%3A1727804664047%2C%22v%22%3A1%7D"
// 	sb       = "Cy3sZisFoVjpIHZIEIDet9N5"
// 	wd       = "1366x681"
// 	xs       = "28%3Ae4c9MvJIwKl1Yg%3A2%3A1727804654%3A-1%3A11298"
// )

func PostSingleItemToMarketplace() {
	// load config files
	config, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// Launch the browser with specific configurations
	u := launcher.New().UserDataDir("~/.config/google-chrome").Leakless(true).NoSandbox(true).Headless(true).MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	defer browser.MustClose() // Ensure the browser closes when main function exits

	// Root directory containing subdirectories with images
	root := config.Root
	entries, err := os.ReadDir(root)
	if err != nil {
		fmt.Println("Error reading root directory:", err)
		return
	}

	// Shuffle the entries slice
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Shuffle(len(entries), func(i, j int) {
		entries[i], entries[j] = entries[j], entries[i]
	})

	for _, entry := range entries {
		// START LOGIN
		dir := "~/.config/google-chrome"

		u := launcher.New().UserDataDir(dir).Leakless(true).NoSandbox(true).Headless(true).MustLaunch()

		browser := rod.New().ControlURL(u).MustConnect()

		defer browser.MustClose()

		// You can now use this session to interact with pages that require authentication
		page := browser.MustPage("https://web.facebook.com/").MustWaitLoad()

		pageHasLoginForm, _, err := page.Has(`form[data-testid="royal_login_form"]`)
		if err != nil {
			log.Println("Error checking if page has login form:", err)
			return
		}

		switch {
		case pageHasLoginForm:
			// Define the session cookies
			cookies := []*proto.NetworkCookieParam{
				{
					Name:     "c_user",
					Value:    cookies.C_user,
					Domain:   ".facebook.com",
					Path:     "/",
					HTTPOnly: true,
					Secure:   true,
					SameSite: "None",
					Priority: "Medium",
				},
				{
					Name:     "datr",
					Value:    cookies.Datr,
					Domain:   ".facebook.com",
					Path:     "/",
					HTTPOnly: true,
					Secure:   true,
					SameSite: "None",
					Priority: "Medium",
				},
				{
					Name:     "fr",
					Value:    cookies.Fr,
					Domain:   ".facebook.com",
					Path:     "/",
					HTTPOnly: true,
					Secure:   true,
					SameSite: "None",
					Priority: "Medium",
				},
				{
					Name:     "presence",
					Value:    cookies.Presence,
					Domain:   ".facebook.com",
					Path:     "/",
					HTTPOnly: true,
					Secure:   true,
					SameSite: "None",
					Priority: "Medium",
				},
				{
					Name:     "sb",
					Value:    cookies.Sb,
					Domain:   ".facebook.com",
					Path:     "/",
					HTTPOnly: true,
					Secure:   true,
					SameSite: "None",
					Priority: "Medium",
				},
				{
					Name:     "wd",
					Value:    cookies.Wd,
					Domain:   ".facebook.com",
					Path:     "/",
					HTTPOnly: true,
					Secure:   true,
					SameSite: "None",
					Priority: "Medium",
				},
				{
					Name:     "xs",
					Value:    cookies.Xs,
					Domain:   ".facebook.com",
					Path:     "/",
					HTTPOnly: true,
					Secure:   true,
					SameSite: "None",
					Priority: "Medium",
				},
			}

			// Inject the session cookie
			err := browser.SetCookies(cookies)
			if err != nil {
				fmt.Println("Failed to set session cookie:", err)
				return
			}

			// check if cookies are valid
			page := browser.MustPage("https://web.facebook.com/").MustWaitLoad()

			pageHasLoginForm, _, err := page.Has(`form[data-testid="royal_login_form"]`)
			if err != nil {
				log.Println("Error checking if page has login form:", err)
				return
			}
			switch {
			case pageHasLoginForm:
				fmt.Println("Invalid or expired cookies 😞")
				os.Exit(1)
			default:
				fmt.Println("Log in complete 😊")
			}
		default:
			fmt.Println("User is logged in 😊")
		}
		// END LOGIN

		// Open the Facebook Marketplace item creation page
		page.MustNavigate("https://web.facebook.com/marketplace/create/item").MustWaitLoad()

		// Locate the file input element on the page
		fileInput := page.MustElement(`input[type="file"]`)

		// Path to the current subdirectory
		subDir := filepath.Join(root, entry.Name())

		fmt.Println("DIRECTORY:", subDir)

		// Read files from the subdirectory
		subEntries, err := os.ReadDir(subDir)
		if err != nil {
			fmt.Println("Error reading subdirectory:", err)
			continue
		}

		var imageFiles []string
		// Loop through files in the subdirectory
		for _, subEntry := range subEntries {
			if !subEntry.IsDir() {
				filePath := filepath.Join(subDir, subEntry.Name())
				ext := filepath.Ext(filePath)
				if ext != ".txt" && ext != ".md" { // Check if the file is not a .txt or .md file
					fmt.Println("IMAGE:", filePath)

					// Collect file paths to attach
					imageFiles = append(imageFiles, filePath)
				}
			}
		}

		// Ensure the image list is shuffled before posting
		r.Shuffle(len(imageFiles), func(i, j int) {
			imageFiles[i], imageFiles[j] = imageFiles[j], imageFiles[i]
		})

		// Attach collected files to the file input element
		if len(imageFiles) > 0 {
			fileInput.MustSetFiles(imageFiles...)
		}

		time.Sleep(5 * time.Second)

		// Open the details.txt file within the subdirectory
		detailsFile := filepath.Join(subDir, "details.txt")
		file, err := os.Open(detailsFile)
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}
		defer file.Close()

		// Initialize variables to hold the extracted fields
		var title, price, category, condition, description, tags string

		// Create a new scanner to read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			switch {
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
				tags = strings.TrimSpace(line[len("tags:"):])
			}
		}

		// Check for errors during scanning
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// Print the extracted fields
		fmt.Println("\nTitle:", title)
		fmt.Println("\nPrice:", price)
		fmt.Println("\nCategory:", category)
		fmt.Println("\nCondition:", condition)

		formattedDescription := formatDescription(description, category)

		fmt.Println("\nDescription:\n" + formattedDescription)

		formattedTags := formatTags(tags)

		fmt.Println("\nTags:", formattedTags)

		fmt.Println("Past tags")

		// get inputs wrapper
		marketplace, err := page.Element(`div[aria-label="Marketplace"]`)
		if err != nil {
			fmt.Println("Error getting marketplace div:", err)
			return
		}

		// get title input
		titleInput, err := marketplace.Element(`label[aria-label="Title"]`)
		if err != nil {
			fmt.Println("Error getting title input:", err)
			return
		}

		// input title text
		err = titleInput.Input(title)
		if err != nil {
			fmt.Println("Error inputing title text:", err)
			return
		}

		fmt.Println("Past title")

		// get price input
		priceInput, err := marketplace.Element(`label[aria-label="Price"]`)
		if err != nil {
			fmt.Println("Error getting price input:", err)
			return
		}

		// input price text
		err = priceInput.Input(price)
		if err != nil {
			fmt.Println("Error inputing price text:", err)
			return
		}

		fmt.Println("Past price")

		// get category input
		categoryInput, err := marketplace.Element(`label[aria-label="Category"]`)
		if err != nil {
			fmt.Println("Error getting category input:", err)
			return
		}

		// click category input
		err = categoryInput.Click("left", 1)
		if err != nil {
			fmt.Println("Error clicking category input:", err)
			return
		}

		// Find all parent divs with data-visualcompletion="ignore-dynamic"
		cats := page.MustElements(`div[data-visualcompletion="ignore-dynamic"]`)

		for _, cat := range cats {
			if strings.ToLower(cat.MustText()) == category {
				cat.MustClick()
			}
		}

		fmt.Println("Past category")

		// get condition input
		conditionInput, err := marketplace.Element(`label[aria-label="Condition"]`)
		if err != nil {
			fmt.Println("Error getting condition input:", err)
			return
		}

		// click conditionInput
		err = conditionInput.Click("left", 1)
		if err != nil {
			fmt.Println("Error clicking condition input:", err)
			return
		}

		// get all conditions
		conditions, err := page.Elements(`div[role="option"]`)
		if err != nil {
			fmt.Println("Error getting all conditions:", err)
			return
		}

		for _, cond := range conditions {
			if strings.ToLower(cond.MustText()) == condition {
				cond.MustClick()
			}
		}

		fmt.Println("Past conditions")

		// get description label
		page.MustElementR("label", "Description").MustWaitVisible()

		// Locate the textarea within the label
		textarea := page.MustElementR("label", "Description").MustElement("textarea")

		// input description text
		err = textarea.Input(formattedDescription)
		if err != nil {
			fmt.Println("Error inputing description text:", err)
			return
		}

		fmt.Println("Past description")

		// get availability inputs
		availabilitySelector := `label[aria-label="Availability"]`
		availabilityInputs, err := page.Elements(availabilitySelector)
		if err != nil {
			fmt.Println("Error getting availability inputs:", err)
			return
		}

		// open available options
		if len(availabilityInputs) > 0 {
			err := availabilityInputs[0].Click("left", 1)
			if err != nil {
				fmt.Println("Error opening available options:", err)
				return
			}
		}

		// find all available options
		divSelector := `div.html-div.xe8uvvx.xdj266r.x11i5rnm.xat24cr.x1mh8g0r.xexx8yu.x4uap5.x18d9i69.xkhd6sd.x6s0dn4.x78zum5.x1q0g3np.x1iyjqo2.x1qughib.xeuugli`

		availableOptions, err := page.Elements(divSelector)
		if err != nil {
			fmt.Println("Error getting available options:", err)
			return
		}

		// click neccessary availability option
		if len(availableOptions) > 0 {
			err := availableOptions[1].Click("left", 1)
			if err != nil {
				fmt.Println("Error selecting available option:", err)
				return
			}
		}

		fmt.Println("Availability option clicked")

		if len(formattedTags) > 0 {
			//get product tags textareas
			productTagsTextareas, err := page.Elements(`label[aria-label="Product tags"]`)
			if err != nil {
				fmt.Println("Error getting product tags text area:", err)
				return
			}

			// input product tags
			if len(productTagsTextareas) > 0 {

				for _, tag := range formattedTags {
					tag = strings.TrimSpace(tag)

					// get product tags textarea
					productTagsTextarea, err := productTagsTextareas[0].Element("textarea")
					if err != nil {
						fmt.Println("Error getting product tags text area:", err)
						return
					}

					// input tag
					err = productTagsTextarea.Input(tag)
					if err != nil {
						fmt.Println("Error inputing product tag:", err)
						return
					}

					err = page.Keyboard.Press(input.Enter)
					if err != nil {
						fmt.Println("Error pressing ENTER key")
						return
					}

				}

			}
		}

		// get "next" button
		nextButton, err := page.Element(`div[aria-label="Next"]`)
		if err != nil {
			fmt.Println("Error getting next button:", err)
			return
		}

		// click "next" button
		err = nextButton.Click("left", 1)
		if err != nil {
			fmt.Println("Error clicking next button:", err)
			return
		}

		fmt.Println("Next button clicked")

		// wait 10 seconds for the next page to load
		time.Sleep(10 * time.Second)

		var sg *rod.Element

		fmt.Println("Getting suggested groups")

		sgs := marketplace.MustElements(`div.x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xurb0ha.x1sxyh0.x1anpbxc.xyorhqc.xzboxd6.x14l7nz5 > div.x1n2onr6.x1ja2u2z.x9f619.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xz9dl7a`)

		totalSgs := len(sgs)

		switch {
		case totalSgs > 0:
			sg = sgs[0]
		default:
			fmt.Println("No suggested groups found")
			page.MustScreenshot("home.png")
			os.Exit(1)
		}

		fmt.Println("SG:", sg) // if sg is nil just publish else go groups way

		fmt.Println("about to get all suggested groups")

		// get suggested groups
		groups, err := sg.Elements(`div[data-visualcompletion="ignore-dynamic"]`)
		if err != nil {
			fmt.Println("Error getting suggested groups:", err)
			return
		}

		fmt.Println("calculating suggested groups")

		totalGroups := len(groups)

		fmt.Println("total suggested groups:", totalGroups)

		r.Shuffle(totalGroups, func(i, j int) { groups[i], groups[j] = groups[j], groups[i] })

		// Regular expression to capture the member count (e.g., "21K members" or "10K+ members")
		memberCountRegex := regexp.MustCompile(`(\d+(\.\d+)?)([KM]?)\s*members`)

		// groups to be selected
		var groupsToSelect []*rod.Element

		for _, group := range groups {
			members := group.MustElements(`div.x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xeuugli.xg83lxy.x1h0ha7o.x1120s5i.x1nn3v0j`)[0].MustText()

			// Find the member count in the text
			match := memberCountRegex.FindStringSubmatch(members)
			if match == nil {
				fmt.Println("No member count found for group")
				continue
			}

			// Parse the member count (e.g., "21K" -> 21000)
			memberCountStr := match[1]
			multiplier := match[3] // K for thousands, M for millions

			// Convert the member count to an integer
			memberCount, err := strconv.ParseFloat(memberCountStr, 64)
			if err != nil {
				fmt.Println("Error parsing member count:", err)
				continue
			}

			// Adjust for K or M
			if multiplier == "K" {
				memberCount *= 1000
			} else if multiplier == "M" {
				memberCount *= 1000000
			}

			//memberCount = RoundToNearestThousand(memberCount)
			if memberCount >= 20000 {
				//group.MustClick()
				//time.Sleep(3 * time.Second)
				//group.MustScreenshot("group.png")
				groupsToSelect = append(groupsToSelect, group)
			}
		}

		totalGroupsToSelect := len(groupsToSelect)

		fmt.Println("total groups to select:", totalGroupsToSelect)

		r.Shuffle(totalGroupsToSelect, func(i, j int) { groupsToSelect[i], groupsToSelect[j] = groupsToSelect[j], groupsToSelect[i] })

		// Determine how many groups to click (up to 50, or the available number of groups)
		limit := 50

		if totalGroupsToSelect < 50 {
			limit = totalGroupsToSelect
		}

		// tracked total selected groups
		totalGroupsSelected := 0

		// Click random groups up to the limit
		for i := 0; i < limit; i++ {
			// Click the group
			groupsToSelect[i].MustClick().MustScreenshot("group.png")

			// increase total groups selected
			totalGroupsSelected++

			if totalGroupsSelected == 20 {
				break
			}

			// Add a small delay between clicks to simulate human interaction (optional)
			time.Sleep(1 * time.Second)
		}

		// sg.MustElements(`div.x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xr9ek0c.xjpr12u.xzboxd6.x14l7nz5`)[1].MustScreenshot("facebook.png")

		// get publish button
		publishButton, err := page.Element(`div[aria-label="Publish"]`)
		if err != nil {
			fmt.Println("Error getting publish button:", err)
			return
		}

		// click publish button to publish ad
		err = publishButton.Click("left", 1)
		if err != nil {
			fmt.Println("Error publishing ad:", err)
			return
		}

		time.Sleep(15 * time.Second)

		fmt.Printf("%s posted successfully 🔥\n", title)

		fmt.Println("==========================================================")
	}
}

func formatDescription(text, category string) string {
	// Split the text by "..."
	parts := strings.Split(text, "...")

	// Trim spaces from each part
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i]) // Trim spaces from each part
	}

	switch {
	case category == strings.ToLower("Mobile phones"):
		// Prepend 📱 to the first line
		if len(parts) > 0 {
			parts[0] = "📱 " + parts[0]
		}

		// Prepend 🔥 to middle lines
		for i := 1; i < len(parts)-1; i++ {
			parts[i] = "🔥 " + parts[i]
		}

		// Prepend 📞 to the last line
		if len(parts) > 1 {
			parts[len(parts)-1] = "📞 " + parts[len(parts)-1]
		}
	default:
		// Prepend ✅ to the first line
		if len(parts) > 0 {
			parts[0] = "✅ " + parts[0]
		}

		// Prepend ✅ to middle lines
		for i := 1; i < len(parts)-1; i++ {
			parts[i] = "✅ " + parts[i]
		}

		// Prepend ✅ to the last line
		if len(parts) > 1 {
			parts[len(parts)-1] = "✅ " + parts[len(parts)-1]
		}
	}

	// Join the parts with "...\n"
	formattedtext := strings.Join(parts, "\n\n")

	return formattedtext
}

func formatTags(text string) []string {
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

// func RoundToNearestThousand(num float64) float64 {
// 	return math.Round(num/1000) * 1000
// }
