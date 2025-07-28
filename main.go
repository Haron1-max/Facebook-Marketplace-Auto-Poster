package main

import (
	"fmt"

	"github.com/haron1996/fb/funcs"
	"github.com/haron1996/fb/funcs/marketplace"
)

func main() {
	browser, page := funcs.LoginToFacebook()
	if browser == nil || page == nil {
		fmt.Println("Browser or page is nil")
		return
	}

	// MARKETPLACE
	marketplace.ListItemForSale(browser, page)
	//marketplace.ListVehicleForSale(browser, page)

	//funcs.ListInMorePlaces(browser, page)
	//funcs.PostToGroups(browser, page)
}
