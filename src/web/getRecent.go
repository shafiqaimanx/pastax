package web

import (
	"fmt"
	"sync"
	"time"

	"github.com/shafiqaimanx/pastax/colors"
	"github.com/shafiqaimanx/pastax/core"
	"github.com/shafiqaimanx/pastax/src/cmd"
	agent "github.com/shafiqaimanx/subEX/src"
	engine "github.com/shafiqaimanx/subEX/src/core"
)

func GetRecentArchive(directoryRecentPath string, sleepTime time.Duration) {
	directoryRecentExist, err := cmd.DirChecker(directoryRecentPath)
	if err != nil {
		fmt.Println(err)
	}

	if !directoryRecentExist {
		cmd.CreateDir(directoryRecentPath)
	}

	fmt.Printf("%s: Retrieve the most recent PasteBin archive.\n\n", colors.BINFO)
	var wg sync.WaitGroup

	for _, recentPasteId := range getRecentArchiveId(core.ARCHIVE_URL) {
		wg.Add(1)
		go func (recentPasteId string)  {
			defer wg.Done()
			DownloadTheContent(core.RECENT_DIR_PATH, recentPasteId)
		}(recentPasteId)
		time.Sleep(sleepTime * time.Second)
	}
	wg.Wait()
}

func getRecentArchiveId(url string) []string {
	search := &engine.Engine{}
	search.Get(url, agent.RandomUserAgent())
	responseBody := search.Body()
	results := cmd.GetPasteIds(responseBody, core.PUBLIC_RGX, core.HREF_RGX)
	return results
}
