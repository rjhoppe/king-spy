package utils

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

// ASCII art for the bare "ks" cmd
func AsciiTitleText() {
	title := figure.NewColorFigure("KING SPY", "", "yellow", true)
	title.Print()
	fmt.Println("")
}
