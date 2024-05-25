package core

import (
	"fmt"

	"github.com/shafiqaimanx/pastax/colors"
)

func HelpMenu() {
	PastaxBanner()
	fmt.Printf("%spastax - pastebin scraper.%s\n\n", colors.DARKGRAY, colors.RESET)
	fmt.Printf("%s%sUsage%s: %spastax%s [flags]\n\n", colors.BOLD, colors.UNDERSCORE, colors.RESET, colors.BOLD, colors.RESET)
	fmt.Printf("%s%sFlags%s:\n", colors.BOLD, colors.UNDERSCORE, colors.RESET)
	fmt.Println("  -u, -username  retrieve user's PasteBin.")
	fmt.Println("  -r, -recent    retrieve the most recent PasteBin.")
	fmt.Println("  -h, -help      show help menu and exit.")
	fmt.Println("")
}