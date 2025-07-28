package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func GoToListingPage(page *rod.Page, listingType string) (*rod.Page, error) {

	pageHasShortcuts, shortcuts, err := page.Has(`div[aria-label="Shortcuts"]`)
	if err != nil {
		return page, fmt.Errorf("error getting shortcuts div: %v", err)
	}

	switch {
	case pageHasShortcuts:
		links := shortcuts.MustElements(`div[data-visualcompletion="ignore-dynamic"]`)

		var marketplaceLink *rod.Element

		for _, link := range links {
			if link.MustText() == "Marketplace" {
				marketplaceLink = link
			}
		}

		marketplaceLink.MustClick()

		time.Sleep(3 * time.Second)

		button := page.MustElement(`a[aria-label="Create new listing"]`)

		button.MustClick()

		time.Sleep(3 * time.Second)

		pageHaseCloseChatBtn, _, _ := page.Has(`div[aria-label="Close chat"]`)

		if pageHaseCloseChatBtn {
			closeChatBtns := page.MustElements(`div[aria-label="Close chat"]`)
			for _, closeChatBtn := range closeChatBtns {
				closeChatBtn.MustClick()
			}
			time.Sleep(2 * time.Second)
		}

		links = page.MustElements(`div.x78zum5.xmrbpvb.x1k70j0n.x1w0mnb.xzueoph.x1mnrxsn.xfj1xho.xlyoakc.x175wuz2`)

		switch {
		case listingType == "Item for sale":
			fmt.Println("Go to item for sale listing page")
			links[0].MustClick()
		case listingType == "Vehicle for sale":
			fmt.Println("Go to vehicle for sale listing page")
			links[1].MustClick()
		case listingType == "Property for sale or rent":
			fmt.Println("Go to property for sale or rent listing page")
			links[2].MustClick()
		default:
			return page, errors.New("invalid listing type")
		}
	default:
		button := page.MustElement(`a[aria-label="Create new listing"]`)

		button.MustClick()

		time.Sleep(5 * time.Second)

		pageHaseCloseChatBtn, _, _ := page.Has(`div[aria-label="Close chat"]`)

		if pageHaseCloseChatBtn {
			closeChatBtns := page.MustElements(`div[aria-label="Close chat"]`)
			for _, closeChatBtn := range closeChatBtns {
				closeChatBtn.MustClick()
			}
			time.Sleep(2 * time.Second)
		}

		links := page.MustElements(`div.x78zum5.xmrbpvb.x1k70j0n.x1w0mnb.xzueoph.x1mnrxsn.xfj1xho.xlyoakc.x175wuz2`)

		switch {
		case listingType == "Item for sale":
			fmt.Println("Go to item for sale listing page")
			links[0].MustClick()
		case listingType == "Vehicle for sale":
			fmt.Println("Go to vehicle for sale listing page")
			links[1].MustClick()
		case listingType == "Property for sale or rent":
			fmt.Println("Go to property for sale or rent listing page")
			links[2].MustClick()
		default:
			return page, errors.New("invalid listing type")
		}
	}

	return page, nil
}
