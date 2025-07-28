package funcs

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/haron1996/fb/config"
	"github.com/haron1996/fb/utils"
)

func PostToGroups(browser *rod.Browser, page *rod.Page) {
	// Launch the browser with specific configurations
	// u := launcher.New().UserDataDir("~/.config/google-chrome").Leakless(true).NoSandbox(true).Headless(true).MustLaunch()
	//browser := rod.New().ControlURL(u).MustConnect()
	defer browser.MustClose() // Ensure the browser closes when main function exits

	page = page.MustNavigate("https://web.facebook.com/groups/joins").MustWaitLoad()

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Println("Error loading config:", err)
		return
	}
	// Root directory containing subdirectories with images
	root := config.Root
	entries, err := os.ReadDir(root)
	if err != nil {
		fmt.Println("Error reading root directory:", err)
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// var anchors rod.Elements

	// for i := 0; i < 100; i++ {
	// 	page.Mouse.MustScroll(0, 500)
	// 	time.Sleep(10 * time.Second)
	// 	// get groups joined container
	// 	classSelector := ".x9f619.x1gryazu.xkrivgy.x1ikqzku.x1h0ha7o.xg83lxy.xh8yej3"

	// 	parents := page.MustElements(classSelector)

	// 	var parent *rod.Element

	// 	parentsLength := len(parents)

	// 	log.Println(parentsLength)

	// 	if parentsLength == 1 {
	// 		parent = parents[0]
	// 	} else if parentsLength == 2 {
	// 		parent = parents[1]
	// 	} else {
	// 		log.Println("no parents found")
	// 		return
	// 	}

	// 	selector := ".x1i10hfl.x1qjc9v5.xjbqb8w.xjqpnuy.xa49m3k.xqeqjp1.x2hbi6w.x13fuv20.xu3j5b3.x1q0q8m5.x26u7qi.x972fbf.xcfux6l.x1qhh985.xm0m39n.x9f619.x1ypdohk.xdl72j9.x2lah0s.xe8uvvx.xdj266r.x11i5rnm.xat24cr.x1mh8g0r.x2lwn1j.xeuugli.xexx8yu.x4uap5.x18d9i69.xkhd6sd.x1n2onr6.x16tdsg8.x1hl2dhg.xggy1nq.x1ja2u2z.x1t137rt.x1o1ewxj.x3x9cwd.x1e5q0jg.x13rtm0m.x1q0g3np.x87ps6o.x1lku1pv.x1rg5ohu.x1a2a7pz.xh8yej3"

	// 	links := parent.MustElements(selector)

	// 	anchors = append(anchors, links...)
	// }

	// VERSION 1

	// for _, a := range anchors {
	// 	href := *a.MustAttribute("href")
	// 	fmt.Println(href)
	// 	page = browser.MustPage(href).MustWaitLoad()
	// 	// time.Sleep(10 * time.Second)
	// 	// page.MustElement(`div[aria-label="Sell Something"]`).MustClick()
	// 	// time.Sleep(10 * time.Second)
	// 	// page.MustElements("span.x1lliihq.x1iyjqo2")[0].MustClick()
	// 	// load config files
	// 	// config, err := config.LoadConfig(".")
	// 	// if err != nil {
	// 	// 	log.Println("Error loading config:", err)
	// 	// 	return
	// 	// }
	// 	// // Root directory containing subdirectories with images
	// 	// root := config.Root
	// 	// entries, err := os.ReadDir(root)
	// 	// if err != nil {
	// 	// 	fmt.Println("Error reading root directory:", err)
	// 	// 	return
	// 	// }
	// 	// // Shuffle the entries slice
	// 	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 	// r.Shuffle(len(entries), func(i, j int) {
	// 	// 	entries[i], entries[j] = entries[j], entries[i]
	// 	// })

	// 	// get images selector
	// 	// imagesSelector := page.MustElement("div.x1gslohp.x1swvt13.x1pi30zi")

	// 	// card := page.MustElement("div.x9f619.x1ja2u2z.x1k90msu.x6o7n8i.x1qfuztq.x10l6tqk.x17qophe.x13vifvy.x1hc1fzr.x71s49j.xh8yej3")

	// 	fmt.Println("reached here...")

	// 	for _, entry := range entries {
	// 		log.Println("started looping entries...")
	// 		time.Sleep(10 * time.Second)
	// 		pageHasSellBtn, sellBtn, err := page.Has(`div[aria-label="Sell Something"]`)
	// 		if err != nil {
	// 			log.Println("Error checking if page/group has sell button:", err)
	// 			return
	// 		}

	// 		if !pageHasSellBtn {
	// 			continue
	// 		}
	// 		// sellBtn, err := page.Element(`div[aria-label="Sell Something"]`)
	// 		// if err != nil {
	// 		// 	log.Println("Error getting sell button element:", err)
	// 		// 	return
	// 		// }
	// 		err = sellBtn.Click("left", 1)
	// 		if err != nil {
	// 			log.Println("Error clicking sale button element:", err)
	// 			return
	// 		}
	// 		fmt.Println("reached here...")
	// 		time.Sleep(10 * time.Second)
	// 		page.MustElements("span.x1lliihq.x1iyjqo2")[0].MustWaitLoad().MustClick()
	// 		fmt.Println("reached here...")
	// 		time.Sleep(10 * time.Second)
	// 		page.MustScreenshot("start.png")
	// 		imagesSelector := page.MustElement("div.x1gslohp.x1swvt13.x1pi30zi")
	// 		fmt.Println("reached here...")
	// 		card := page.MustElement("div.x9f619.x1ja2u2z.x1k90msu.x6o7n8i.x1qfuztq.x10l6tqk.x17qophe.x13vifvy.x1hc1fzr.x71s49j.xh8yej3")
	// 		fmt.Println("reached here...?")
	// 		// Locate the file input element on the page
	// 		fileInput := imagesSelector.MustElement(`input[type="file"]`)
	// 		fmt.Println("file input located....")

	// 		//log.Println(fileInput.HTML())

	// 		// Path to the current subdirectory
	// 		subDir := filepath.Join(root, entry.Name())
	// 		fmt.Println("DIRECTORY:", subDir)

	// 		// Read files from the subdirectory
	// 		subEntries, err := os.ReadDir(subDir)
	// 		if err != nil {
	// 			fmt.Println("Error reading subdirectory:", err)
	// 			continue
	// 		}

	// 		var imageFiles []string
	// 		// Loop through files in the subdirectory
	// 		for _, subEntry := range subEntries {
	// 			if !subEntry.IsDir() {
	// 				filePath := filepath.Join(subDir, subEntry.Name())
	// 				if filepath.Ext(filePath) != ".txt" { // Check if the file is not a .txt file
	// 					fmt.Println("IMAGE:", filePath)

	// 					// Collect file paths to attach
	// 					imageFiles = append(imageFiles, filePath)
	// 				}
	// 			}
	// 		}

	// 		// Attach collected files to the file input element
	// 		if len(imageFiles) > 0 {
	// 			fileInput.MustSetFiles(imageFiles...)
	// 		}

	// 		// Open the details.txt file within the subdirectory
	// 		detailsFile := filepath.Join(subDir, "details.txt")
	// 		file, err := os.Open(detailsFile)
	// 		if err != nil {
	// 			fmt.Println("Error opening file:", err)
	// 			continue
	// 		}
	// 		defer file.Close()

	// 		// Initialize variables to hold the extracted fields
	// 		var title, price, description string

	// 		// Create a new scanner to read the file line by line
	// 		scanner := bufio.NewScanner(file)
	// 		for scanner.Scan() {
	// 			line := scanner.Text()

	// 			switch {
	// 			case strings.HasPrefix(line, "title:"):
	// 				title = strings.TrimSpace(line[len("title:"):])
	// 			case strings.HasPrefix(line, "price:"):
	// 				price = strings.TrimSpace(line[len("price:"):])
	// 			case strings.HasPrefix(line, "description:"):
	// 				description = strings.TrimSpace(line[len("description:"):])
	// 			}
	// 		}

	// 		// Check for errors during scanning
	// 		if err := scanner.Err(); err != nil {
	// 			fmt.Println("Error reading file:", err)
	// 			return
	// 		}

	// 		// Print the extracted fields
	// 		fmt.Println("Title:", title)
	// 		fmt.Println("Price:", price)

	// 		// Split the text by "..."
	// 		parts := strings.Split(description, "...")

	// 		// Trim spaces from each part
	// 		for i := range parts {
	// 			parts[i] = strings.TrimSpace(parts[i])
	// 		}

	// 		// Join the parts with "...\n"
	// 		formattedDescription := strings.Join(parts, "\n\n")

	// 		fmt.Println("Description: " + formattedDescription)

	// 		fmt.Println("past description...")

	// 		page.MustScreenshot("process.png")

	// 		// set title
	// 		titleInput, err := page.Element(`label[aria-label="Title"]`)
	// 		if err != nil {
	// 			log.Println("Error getting title input:", err)
	// 			return
	// 		}

	// 		err = titleInput.Input(title)
	// 		if err != nil {
	// 			log.Println("Error inputing title:", err)
	// 			return
	// 		}

	// 		fmt.Println("past card...")

	// 		// set price
	// 		card.MustElement(`label[aria-label="Price"]`).MustInput(price)

	// 		fmt.Println("past here...")

	// 		// set condition
	// 		condsInput, err := card.Element(`label[aria-label="Condition"]`)
	// 		if err != nil {
	// 			log.Println("Error getting conditions input:", err)
	// 			return
	// 		}

	// 		err = condsInput.Click("left", 1)
	// 		if err != nil {
	// 			log.Println("Error clicking conditions input:", err)
	// 			return
	// 		}

	// 		conds, err := page.Elements(`div[role="option"]`)
	// 		if err != nil {
	// 			log.Println("Error getting all conditions:", err)
	// 			return
	// 		}

	// 		if len(conds) > 0 {
	// 			err := conds[0].Click("left", 1)
	// 			if err != nil {
	// 				log.Println("Error selecting condition:", err)
	// 				return
	// 			}
	// 		}

	// 		fmt.Println("past here...")

	// 		// show more details
	// 		moreDetails, err := card.Element("div.x6s0dn4.xkh2ocl.x1q0q8m5.x1qhh985.xu3j5b3.xcfux6l.x26u7qi.xm0m39n.x13fuv20.x972fbf.x9f619.x78zum5.x1q0g3np.x1iyjqo2.xs83m0k.x1qughib.xat24cr.x11i5rnm.x1mh8g0r.xdj266r.x2lwn1j.xeuugli.x18d9i69.x4uap5.xkhd6sd.xexx8yu.x1n2onr6.x1ja2u2z")
	// 		if err != nil {
	// 			log.Println("Error getting more details:", err)
	// 			return
	// 		}

	// 		err = moreDetails.Click("left", 1)
	// 		if err != nil {
	// 			log.Println("Error clicking on more details:", err)
	// 			return
	// 		}

	// 		fmt.Println("past here...")

	// 		//set description
	// 		page.MustElementR("label", "Description").MustWaitVisible()

	// 		textarea := page.MustElementR("label", "Description").MustElement("textarea")

	// 		err = textarea.Input(formattedDescription)
	// 		if err != nil {
	// 			log.Println("Error inputing description text:", err)
	// 			return
	// 		}

	// 		fmt.Println("past here...")

	// 		// set product tags
	// 		productTagsTextareas, err := card.Elements(`label[aria-label="Product tags"]`)
	// 		if err != nil {
	// 			log.Println("Error getting product tags text area:", err)
	// 			return
	// 		}

	// 		// input product tags
	// 		if len(productTagsTextareas) > 0 {
	// 			tags := []string{
	// 				"lipa mdogo mdogo smartphones",
	// 				"samsung lipa mdogo mdogo",
	// 				"iphone lipa mdogo mdogo",
	// 				"lipa mdogo mdogo phones",
	// 				"mkopa phones",
	// 				"m-kopa phones",
	// 				"lipa pole pole smartphones",
	// 				"samsung lipa pole pole",
	// 				"iphone lipa pole pole",
	// 			}

	// 			for _, tag := range tags {
	// 				// get product tags textarea
	// 				productTagsTextarea, err := productTagsTextareas[0].Element("textarea")
	// 				if err != nil {
	// 					log.Println("Error getting product tags text area:", err)
	// 					return
	// 				}

	// 				// input tag
	// 				err = productTagsTextarea.Input(tag)
	// 				if err != nil {
	// 					log.Println("Error inputing product tag:", err)
	// 					return
	// 				}
	// 				page.Keyboard.Press(input.Enter)
	// 			}
	// 		}

	// 		time.Sleep(10 * time.Second)

	// 		// go to next page
	// 		card.MustElement(`div[aria-label="Next"]`).MustClick()

	// 		time.Sleep(10 * time.Second)

	// 		// list in marketplace
	// 		page.MustElements("div.x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xr9ek0c.xjpr12u.xzboxd6.x14l7nz5")[2].MustClick()
	// 		fmt.Println("listed in marketplace...")

	// 		// list in groups
	// 		suggestedGroups, err := page.Elements(".x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xr9ek0c.xjpr12u.xzboxd6.x14l7nz5")
	// 		if err != nil {
	// 			log.Println("Error getting suggested groups wrappers:", err)
	// 			return
	// 		}

	// 		groups, err := suggestedGroups[4].Elements(`div[data-visualcompletion="ignore-dynamic"]`)
	// 		if err != nil {
	// 			log.Println("Error getting suggested groups:", err)
	// 			return
	// 		}

	// 		// calculate and log total groups
	// 		totalGroups := len(groups)

	// 		fmt.Printf("Total groups: %d\n", totalGroups)

	// 		// select up to 20 groups
	// 		if totalGroups > 20 {
	// 			//Click on up to 20 divs randomly
	// 			for i := 0; i < 1000 && i < len(groups); i++ {
	// 				// Generate a random index within the range of groups slice
	// 				randomIndex := r.Intn(len(groups))

	// 				// Click on the div at the random index
	// 				err := groups[randomIndex].Click("left", 1)
	// 				if err != nil {
	// 					log.Println("Error selecting suggested group randomly:", err)
	// 					return
	// 				}

	// 				// Remove the clicked element from the slice to avoid clicking it again
	// 				groups = append(groups[:randomIndex], groups[randomIndex+1:]...)
	// 			}

	// 		} else {
	// 			// Click all groups if the total is 20 or fewer
	// 			for i := 0; i < totalGroups; i++ {
	// 				err := groups[i].Click("left", 1)
	// 				if err != nil {
	// 					log.Println("Error selecting suggested group:", err)
	// 					return
	// 				}
	// 			}

	// 		}

	// 		// post ad
	// 		page.MustElement(`div[aria-label="Post"]`).MustClick()

	// 		log.Printf("%s posted successfully", title)

	// 		time.Sleep(1 * time.Minute)
	// 	}
	// 	// break after one group
	// }

	// VERSION 2
	var anchors []rod.Element
	seen := make(map[string]struct{}) // To track unique anchors

	for i := 0; i < 100; i++ {
		page.Mouse.MustScroll(0, 500)

		time.Sleep(10 * time.Second)

		classSelector := ".x9f619.x1gryazu.xkrivgy.x1ikqzku.x1h0ha7o.xg83lxy.xh8yej3"
		parents := page.MustElements(classSelector)

		var parent *rod.Element
		parentsLength := len(parents)

		if parentsLength == 1 {
			parent = parents[0]
		} else if parentsLength == 2 {
			parent = parents[1]
		} else {
			fmt.Println("no parents found")
			return
		}

		selector := ".x1i10hfl.x1qjc9v5.xjbqb8w.xjqpnuy.xa49m3k.xqeqjp1.x2hbi6w.x13fuv20.xu3j5b3.x1q0q8m5.x26u7qi.x972fbf.xcfux6l.x1qhh985.xm0m39n.x9f619.x1ypdohk.xdl72j9.x2lah0s.xe8uvvx.xdj266r.x11i5rnm.xat24cr.x1mh8g0r.x2lwn1j.xeuugli.xexx8yu.x4uap5.x18d9i69.xkhd6sd.x1n2onr6.x16tdsg8.x1hl2dhg.xggy1nq.x1ja2u2z.x1t137rt.x1o1ewxj.x3x9cwd.x1e5q0jg.x13rtm0m.x1q0g3np.x87ps6o.x1lku1pv.x1rg5ohu.x1a2a7pz.xh8yej3"
		links := parent.MustElements(selector)

		for _, link := range links {
			anchorIDPtr := link.MustAttribute("href") // Get the pointer to the attribute
			if anchorIDPtr != nil {                   // Check if the attribute is not nil
				anchorID := *anchorIDPtr // Dereference the pointer
				if _, exists := seen[anchorID]; !exists {
					seen[anchorID] = struct{}{}      // Mark this anchor as seen
					anchors = append(anchors, *link) // Dereference link to append
				}
			}
		}
	}

	fmt.Println("Total Unique Anchors:", len(anchors))

	// Shuffle the anchors slice
	r.Shuffle(len(anchors), func(i, j int) {
		anchors[i], anchors[j] = anchors[j], anchors[i]
	})

	var hrefs []string

	for _, anchor := range anchors {
		hrefPtr := anchor.MustAttribute("href")
		if hrefPtr != nil {
			hrefs = append(hrefs, *hrefPtr) // Add hrefs to the slice
		}
	}

	// Shuffle the hrefs slice
	r.Shuffle(len(hrefs), func(i, j int) {
		hrefs[i], hrefs[j] = hrefs[j], hrefs[i]
	})

	// Ensure we have at least 10 hrefs to avoid runtime errors
	if len(hrefs) >= 20 {
		hrefs = hrefs[:20]
	}

	// Now `hrefs` contains up to 20 random unique links

	for _, entry := range entries {

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
				if filepath.Ext(filePath) != ".txt" { // Check if the file is not a .txt file
					fmt.Println("IMAGE:", filePath)

					// Collect file paths to attach
					imageFiles = append(imageFiles, filePath)
				}
			}
		}

		// Open the details.txt file within the subdirectory
		detailsFile := filepath.Join(subDir, "details.txt")
		file, err := os.Open(detailsFile)
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}
		defer file.Close()

		// Initialize variables to hold the extracted fields
		var title, price, description string

		// Create a new scanner to read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			switch {
			case strings.HasPrefix(line, "title:"):
				title = strings.TrimSpace(line[len("title:"):])
			case strings.HasPrefix(line, "price:"):
				price = strings.TrimSpace(line[len("price:"):])
			case strings.HasPrefix(line, "description:"):
				description = strings.TrimSpace(line[len("description:"):])
			}
		}

		// Check for errors during scanning
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// Print the extracted fields
		fmt.Println("Title:", title)
		fmt.Println("Price:", price)

		// Split the text by "..."
		parts := strings.Split(description, "...")

		// Trim spaces from each part
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}

		// Join the parts with "...\n"
		description = strings.Join(parts, "\n\n")

		fmt.Println("Description: " + description)

		for _, href := range hrefs {
			fmt.Println(href)
			page = page.MustNavigate(href).MustWaitLoad().MustWaitDOMStable()
			page, err = utils.CheckSellerLoginStatus(browser, page)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Checked login status")

			if page.MustHas(`div.x1yztbdb.x1j9u4d2`) {
				fmt.Println("Blocked from accessing this group")
				continue
			}

			pageHasSellBtn, sellBtn, _ := page.Has(`div[aria-label="Sell Something"]`)

			switch {
			case pageHasSellBtn:
				fmt.Println("This is a buy and sell group")
				fmt.Println("Handle a case where page has sell button")
				sellBtn.MustScreenshot("sellButton.png")
				continue
			default:
				fmt.Println("This is a normal group")

				page.MustElement(`div.x1i10hfl.x1ejq31n.xd10rxx.x1sy0etr.x17r0tee.x972fbf.xcfux6l.x1qhh985.xm0m39n.x9f619.x1ypdohk.xe8uvvx.xdj266r.x11i5rnm.xat24cr.x1mh8g0r.x16tdsg8.x1hl2dhg.xggy1nq.x87ps6o.x1lku1pv.x1a2a7pz.x6s0dn4.xmjcpbm.x107yiy2.xv8uw2v.x1tfwpuw.x2g32xy.x78zum5.x1q0g3np.x1iyjqo2.x1nhvcw1.x1n2onr6.xt7dq6l.x1ba4aug.x1y1aw1k.xn6708d.xwib8y2.x1ye3gou > div.xi81zsa.x1lkfr7t.xkjl1po.x1mzt3pk.xh8yej3.x13faqbe > span.x1lliihq.x6ikm8r.x10wlt62.x1n2onr6`).MustClick()

				time.Sleep(10 * time.Second)

				dialog := page.MustElement(`div.x1n2onr6.x1ja2u2z.x1afcbsf.x78zum5.xdt5ytf.x1a2a7pz.x6ikm8r.x10wlt62.x71s49j.x1jx94hy.x1qpq9i9.xdney7k.xu5ydu1.xt3gfkd.x104qc98.x1g2kw80.x16n5opg.xl7ujzl.xhkep3z.x193iq5w[role="dialog"]`)

				dialog.MustElement(`div[aria-label="Photo/video"]`).MustClick()

				dialog.MustElement(`input[type="file"]`).MustSetFiles(imageFiles...)

				dialog.MustElement(`div[aria-label="Create a public post…"]`).MustInput(description)

				time.Sleep(5 * time.Second)

				dialog.MustElement(`div[aria-label="Post"]`).MustClick()

				time.Sleep(1 * time.Minute)

				page.MustScreenshot("home.png")

				// wait for 5 minutes before posting item in another group
				time.Sleep(1 * time.Minute)

				continue
			}

		}

		// wait for 5 minutes before posting another item
		time.Sleep(5 * time.Minute)
	}
}
