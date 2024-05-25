package cmd

import (
	"regexp"
	"strings"
)

func GetPasteIds(content string, classRegex string, hrefRegex string) []string {
	var results []string
	
	re := regexp.MustCompile(classRegex)
	classes := re.FindAllString(content, -1)
	for _, class := range classes {
		re := regexp.MustCompile(hrefRegex)
		pasteIds := re.FindAllString(class, -1)
		for _, pasteId := range pasteIds {
			result := strings.Replace(strings.Replace(pasteId, "\"", "", -1), "/", "", -1)
			results = append(results, result)
		}
	}
	return results
}