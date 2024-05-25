package cmd

import (
	"fmt"
	"os"

	"github.com/shafiqaimanx/pastax/colors"
)

func DirChecker(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, fmt.Errorf("%s: %s%s%s does not exist", colors.BWARN, colors.YELLOW, path, colors.RESET)
	}
	return true, fmt.Errorf("%s: Directory %s%s%s exists", colors.BINFO, colors.YELLOW, path, colors.RESET)
}

func CreateDir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("%s: Failed creating directories.\n", colors.BERROR)
	}
	fmt.Printf("%s: %s%s%s created successfully.\n", colors.BOK, colors.YELLOW, path, colors.RESET)
}

func WriteToFile(path string, content []byte) {
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		fmt.Printf("%s: Failed to save into a file.\n", colors.BERROR)
		return
	}
	fmt.Printf("%s: Save into a file %s%s%s\n", colors.OK, colors.YELLOW, path, colors.RESET)
}

