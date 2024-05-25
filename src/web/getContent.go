package web

import (
	"fmt"
	"path/filepath"

	"github.com/shafiqaimanx/pastax/core"
	"github.com/shafiqaimanx/pastax/src/cmd"
	agent "github.com/shafiqaimanx/subEX/src"
	engine "github.com/shafiqaimanx/subEX/src/core"
)

func DownloadTheContent(outputDirectoryPath string, pasteId string) {
	search := &engine.Engine{}
	search.Get(core.DOWNLOAD_URL+pasteId, agent.RandomUserAgent())
	res := search.Body()

	fileName := fmt.Sprintf("%s-pastebins-%s.txt", cmd.DateAndTime(), pasteId)
	cmd.WriteToFile(filepath.Join(outputDirectoryPath, fileName), []byte(res))
}
