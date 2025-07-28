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
	"github.com/haron1996/fb/cookies"
	"github.com/haron1996/fb/utils"
)

func CreatePost(browser *rod.Browser, page *rod.Page) {
	defer browser.MustClose()

	config, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	root := config.Root

	entries, err := os.ReadDir(root)
	if err != nil {
		fmt.Println("Error reading root directory:", err)
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Shuffle(len(entries), func(i, j int) {
		entries[i], entries[j] = entries[j], entries[i]
	})

	for _, entry := range entries {
		// check seller login status
		page, err := utils.CheckSellerLoginStatus(browser, page)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Checked login status")

		currentURL := page.MustInfo().URL
		profileURL := fmt.Sprintf("https://www.facebook.com/profile.php?id=%s", cookies.C_user)

		switch {
		case currentURL != profileURL:
			profileURL := fmt.Sprintf("https://www.facebook.com/profile.php?id=%s", cookies.C_user)

			page = page.MustNavigate(profileURL).MustWaitLoad().MustWaitDOMStable()
		}

		fmt.Println("Reached profile page")

		container := page.MustElement(`div.xqmpxtq.x13fuv20.x178xt8z.x78zum5.x1a02dak.x1vqgdyp.x1l1ennw.x14vqqas.x6ikm8r.x10wlt62.x1y1aw1k`)

		buttons := container.MustElements(`div[role="button"]`)

		if len(buttons) == 0 {
			fmt.Println("Buttons not found")
			return
		}

		for _, button := range buttons {
			if button.MustText() == "Photo/video" {
				button.MustClick()
			}
		}

		time.Sleep(10 * time.Second)

		page.MustScreenshot("fb.png")

		pageHasDialog, dialog, err := page.Has(`div[role="dialog"]`)
		if err != nil {
			log.Println("Error checkin if pages has dialog:", err)
			return
		}

		if !pageHasDialog {
			log.Println("Page ha no dialog:", err)
			return
		}

		dialogHasFileInput, fileInput, err := dialog.Has(`input[type="file"]`)
		if err != nil {
			log.Println("Error checking if dialog has file input:", err)
			return
		}

		if !dialogHasFileInput {
			log.Println("Dialog has no file input")
			return
		}

		subDir := filepath.Join(root, entry.Name())
		fmt.Println("DIRECTORY:", subDir)

		subEntries, err := os.ReadDir(subDir)
		if err != nil {
			fmt.Println("Error reading subdirectory:", err)
			continue
		}

		var imageFiles []string

		for _, subEntry := range subEntries {
			if !subEntry.IsDir() {
				filePath := filepath.Join(subDir, subEntry.Name())
				if filepath.Ext(filePath) != ".txt" {
					fmt.Println("IMAGE:", filePath)
					imageFiles = append(imageFiles, filePath)
				}
			}
		}

		if len(imageFiles) > 0 {
			fileInput.MustSetFiles(imageFiles...)
		}

		detailsFile := filepath.Join(subDir, "details.txt")
		file, err := os.Open(detailsFile)
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}
		defer file.Close()

		var title, description string

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			switch {
			case strings.HasPrefix(line, "title:"):
				title = strings.TrimSpace(line[len("title:"):])
			case strings.HasPrefix(line, "description:"):
				description = strings.TrimSpace(line[len("description:"):])
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		parts := strings.Split(description, "...")

		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}

		formattedDescription := strings.Join(parts, "\n\n")

		fmt.Println("Description: " + formattedDescription)

		contentEditableDiv := dialog.MustElement(`div[aria-label="What's on your mind?"]`)

		contentEditableDiv.MustInput(formattedDescription)

		postBtn := dialog.MustElement(`div[aria-label="Post"]`)

		time.Sleep(30 * time.Second)

		postBtn.MustClick()

		time.Sleep(30 * time.Second)

		fmt.Printf("%s posted 🥳\n", title)
	}
}
