package funcs

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/haron1996/fb/config"
	"github.com/haron1996/fb/utils"
)

func PostToCommunities() {
	config, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	u := launcher.New().UserDataDir("~/.config/google-chrome").Leakless(true).NoSandbox(true).Headless(true).MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	defer browser.MustClose()

	page := browser.MustPage("https://web.facebook.com/messages/t/").MustWaitLoad()

	time.Sleep(10 * time.Second)

	spans := page.MustElements(`span.x1lliihq.x6ikm8r.x10wlt62.x1n2onr6.xlyipyv.xuxw1ft`)

	var communitiesSpan *rod.Element

	for _, span := range spans {
		if span.MustText() == "Communities" {
			communitiesSpan = span
		}
	}

	err = communitiesSpan.Click("left", 1)
	if err != nil {
		fmt.Println("Error clicking communities span")
		return
	}

	time.Sleep(10 * time.Second)

	chats := page.MustElement(`div[aria-label="Chats"]`)

	chatsHasSeeAllBtn, seeAllBtn, err := chats.Has(`div[aria-label="See all"]`)
	if err != nil {
		fmt.Println("Error checking if page has see all button")
		return
	}

	if chatsHasSeeAllBtn {
		seeAllBtn.MustClick()
	}

	time.Sleep(10 * time.Second)

	var hrefs []string

	for i := 0; i < 5; i++ {
		chats.Page().Mouse.MustScroll(0, 1000)
		communities := chats.MustElements(`div.html-div.xdj266r.x11i5rnm.xat24cr.x1mh8g0r.xexx8yu.x18d9i69.xurb0ha.x1sxyh0.x1n2onr6`)
		for _, community := range communities {
			a := community.MustElement(`a[role="link"]`)
			href := *a.MustAttribute("href")
			hrefs = append(hrefs, href)
		}
	}

	hrefs = utils.RemoveDuplicateURLs(hrefs)

	//fmt.Println(hrefs[len(hrefs)-1]) // gets the last href
	//fmt.Println(hrefs[0])

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

	for _, href := range hrefs {
		// visit each community
		// join and post if not already joined
		// post if joined
		href = fmt.Sprintf(`https://www.facebook.com%s`, href)

		fmt.Println(href)

		page := browser.MustPage(href).MustWaitLoad()

		//time.Sleep(30 * time.Second)

		pageHasJoin, join, err := page.Has(`div[aria-label="Join"]`)
		if err != nil {
			fmt.Println("Error checking if page has join button")
			continue
		}

		if pageHasJoin {
			fmt.Println("Page has join button")

			err := join.Click("left", 1)
			if err != nil {
				fmt.Println("Error clicking join button:", err)
				continue
			}

			time.Sleep(10 * time.Second)

			pageHasOkBtn, oKBtn, err := page.Has(`div[aria-label="OK"]`)
			if err != nil {
				fmt.Println("Error checking if page has OK button:", err)
				continue
			}

			if pageHasOkBtn {
				fmt.Println("Page has OK button")

				err := oKBtn.Click("left", 1)
				if err != nil {
					fmt.Println("Error clicking ok button:", err)
					continue
				}
			}
		}

		time.Sleep(10 * time.Second)

		fmt.Println("Post to community")

		for _, entry := range entries {

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
					if filepath.Ext(filePath) != ".txt" { // Check if the file is not a .txt file
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
			var description, category string

			// Create a new scanner to read the file line by line
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()

				switch {
				case strings.HasPrefix(line, "description:"):
					description = strings.TrimSpace(line[len("description:"):])
				case strings.HasPrefix(line, "category"):
					category = strings.ToLower(strings.TrimSpace(line[len("category:"):]))
				}
			}

			// Check for errors during scanning
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			formattedDescription := formatDescription(description, category)

			fmt.Println("Description: " + formattedDescription)

			textBox, err := page.Element(`div[aria-placeholder="Aa"][aria-label="Message"]`)
			if err != nil {
				fmt.Println("Error getting text box:", err)
				continue
			}

			err = textBox.Input(formattedDescription)
			if err != nil {
				fmt.Println("Error inputing product description to textbox:", err)
				continue
			}

			if len(imageFiles) > 0 {
				fileInput.MustSetFiles(imageFiles...)
			}

			time.Sleep(5 * time.Second)

			page.MustScreenshot("facebook.png")

			postBtn, err := page.Element(`div[aria-label="Press Enter to send"]`)
			if err != nil {
				fmt.Println("Error getting post button:", err)
				continue
			}

			err = postBtn.Click("left", 1)
			if err != nil {
				fmt.Println("Error clicking post button:", err)
				continue
			}

			time.Sleep(30 * time.Second)
		}

	}

}
