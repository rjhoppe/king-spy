package all

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type AllInputs struct {
	spyPerf      float64
	ticker1YPerf float64
	ticker       string
}

func FormatAll(i AllInputs) {
	if i.spyPerf > 0 {
		spyPositive = "+"
	} else {
		spyPositive = ""
	}

	deltaPerf := i.ticker1YPerf - i.spyPerf
	if deltaPerf > 0 {
		deltaPositive = "+"
	} else {
		deltaPositive = ""
	}

	if spyPositive == "+" {
		spyC := color.New(color.FgGreen)
		spyTextC := color.YellowString("SPY")
		fmt.Println("")
		fmt.Printf("%v 1Y performance: ", spyTextC)
		spyC.Printf("%v%.2f%% \n", spyPositive, i.spyPerf)
	} else {
		spyC := color.New(color.FgRed)
		spyTextC := color.YellowString("SPY")
		fmt.Println("")
		fmt.Printf("%v 1Y performance: ", spyTextC)
		spyC.Printf("%v%.2f%% \n", spyPositive, i.spyPerf)
	}

	if deltaPositive == "+" {
		deltaC := color.New(color.FgGreen)
		fmt.Printf("%v 1Y performance vs SPY: ", color.YellowString(strings.ToUpper(i.ticker)))
		deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
	} else {
		deltaC := color.New(color.FgRed)
		fmt.Printf("%v 1Y performance vs SPY: ", color.YellowString(strings.ToUpper(i.ticker)))
		deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
	}
}
