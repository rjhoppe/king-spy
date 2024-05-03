package utils

import "github.com/common-nighthawk/go-figure"

func AsciiTitleText() {
	title := figure.NewColorFigure("KING SPY", "", "yellow", true)
  title.Print()
	println()
}