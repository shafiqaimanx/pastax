package web

import (
	"fmt"
	"os"

	"github.com/shafiqaimanx/pastax/colors"
	"github.com/shafiqaimanx/pastax/core"
	"github.com/shafiqaimanx/pastax/src/cmd"
	agent "github.com/shafiqaimanx/subEX/src"
	engine "github.com/shafiqaimanx/subEX/src/core"
)

func GetUserContent(directoryUserPath string, usernameFlag string) []string {
	directoryUserExist, err := cmd.DirChecker(directoryUserPath)
	if err != nil {
		fmt.Println(err)
	}

	if !directoryUserExist {
		cmd.CreateDir(directoryUserPath)
	}

	userExist, contentIds := checkUserExist(core.USERNAME_URL, usernameFlag)
	if !userExist {
		fmt.Printf("%s: User %s didn't exist.\n", colors.BERROR, usernameFlag)
		os.Exit(0)
	}
	fmt.Printf("%s: User %s exist.\n", colors.BINFO, usernameFlag)

	userDirExist, err := cmd.DirChecker(core.USERS_DIR_PATH+"/"+usernameFlag)
	if err != nil {
		fmt.Println(err)
	}

	if !userDirExist {
		cmd.CreateDir(core.USERS_DIR_PATH+"/"+usernameFlag)
	}

	fmt.Printf("%s: Retrieve %s's PasteBin.\n\n", colors.BINFO, usernameFlag)		
	return contentIds
}

func checkUserExist(url string, username string) (bool, []string) {
	search := &engine.Engine{}
	search.Get(url+username, agent.RandomUserAgent())
	responseCode := search.Response.StatusCode

	if responseCode != 404 {
		results := cmd.GetPasteIds(search.Body(), core.USER_PASTE_RGX, core.HREF_RGX)
		return true, results
	}
	return false, nil
}
