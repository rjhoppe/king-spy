package high

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type HighOutput struct {
	ticker      string
	timeVal     string
	priceDiff   float64
	highestVal  float64
	highestDate string
	percDiff    float64
	cmdArgs     string
}

func formatOutputHigh(h HighOutput) {
	if h.cmdArgs == "high" {
		fmt.Println("")
	}

	fmt.Println("==================================================================================")
	fmt.Printf("The highest price of %v in the last %v time period was: %v on %v \n", color.YellowString(strings.ToUpper(h.ticker)), timeVal, color.GreenString("$"+strconv.FormatFloat(h.highestVal, 'f', 2, 64)), h.highestDate[:10])
	fmt.Printf("Price decrease off %v high: %v which is a %v decrease. \n", timeVal, color.RedString("-$"+strconv.FormatFloat(h.priceDiff, 'f', 2, 64)), color.RedString(strconv.FormatFloat(h.percDiff, 'f', 2, 64)+"%"))
	fmt.Println("==================================================================================")
}
