package utils

import (
	"errors"
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/haron1996/fb/cookies"
)

func CheckSellerLoginStatus(browser *rod.Browser, page *rod.Page) (*rod.Page, error) {
	pageHasLoginButton := page.MustHas(`button[name="login"]`)

	switch {
	case pageHasLoginButton:
		// Define the session cookies
		cookies := []*proto.NetworkCookieParam{
			{
				Name:     "c_user",
				Value:    cookies.C_user,
				Domain:   ".facebook.com",
				Path:     "/",
				HTTPOnly: true,
				Secure:   true,
				SameSite: "None",
				Priority: "Medium",
			},
			{
				Name:     "datr",
				Value:    cookies.Datr,
				Domain:   ".facebook.com",
				Path:     "/",
				HTTPOnly: true,
				Secure:   true,
				SameSite: "None",
				Priority: "Medium",
			},
			{
				Name:     "fr",
				Value:    cookies.Fr,
				Domain:   ".facebook.com",
				Path:     "/",
				HTTPOnly: true,
				Secure:   true,
				SameSite: "None",
				Priority: "Medium",
			},
			{
				Name:     "presence",
				Value:    cookies.Presence,
				Domain:   ".facebook.com",
				Path:     "/",
				HTTPOnly: true,
				Secure:   true,
				SameSite: "None",
				Priority: "Medium",
			},
			{
				Name:     "sb",
				Value:    cookies.Sb,
				Domain:   ".facebook.com",
				Path:     "/",
				HTTPOnly: true,
				Secure:   true,
				SameSite: "None",
				Priority: "Medium",
			},
			{
				Name:     "wd",
				Value:    cookies.Wd,
				Domain:   ".facebook.com",
				Path:     "/",
				HTTPOnly: true,
				Secure:   true,
				SameSite: "None",
				Priority: "Medium",
			},
			{
				Name:     "xs",
				Value:    cookies.Xs,
				Domain:   ".facebook.com",
				Path:     "/",
				HTTPOnly: true,
				Secure:   true,
				SameSite: "None",
				Priority: "Medium",
			},
		}

		// Inject the session cookie
		err := browser.SetCookies(cookies)
		if err != nil {
			return page, fmt.Errorf("error setting cookies to browser: %v", err)
		}

		// check if cookies are valid
		page.MustNavigate("https://web.facebook.com/").MustWaitLoad().MustWaitDOMStable()

		pageHasLoginButton := page.MustHas(`button[name="login"]`)

		switch {
		case pageHasLoginButton:
			return page, errors.New("invalid or expired cookies 😞")
		default:
			fmt.Println("Log in complete 😊")
			page.MustScreenshot("home.png")
		}
	default:
		fmt.Println("User is logged in 😊")
		page.MustScreenshot("home.png")
	}

	return page, nil
}
