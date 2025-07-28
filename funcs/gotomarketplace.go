package funcs

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func GoToListSingleItemPage() *rod.Page {
	dir := "~/.config/google-chrome"

	u := launcher.New().UserDataDir(dir).Leakless(true).NoSandbox(true).Headless(true).MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	defer browser.MustClose()

	page := browser.MustPage("https://web.facebook.com/").MustWaitLoad()

	page.MustScreenshot("home.png")

	shortcuts, err := page.Element(`div[aria-label="Shortcuts"]`)
	if err != nil {
		fmt.Println("Error getting shortcuts card:", err)
		return nil
	}

	links, err := shortcuts.Elements(`div[data-visualcompletion="ignore-dynamic"]`)
	if err != nil {
		fmt.Println("Error getting links:", err)
		return nil
	}

	var marketplace *rod.Element

	for _, link := range links {
		if link.MustText() == "Marketplace" {
			marketplace = link
		}
	}

	err = marketplace.Click("left", 1)
	if err != nil {
		fmt.Println("Error clicking link:", err)
		return nil
	}

	time.Sleep(5 * time.Second)

	button, err := page.Element(`a[aria-label="Create new listing"]`)
	if err != nil {
		fmt.Println("Error getting create new listing button:", err)
		return nil
	}

	err = button.Click("left", 1)
	if err != nil {
		fmt.Println("Error clicking create new listing button:", err)
		return nil
	}

	time.Sleep(5 * time.Second)

	links, err = page.Elements(`div.x78zum5.xmrbpvb.x1k70j0n.x1w0mnb.xzueoph.x1mnrxsn.xfj1xho.xlyoakc.x175wuz2`)
	if err != nil {
		fmt.Println("Error getting listing type cards:", err)
		return nil
	}

	var listSingleItemLink *rod.Element

	for _, link := range links {
		if link.MustText() == "Item for sale\nCreate a single listing for one or more items to sell." {
			listSingleItemLink = link
		}
	}

	err = listSingleItemLink.Click("left", 1)
	if err != nil {
		fmt.Println("Error clicking list single item link:", err)
		return nil
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Reached list single item page")

	return page
}
