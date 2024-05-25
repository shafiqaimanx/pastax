package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shafiqaimanx/pastax/colors"
	"github.com/shafiqaimanx/pastax/core"
	"github.com/shafiqaimanx/pastax/src/web"
)

func main() {
	var recentFlag 	bool
	var usernameFlag string

	flag.BoolVar(&recentFlag, "recent", false, ""); flag.BoolVar(&recentFlag, "r", false, "")
	flag.StringVar(&usernameFlag, "username", "", ""); flag.StringVar(&usernameFlag, "u", "", "")
	flag.Usage = func() {
		core.HelpMenu()
	}
    flag.Parse()

    if flag.NFlag() == 0 {
        flag.Usage()
        os.Exit(0)
    }

	if recentFlag && usernameFlag != "" {
		fmt.Printf("%s: Only one of -recent or -username can be used at a time.\n\n", colors.BERROR)
        flag.Usage()
        os.Exit(0)
	}

	if recentFlag {
		core.PastaxBanner(); fmt.Println("")
		web.GetRecentArchive(core.RECENT_DIR_PATH, 2)
	}

	if usernameFlag != "" {
		core.PastaxBanner(); fmt.Println("")
		pasteIds := web.GetUserContent(core.USERS_DIR_PATH, usernameFlag)
		for _, pasteId := range pasteIds {
			web.DownloadTheContent(core.USERS_DIR_PATH+"/"+usernameFlag, pasteId)
		}
	}
}