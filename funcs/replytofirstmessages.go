package funcs

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func ReplyToFirstMessages() {
	u := launcher.New().UserDataDir("~/.config/google-chrome").Leakless(true).NoSandbox(true).Headless(true).MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	defer browser.MustClose()

	page := browser.MustPage("https://web.facebook.com/messages/t/").MustWaitLoad()

	time.Sleep(10 * time.Second)

	fmt.Println("processing...")

	chats := page.MustElement(`div[aria-label="Chats"]`)

	marketplaceChats := chats.MustElement(`div.xurb0ha.x1sxyh0[role="gridcell"]`)

	if err := marketplaceChats.Click("left", 1); err != nil {
		fmt.Println("Error clicking on marketplace chats:", err)
		return
	}

	fmt.Println("processing...")

	time.Sleep(10 * time.Second)

	marketplace := page.MustElement(`div[aria-label="Marketplace"]`)

	fmt.Println("processing...")

	err := marketplace.Page().Mouse.Scroll(0, 50000, 10000)
	if err != nil {
		fmt.Println("Error scrolling marketplace chats:", err)
		return
	}

	fmt.Println("DONE SCROLLING")
	page.MustScreenshot("fb.png")

	fmt.Println(len(marketplace.MustElements(`a.x1i10hfl.x1qjc9v5.xjbqb8w.xjqpnuy.xa49m3k.xqeqjp1.x2hbi6w.x13fuv20.xu3j5b3.x1q0q8m5.x26u7qi.x972fbf.xcfux6l.x1qhh985.xm0m39n.x9f619.x1ypdohk.xdl72j9.x2lah0s.xe8uvvx.xdj266r.x11i5rnm.xat24cr.x1mh8g0r.x2lwn1j.xeuugli.xexx8yu.x4uap5.x18d9i69.xkhd6sd.x1n2onr6.x16tdsg8.x1hl2dhg.xggy1nq.x1ja2u2z.x1t137rt.x1o1ewxj.x3x9cwd.x1e5q0jg.x13rtm0m.x1q0g3np.x87ps6o.x1lku1pv.x1a2a7pz.x1lliihq`)))

	messages := marketplace.MustElements(`div.x78zum5.xdt5ytf`)

	fmt.Println("Messages found:", len(messages))

	//return

	for _, message := range messages {
		fmt.Println("Looping messages...")
		page.MustScreenshot("facebook.png")
		message.MustHover()
		fmt.Println("messanger hovering...")
		time.Sleep(5 * time.Second)
		fmt.Println("messanger hovered...")
		menu := message.MustElement(`div[aria-label="Menu"]`)
		menu.MustScrollIntoView()
		menu.MustClick()
		time.Sleep(5 * time.Second)
		fmt.Println("menu clicked...")

		optionsCard := page.MustElement(`div[aria-label="More options for this chat"]`)
		options := optionsCard.MustElements(`div.x1i10hfl.xjbqb8w.x1ejq31n.xd10rxx.x1sy0etr.x17r0tee.x972fbf.xcfux6l.x1qhh985.xm0m39n.xe8uvvx.x1hl2dhg.xggy1nq.x1o1ewxj.x3x9cwd.x1e5q0jg.x13rtm0m.x87ps6o.x1lku1pv.x1a2a7pz.xjyslct.x9f619.x1ypdohk.x78zum5.x1q0g3np.x2lah0s.x1i6fsjq.xfvfia3.xnqzcj9.x1gh759c.x1n2onr6.x16tdsg8.x1ja2u2z.x6s0dn4.x1y1aw1k.xwib8y2.x1q8cg2c.xnjli0[role="menuitem"]`)

		fmt.Println("Len options:", len(options))

		for _, option := range options {
			fmt.Println("started options loop")
			fmt.Println("reached here")
			fmt.Println(option.MustText())
			if option.MustText() == "Mark as read" {
				fmt.Println("message is unread")
				a := message.MustElement(`a[role="link"]`)
				fmt.Println("past a")
				href := *a.MustAttribute("href")
				fmt.Println("past href")
				threadUrl := fmt.Sprintf("https://web.facebook.com%s", href)
				fmt.Println(threadUrl)
				page := browser.MustPage(threadUrl).MustWaitLoad()
				fmt.Println("past new page")
				time.Sleep(10 * time.Second)
				fmt.Println("New tab screenshot taken...")
				fmt.Println("weri chorwenyun....")
				messages := page.MustElements(`div.html-div.xe8uvvx.xexx8yu.x4uap5.x18d9i69.xkhd6sd.x1gslohp.x11i5rnm.x12nagc.x1mh8g0r.x1yc453h.x126k92a.x18lvrbx[dir="auto"]`)
				fmt.Println("past getting messages...")
				fmt.Println("Len of found messages:", len(messages))
				fmt.Println("LAST STEP BRO")
				lastMessage := messages[len(messages)-1]
				lastMessageText := lastMessage.MustText()
				fmt.Println(lastMessageText)
				fmt.Println("past last text message text")
				keywords := []string{"hujambo", "available", "habari"}
				fmt.Println("past keywords...")
				for _, kw := range keywords {
					fmt.Println("STARTED LOOPING KEYWORDS")
					fmt.Println("KEYWORD:", kw)
					if strings.Contains(strings.ToLower(lastMessageText), kw) {
						fmt.Println("Last message contains key word:", kw)
						fmt.Println(lastMessageText)
						input := page.MustElement(`div[aria-placeholder="Aa"]`)
						reply := `
						Hi.

It's available.

Nairobi and its environs only.

Free delivery. Pay on Delivery.

For more info, reply, call, or WhatsApp 0718448461.
						`
						input.MustInput(reply)
						time.Sleep(5 * time.Second)
						sendBtn := page.MustElement(`div[aria-label="Press Enter to send"]`)
						sendBtn.MustClick()
						time.Sleep(5 * time.Second)
						page.MustScreenshot("facebook.png")
						fmt.Println("MESSAGE REPLIED")
					}
				}

				menu.MustClick()
				time.Sleep(5 * time.Second)
			}
		}

	}

	fmt.Println("done")
}
