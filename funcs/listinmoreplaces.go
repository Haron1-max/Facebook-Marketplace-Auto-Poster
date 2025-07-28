package funcs

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/go-rod/rod"
)

func ListInMorePlaces(browser *rod.Browser, page *rod.Page) {

	defer browser.MustClose()

	page = page.MustNavigate("https://web.facebook.com/marketplace/you/selling").MustWaitLoad().MustWaitDOMStable()

	page.MustScreenshot("home.png")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	memberCountRegex := regexp.MustCompile(`(\d+(\.\d+)?)([KM]?)\s*members`)

	pageHasCollection, collection, _ := page.Has(`div[aria-label="Collection of your Marketplace items"]`)

	if !pageHasCollection {
		fmt.Println("Page has not collection of marketplace items")
		return
	}

	needsAttention := collection.MustElement(`div.xyamay9.x13mt7qq.x1w9j1nh.x1rik9be.x1mtsufr > div.x1n2onr6.x1ja2u2z.x9f619.x78zum5.xdt5ytf.x2lah0s.x193iq5w > div.x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x1iyjqo2.x2lwn1j`)

	needsAttention.MustScreenshot("needsAttention.png")

	allListings := collection.MustElement(`span.x1lliihq.x1iyjqo2 > div.x1n2onr6.x1ja2u2z.x9f619.x78zum5.xdt5ytf.x2lah0s.x193iq5w > div.x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x1iyjqo2.x2lwn1j`)

	allListings.MustScreenshot("allListings.png")

	for i := 0; i < 100; i++ {

		page.Mouse.MustScroll(0, 1000)

		time.Sleep(10 * time.Second)

		fmt.Println("fetching cards...")

		// alternate between needs attention and all listings
		cards := needsAttention.MustElements(`div.x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x2lah0s.x193iq5w.x1k70j0n.xzueoph.xzboxd6.x14l7nz5 > div.x78zum5.x1n2onr6.xh8yej3 > div.x9f619.x1n2onr6.x1ja2u2z.x1jx94hy.x1qpq9i9.xdney7k.xu5ydu1.xt3gfkd.xh8yej3.x6ikm8r.x10wlt62.xquyuld`)

		fmt.Println("cards fetched")

		totalCards := len(cards)

		fmt.Println("Total cards:", totalCards)

		r.Shuffle(totalCards, func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})

		if totalCards > 0 {
			for _, card := range cards {

				card.MustScreenshot("home.png")

				card.MustScrollIntoView()

				pageHaseCloseChatBtn := page.MustHas(`div[aria-label="Close chat"]`)

				if pageHaseCloseChatBtn {
					closeChatBtns := page.MustElements(`div[aria-label="Close chat"]`)
					for _, closeChatBtn := range closeChatBtns {
						closeChatBtn.MustClick()
					}
					time.Sleep(3 * time.Second)
				}

				cardHasInfoIcon, infoIcon, err := card.Has(`div[aria-label="The number of times that people viewed the details page of your Marketplace listing in the last 14 days."]`)
				if err != nil {
					fmt.Println("error checking if card has info icon")
					return
				}

				selectedGroups := 0

				switch {
				case cardHasInfoIcon:

					fmt.Println("card has info icon")

					infoIcon.MustClick()

					fmt.Println("card info icon clicked")

					time.Sleep(10 * time.Second)

					page.MustScreenshot("home.png")

					pageHasListToMorePlacesBtn, listToMorePlacesBtn, _ := page.Has(`div[aria-label="List to more places"]`)

					if !pageHasListToMorePlacesBtn {
						fmt.Println("Page has no list to more places button")
						continue
					}

					fmt.Println("Page has list to more places button")

					listToMorePlacesBtn.MustClick()

					time.Sleep(10 * time.Second)

					fmt.Println("List to more places button clicked")

					page.MustScreenshot("home.png")

					pageHasDialog, dialog, _ := page.Has(`div[aria-label="List in more places"]`)

					if !pageHasDialog {
						fmt.Println("Page has no list in more places dialog")
						continue
					}

					fmt.Println("Page has dialog")
					dialog.MustScreenshot("home.png")

					suggestedGroupsContainer := dialog.MustElement("div.x1n2onr6.x1ja2u2z.x9f619.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xsag5q8.xexx8yu")

					suggestedGroups := suggestedGroupsContainer.MustElements(`div[data-visualcompletion="ignore-dynamic"][style="padding-left: 8px; padding-right: 8px;"]`)

					totalSuggestedGroups := len(suggestedGroups)

					fmt.Println("Total suggested groups:", totalSuggestedGroups)

					if totalSuggestedGroups > 20 {
						for _, suggestedGroup := range suggestedGroups {
							members := suggestedGroup.MustElement(`div.x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xeuugli.xg83lxy.x1h0ha7o.x1120s5i.x1nn3v0j`).MustText()

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
								log.Println("Error parsing member count:", err)
								continue
							}

							// Adjust for K or M
							if multiplier == "K" {
								memberCount *= 1000
							} else if multiplier == "M" {
								memberCount *= 1000000
							}

							//memberCount = RoundToNearestThousand(memberCount)

							if memberCount >= 10000 {
								suggestedGroup.MustClick().MustScreenshot("home.png")
								selectedGroups++
								if selectedGroups == 20 {
									break
								}
								time.Sleep(1 * time.Second)
							}

							//suggestedGroupsContainer.MustScreenshot("facebook.png")
						}
					} else {
						for _, suggestedGroup := range suggestedGroups {
							suggestedGroup.MustClick().MustScreenshot("home.png")
							selectedGroups++
							if selectedGroups == 20 {
								break
							}
							time.Sleep(1 * time.Second)
						}
					}

					fmt.Println("Selected groups:", selectedGroups)

					dialog.MustElements("div.x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x193iq5w.xeuugli.x1iyjqo2.xs83m0k.x150jy0e.x1e558r4.xjkvuk6.x1iorvi4.xdl72j9")[1].MustElement(`div[aria-label="Post"]`).MustClick()

					fmt.Println("Clicked post button")

					time.Sleep(10 * time.Second)

					page.MustElement(`div[aria-label="Close"]`).MustClick()

					time.Sleep(10 * time.Second)

				default:
					fmt.Println("Card has no info icon")

					card.MustClick()

					fmt.Println("Card clicked")

					time.Sleep(10 * time.Second)

					pageHasListToMorePlacesBtn, listToMorePlacesBtn, _ := page.Has(`div[aria-label="List to more places"]`)

					if !pageHasListToMorePlacesBtn {
						fmt.Println("Page has no list to more places button")
						continue
					}

					listToMorePlacesBtn.MustClick()

					time.Sleep(10 * time.Second)

					fmt.Println("List to more places button clicked")

					pageHasDialog, dialog, _ := page.Has(`div[aria-label="List in more places"]`)

					if !pageHasDialog {
						fmt.Println("Page has no list in more places dialog")
						continue
					}

					fmt.Println("Page has dialog")
					dialog.MustScreenshot("home.png")

					suggestedGroupsContainer := dialog.MustElement("div.x1n2onr6.x1ja2u2z.x9f619.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xsag5q8.xexx8yu")

					suggestedGroups := suggestedGroupsContainer.MustElements(`div[data-visualcompletion="ignore-dynamic"][style="padding-left: 8px; padding-right: 8px;"]`)

					totalSuggestedGroups := len(suggestedGroups)

					fmt.Println("Total suggested groups:", totalSuggestedGroups)

					fmt.Println("past here")

					if totalSuggestedGroups > 20 {
						for _, suggestedGroup := range suggestedGroups {
							members := suggestedGroup.MustElement(`div.x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xeuugli.xg83lxy.x1h0ha7o.x1120s5i.x1nn3v0j`).MustText()

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
								log.Println("Error parsing member count:", err)
								continue
							}

							// Adjust for K or M
							if multiplier == "K" {
								memberCount *= 1000
							} else if multiplier == "M" {
								memberCount *= 1000000
							}

							//memberCount = RoundToNearestThousand(memberCount)

							if memberCount >= 10000 {
								suggestedGroup.MustClick()
								selectedGroups++
								if selectedGroups == 20 {
									break
								}
								time.Sleep(1 * time.Second)
							}

							//suggestedGroupsContainer.MustScreenshot("facebook.png")
						}
					} else {
						for _, suggestedGroup := range suggestedGroups {
							suggestedGroup.MustClick().MustScreenshot("home.png")
							selectedGroups++
							if selectedGroups == 20 {
								break
							}
							time.Sleep(1 * time.Second)
						}
					}

					dialog.MustElements("div.x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x193iq5w.xeuugli.x1iyjqo2.xs83m0k.x150jy0e.x1e558r4.xjkvuk6.x1iorvi4.xdl72j9")[1].MustElement(`div[aria-label="Post"]`).MustClick()

					time.Sleep(10 * time.Second)

					page.MustElement(`div[aria-label="Close"]`).MustClick()

					time.Sleep(10 * time.Second)

				}

				fmt.Println("Item shared successfully!")

				page.MustScreenshot("home.png")

				time.Sleep(60 * time.Second)
			}
		} else {
			break
		}
	}
}
