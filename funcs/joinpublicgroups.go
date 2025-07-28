package funcs

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/haron1996/fb/utils"
)

func JoinPublicGroups() {
	estates := []string{
		"Kilimani", "Kileleshwa", "Lavington", "Westlands", "Parklands",
		"Hurlingham", "Upper Hill", "Nairobi",
		"Embakasi", "Donholm", "Umoja", "Buruburu", "Komarock", "Kayole",
		"Tassia", "Pipeline", "Saika", "Ruai", "Utawala",
		"Fedha", "Mukuru kwa Njenga",
		"Kangemi", "Uthiru", "Kinoo", "Kawangware",
		"Riruta", "Kabete", "Dagoretti", "Waithaka", "Limuru",
		"South B", "South C", "Lang’ata", "Karen", "Nyayo Estate", "Nairobi West",
		"Madaraka", "Ngong Road", "Lucky Summer",
		"Runda", "Muthaiga",
		"Kasarani", "Roysambu", "Githurai", "Zimmerman", "Kahawa", "Kahawa West",
		"Kahawa Sukari", "Kahawa Wendani", "Membley", "Marurui", "Garden City",
		"Roysambu",
		"Syokimau", "Athi River", "Kitengela", "Ngong",
		"Mlolongo", "Ruaka", "Thindigua", "Kikuyu", "Ruiru",
		"Kiambu",
		"Kamakis", "Gwa Kairu",
		"Kahawa Sukari", "Kahawa Wendani", "Githurai", "Ruiru Ndani",
		"Thika", "Makongeni", "Ngoingwa", "Landless", "Juja",
		"Witeithie", "Muguga",
		"Kikuyu",
		"Kabete", "Wangige", "Uthiru",
		"Ruaka", "Ndenderu", "Muchatha",
		"Kahawa Sukari", "Kahawa Wendani", "Githurai 44",
		"Gachie",
		"ongata rongai", "kiserian", "Thika Road",
	}

	// estates := []string{
	// 	"Eldoret", "Kisumu", "Mombasa", "Kisii", "Busia", "Kakamega", "Nakuru", "Garissa", "Kilifi", "Malindi", "Kericho", "Siaya", "Kericho", "Kilgoris", "Nyamira", "Meru", "Embu", "Nyeri", "Kirinyaga", "Nairobi", "Narok", "Naivasha", "Gilgil", "Tana River", "Turkana", "Machakos", "Kitui", "Migori", "Bomet", "Bungoma", "Lamu", "Voi", "Kiambu", "Muranga", "Kajiado", "Baringo", "Elgeyo Marakwet", "Homa Bay", "Eldama Ravine", "Isiolo", "Kerugoya", "Kwale", "Laikipia", "Makueni", "Mandera", "Marsabit", "Nandi", "Kitale", "Nyandarua", "Vihiga", "Nyahururu", "Tharakanithi", "Trans-nzoia", "Wajir", "Pokot", "West Pokot", "Samburu", "Malaba",
	// }

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomIndex := r.Intn(len(estates))

	randomEstate := estates[randomIndex]

	fmt.Println("Random Estate:", randomEstate)

	q := strings.ReplaceAll(randomEstate, " ", "%20")

	url := fmt.Sprintf(`https://www.facebook.com/search/groups?q=%s&filters=eyJwdWJsaWNfZ3JvdXBzOjAiOiJ7XCJuYW1lXCI6XCJwdWJsaWNfZ3JvdXBzXCIsXCJhcmdzXCI6XCJcIn0ifQ`, strings.ToLower(q))

	// Launch the browser with specific configurations
	u := launcher.New().UserDataDir("~/.config/google-chrome").Leakless(true).NoSandbox(true).Headless(true).MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	defer browser.MustClose() // Ensure the browser closes when main function exits

	page := browser.MustPage(url).MustWaitLoad()

	var groups []string

	for i := 0; i < 10; i++ {
		page.Mouse.MustScroll(0, 500)

		time.Sleep(5 * time.Second)

		articles := page.MustElements(`div[role="article"]`)

		// Regular expression to capture the member count (e.g., "21K members" or "10K+ members")
		memberCountRegex := regexp.MustCompile(`(\d+(\.\d+)?)([KM]?)\s*members`)

		for _, article := range articles {
			spans := article.MustElements(`span.x1lliihq.x6ikm8r.x10wlt62.x1n2onr6`)

			groupInfo := spans[0].MustText()

			parts := strings.Split(groupInfo, " · ")

			// Find the member count in the text
			match := memberCountRegex.FindStringSubmatch(parts[1])
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

			if len(parts) > 2 {
				numberOfPosts := parts[2]
				if memberCount >= 10000 && numberOfPosts == "10+ posts a day" {
					spans := article.MustElements("span")
					for _, span := range spans {
						text := span.MustText()
						if text == "Join" {
							fmt.Println(groupInfo)
							a := article.MustElement(`a.x1i10hfl.xjbqb8w.x1ejq31n.xd10rxx.x1sy0etr.x17r0tee.x972fbf.xcfux6l.x1qhh985.xm0m39n.x9f619.x1ypdohk.xt0psk2.xe8uvvx.xdj266r.x11i5rnm.xat24cr.x1mh8g0r.xexx8yu.x4uap5.x18d9i69.xkhd6sd.x16tdsg8.x1hl2dhg.xggy1nq.x1a2a7pz.x1sur9pj.xkrqix3.xzsf02u.x1s688f[role="presentation"]`)
							href := *a.MustAttribute("href")
							fmt.Println(href)
							groups = append(groups, href)
						}
					}
				}
			}

		}
	}

	groups = utils.RemoveDuplicateURLs(groups)

	r.Shuffle(len(groups), func(i, j int) {
		groups[i], groups[j] = groups[j], groups[i]
	})

	totalGroups := len(groups)

	fmt.Println("Total Groups:", totalGroups)

	var randomGroups []string

	switch {
	case totalGroups > 10:
		randomGroups = groups[:10]
	default:
		randomGroups = groups
	}

	fmt.Println("Random Groups:", len(randomGroups))

	for _, href := range randomGroups {
		page := browser.MustPage(href).MustWaitLoad()

		joinButton := page.MustElement(`div[aria-label="Join Group"]`)

		joinButton.MustClick()

		fmt.Println("Group Joined")

		time.Sleep(5 * time.Second)

		page.MustScreenshot("facebook.png")
	}
}
