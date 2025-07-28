package funcs

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/haron1996/fb/utils"
)

func FirstReply() {
	u := launcher.New().UserDataDir("~/.config/google-chrome").Leakless(true).NoSandbox(true).Headless(true).MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	defer browser.MustClose()

	page := browser.MustPage("https://web.facebook.com/messages/t/").MustWaitLoad()

	time.Sleep(5 * time.Second)

	chats := page.MustElement(`div[aria-label="Chats"]`)

	marketplaceChats := chats.MustElement(`div.xurb0ha.x1sxyh0[role="gridcell"]`)

	if err := marketplaceChats.Click("left", 1); err != nil {
		fmt.Println("Error clicking on marketplace chats:", err)
		return
	}

	time.Sleep(5 * time.Second)

	marketplace := page.MustElement(`div[aria-label="Marketplace"]`)

	var hrefs []string

	fmt.Println("Loading messages...")

	for i := 0; i < 30; i++ {
		marketplace.Page().Mouse.MustScroll(0, 500)
		time.Sleep(5 * time.Second)
		messages := marketplace.MustElements(`div.x78zum5.xdt5ytf[data-virtualized="false"]`)
		for _, message := range messages {
			a := message.MustElement(`a[role="link"]`)
			href := *a.MustAttribute("href")
			hrefs = append(hrefs, href)
		}
	}

	uniqueHrefs := utils.RemoveDuplicateURLs(hrefs)

	fmt.Println("Unique Hrefs:", len(uniqueHrefs))

	reply := `
              Its still available.

Are you still interested?

We are located at Pioneer Building, Kimathi Street.

We also deliver within Nairobi and its environs.

Free delivery. Pay on delivery.

If you're still interested, reply to this thread, call, or WhatsApp me at 0718448461.

Have a great day!

Regards,
Haron.

                `

	for _, href := range uniqueHrefs {
		threadUrl := fmt.Sprintf("https://web.facebook.com%s", href)

		fmt.Println(threadUrl)

		page := browser.MustPage(threadUrl).MustWaitLoad()

		time.Sleep(5 * time.Second)

		myMessages := page.MustElements(`div.html-div.xe8uvvx.xexx8yu.x4uap5.x18d9i69.xkhd6sd.x1gslohp.x11i5rnm.x12nagc.x1mh8g0r.x1yc453h.x126k92a.xyk4ms5`)

		if len(myMessages) > 0 {
			fmt.Println("Thread has my messages")
			fmt.Println("Length of my messages in thread:", len(myMessages))
			continue
		}

		leadMessages := page.MustElements(`div.html-div.xe8uvvx.xexx8yu.x4uap5.x18d9i69.xkhd6sd.x1gslohp.x11i5rnm.x12nagc.x1mh8g0r.x1yc453h.x126k92a.x18lvrbx`)

		lenLeadMessages := len(leadMessages)

		if lenLeadMessages == 0 {
			fmt.Println("No lead messages found")
			continue
		}

		input := page.MustElement(`div[aria-placeholder="Aa"]`)

		msg := reply

		input.MustInput(msg)

		time.Sleep(5 * time.Second)

		sendBtn := page.MustElement(`div[aria-label="Press Enter to send"]`)

		sendBtn.MustClick()

		time.Sleep(5 * time.Second)

		fmt.Println("OK")
	}

}
