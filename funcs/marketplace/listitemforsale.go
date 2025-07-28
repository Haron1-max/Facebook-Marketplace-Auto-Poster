package marketplace

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/go-rod/rod"
	"github.com/haron1996/fb/config"
	"github.com/haron1996/fb/utils"
)

// Create a single listing for one or more items to sell
func ListItemForSale(browser *rod.Browser, page *rod.Page) {
	startTime := time.Now()
	// Ensure the browser closes when main function exits
	defer browser.MustClose()

	// load config files
	config, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// Root directory containing subdirectories with images
	root := config.Root

	// get all directories in root
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

		page.MustScreenshot("fb.png")

		page, err := utils.CheckSellerLoginStatus(browser, page)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Checked seller login status")

		page = page.MustNavigate("https://www.facebook.com/marketplace/create/item").MustWaitLoad().MustWaitDOMStable()

		//page.MustScreenshot("marketplace.png")

		fmt.Println("Navigated to create item link")

		//os.Exit(1)

		// Path to the current subdirectory
		subDir := filepath.Join(root, entry.Name())

		fmt.Println("DIRECTORY:", subDir)

		imageFiles, err := utils.ReturnEntryImages(subDir, r)
		if err != nil {
			fmt.Println(err)
			return
		}

		item, err := utils.ReturnItemDetails(subDir, page, r, imageFiles)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = item.ListItemForSale()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("🥳🥳🥳🥳🥳🥳🥳🥳🥳🥳🥳🥳🥳🥳🥳🥳🥳")

		time.Sleep(1 * time.Minute)
	}

	finishTime := time.Since(startTime).Minutes()

	fmt.Printf("%d products in %s have been listed in %f minutes. Finish time: %s\n", len(entries), root, finishTime, time.Now().Local())
}
