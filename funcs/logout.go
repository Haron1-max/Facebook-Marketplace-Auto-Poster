package funcs

// import (
// 	"github.com/go-rod/rod"
// 	"github.com/go-rod/rod/lib/launcher"
// )

// func LogOut() {
// 	// Launch the browser with specific configurations
// 	u := launcher.New().UserDataDir("~/.config/google-chrome").Leakless(true).NoSandbox(true).Headless(true).MustLaunch()
// 	browser := rod.New().ControlURL(u).MustConnect()
// 	defer browser.MustClose() // Ensure the browser closes when main function exits

// 	page := browser.MustPage("https://web.facebook.com").MustWaitLoad()
// }
