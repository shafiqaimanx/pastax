package core

import (
	"fmt"

	"github.com/shafiqaimanx/pastax/colors"
)

func PastaxBanner() {
	fmt.Printf("%s╔═╗╔═╗╔═╗╔╦╗╔═╗═╗╔═%s\n", colors.SNOW, colors.RESET)
	fmt.Printf("%s╠═╝╠═╣╚═╗ ║ ╠═╣ ║╣ %s\n", colors.SNOW, colors.RESET)
	fmt.Printf("%s╩  ╩ ╩╚═╝ ╩ ╩ ╩═╝╚═%s\n", colors.SNOW, colors.RESET)
}