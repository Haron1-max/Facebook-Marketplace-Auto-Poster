package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"

	"github.com/go-rod/rod"
)

func FilterGroupsByMemberCount(groups rod.Elements, r *rand.Rand) ([]*rod.Element, int) {
	// Regular expression to capture the member count (e.g., "21K members" or "10K+ members")
	memberCountRegex := regexp.MustCompile(`(\d+(\.\d+)?)([KM]?)\s*members`)

	// groups to be selected
	var filteredGroups []*rod.Element

	for _, group := range groups {
		members := group.MustElements(`div.x9f619.x1n2onr6.x1ja2u2z.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xeuugli.xg83lxy.x1h0ha7o.x1120s5i.x1nn3v0j`)[0].MustText()

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
			fmt.Println("Error parsing member count:", err)
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
			filteredGroups = append(filteredGroups, group)
		}
	}

	totalFilteredGroups := len(filteredGroups)

	r.Shuffle(totalFilteredGroups, func(i, j int) { filteredGroups[i], filteredGroups[j] = filteredGroups[j], filteredGroups[i] })

	return filteredGroups, totalFilteredGroups
}
