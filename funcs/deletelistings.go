package funcs

import (
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func DeleteListings() {

	u := launcher.New().UserDataDir("~/.config/google-chrome").Leakless(true).NoSandbox(true).Headless(true).MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	defer browser.MustClose()

	page := browser.MustPage("https://web.facebook.com/marketplace/you/selling").MustWaitLoad()

	for i := 0; i < 100; i++ {
		page.MustReload()

		err := page.Mouse.Scroll(0, 1000, 0)
		if err != nil {
			log.Println("Error scrolling page:", err)
			return
		}

		time.Sleep(10 * time.Second)

		fmt.Println("fetching cards...")

		cards := page.MustElements(`div.x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x2lah0s.x193iq5w.x1k70j0n.xzueoph.xzboxd6.x14l7nz5 > div.x78zum5.x1n2onr6.xh8yej3 > div.x9f619.x1n2onr6.x1ja2u2z.x1jx94hy.x1qpq9i9.xdney7k.xu5ydu1.xt3gfkd.xh8yej3.x6ikm8r.x10wlt62.xquyuld`)

		fmt.Println("cards fetched")

		totalCards := len(cards)

		fmt.Println("Total cards:", totalCards)

		if totalCards > 0 {
			for _, card := range cards {
				fmt.Println("product deletion started...")

				pageHaseCloseChatBtn, _, err := page.Has(`div[aria-label="Close chat"]`)
				if err != nil {
					fmt.Println("Error checking if page has close chat button")
					return
				}

				if pageHaseCloseChatBtn {
					closeChatBtns := page.MustElements(`div[aria-label="Close chat"]`)
					for _, closeChatBtn := range closeChatBtns {
						closeChatBtn.MustClick()
					}
					time.Sleep(10 * time.Second)
				}

				cardHasInfoIcon, infoIcon, err := card.Has(`div[aria-label="The number of times people viewed the details page of your Marketplace listing in the last 14 days."][role="tooltip"]`)
				if err != nil {
					fmt.Println("error checking if card has info icon")
					return
				}

				switch {
				case cardHasInfoIcon:

					infoIcon.MustClick()

					fmt.Println("card info icon clicked")

					time.Sleep(10 * time.Second)

					deleteIcon := page.MustElement(`div[aria-label="Delete"]`)

					deleteIcon.MustClick()

					fmt.Println("delete icon clicked")

					time.Sleep(3 * time.Second)

					deleteButtons := page.MustElements(`div[aria-label="Delete"]`)

					totalDeleteButtons := len(deleteButtons)

					switch {
					case totalDeleteButtons == 3:
						confirmDelete := deleteButtons[2]

						confirmDelete.MustClick()

						fmt.Println("confirm button clicked")

						time.Sleep(10 * time.Second)
					default:

						confirmDelete := deleteButtons[1]

						confirmDelete.MustClick()

						fmt.Println("confirm button clicked")

						time.Sleep(10 * time.Second)
					}

				default:
					card.MustClick()

					fmt.Println("card clicked")

					time.Sleep(10 * time.Second)

					page.MustScreenshot("facebook.png")

					deleteIcon := page.MustElement(`div[aria-label="Delete"]`)

					deleteIcon.MustClick()

					fmt.Println("delete icon clicked")

					time.Sleep(3 * time.Second)

					deleteButtons := page.MustElements(`div[aria-label="Delete"]`)

					totalDeleteButtons := len(deleteButtons)

					switch {
					case totalDeleteButtons == 3:
						confirmDelete := deleteButtons[2]

						confirmDelete.MustClick()

						fmt.Println("confirm button clicked")

						time.Sleep(10 * time.Second)
					default:

						confirmDelete := deleteButtons[1]

						confirmDelete.MustClick()

						fmt.Println("confirm button clicked")

						time.Sleep(10 * time.Second)
					}

				}

				fmt.Println("product deleted successfully!")
			}
		} else {
			break
		}
	}
}
