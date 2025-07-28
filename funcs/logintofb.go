package funcs

import (
	"fmt"
	"log"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	"github.com/go-rod/rod/lib/proto"
	"github.com/haron1996/fb/cookies"
)

func LoginToFacebook() (*rod.Browser, *rod.Page) {
	dir := "~/.config/google-chrome"

	u := launcher.New().UserDataDir(dir).Leakless(true).NoSandbox(true).Headless(true).MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	//defer browser.MustClose()

	page := browser.MustPage("https://web.facebook.com/").MustWaitLoad().MustWaitDOMStable().MustSetViewport(1920, 1080, 1, false).MustWindowMaximize()

	//page = page.MustWindowMaximize()

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
			fmt.Println("Failed to set session cookie:", err)
			return nil, nil
		}

		// check if cookies are valid
		page = page.MustNavigate("https://web.facebook.com/").MustWaitLoad().MustWaitDOMStable()

		pageHasLoginForm, _, err := page.Has(`form[data-testid="royal_login_form"]`)
		if err != nil {
			log.Println("Error checking if page has login form:", err)
			return nil, nil
		}
		switch {
		case pageHasLoginForm:
			fmt.Println("Invalid or expired cookies 😞")
			os.Exit(1)
		default:
			fmt.Println("Log in complete 😊")
			page.MustScreenshot("home.png")
			return browser, page
		}
	default:
		fmt.Println("User is logged in 😊")
		page.MustScreenshot("home.png")
	}

	return browser, page
}
