package low

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type LowOutput struct {
	ticker     string
	timeVal    string
	priceDiff  float64
	lowestVal  float64
	lowestDate string
	percDiff   float64
	cmdArgs    string
}

func formatOutputLow(l LowOutput) {
	if l.cmdArgs == "low" {
		fmt.Println("")
	}

	fmt.Println("==================================================================================")
	fmt.Printf("The lowest price of %v in the last %v time period was: %v on %v \n", color.YellowString(strings.ToUpper(l.ticker)), timeVal, color.RedString("$"+strconv.FormatFloat(l.lowestVal, 'f', 2, 64)), l.lowestDate[:10])
	fmt.Printf("Price increase off %v low: %v which is a %v increase. \n", timeVal, color.GreenString("+$"+strconv.FormatFloat(l.priceDiff, 'f', 2, 64)), color.GreenString(strconv.FormatFloat(l.percDiff, 'f', 2, 64)+"%"))
	fmt.Println("==================================================================================")
}
