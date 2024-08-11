package c2s

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Output struct {
	spyHist      float64
	spyLatest    float64
	tickerHist   float64
	tickerLatest float64
	tickerVal    string
	timeVal      string
}

func formatOutputC2S(o Output) {
	spyPerf := ((o.spyLatest - o.spyHist) / o.spyHist) * 100
	if spyPerf > 0 {
		spyPositive = "+"
	} else {
		spyPositive = ""
	}

	tickerPerf := ((o.tickerLatest - o.tickerHist) / o.tickerHist) * 100
	if tickerPerf > 0 {
		tickerPositive = "+"
	} else {
		tickerPositive = ""
	}

	deltaPerf := tickerPerf - spyPerf
	if deltaPerf > 0 {
		deltaPositive = "+"
	} else {
		deltaPositive = ""
	}

	spacing := FormatOutput(o.tickerVal)
	fmt.Println("")
	fmt.Printf("%v%v%v \n", color.YellowString(strings.ToUpper(o.tickerVal)), color.YellowString(" Performance VS SPY "), color.YellowString(o.timeVal))
	if spyPositive == "+" {
		spyC := color.New(color.FgGreen)
		spyTextC := color.YellowString("SPY")
		fmt.Println("==================================================================================")
		fmt.Printf("%v %v performance: %v", spyTextC, o.timeVal, spacing)
		spyC.Printf("%v%.2f%% \n", spyPositive, spyPerf)
	} else {
		spyC := color.New(color.FgRed)
		spyTextC := color.YellowString("SPY")
		fmt.Println("==================================================================================")
		fmt.Printf("%v %v performance: %v", spyTextC, o.timeVal, spacing)
		spyC.Printf("%v%.2f%% \n", spyPositive, spyPerf)
	}

	if tickerPositive == "+" {
		tickerC := color.New(color.FgGreen)
		fmt.Printf("%v %v performance:        ", color.YellowString(strings.ToUpper(o.tickerVal)), o.timeVal)
		tickerC.Printf("%v%.2f%% \n", tickerPositive, tickerPerf)
	} else {
		tickerC := color.New(color.FgRed)
		fmt.Printf("%v %v performance:        ", color.YellowString(strings.ToUpper(o.tickerVal)), o.timeVal)
		tickerC.Printf("%v%.2f%% \n", tickerPositive, tickerPerf)
	}

	if deltaPositive == "+" {
		deltaC := color.New(color.FgGreen)
		fmt.Printf("%v %v performance vs SPY: ", color.YellowString(strings.ToUpper(o.tickerVal)), o.timeVal)
		deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
	} else {
		deltaC := color.New(color.FgRed)
		fmt.Printf("%v %v performance vs SPY: ", color.YellowString(strings.ToUpper(o.tickerVal)), o.timeVal)
		deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
	}
	fmt.Println("==================================================================================")
}

func FormatOutput(ticker string) (spacing string) {
	spaceLen := len(ticker) - 3
	switch spaceLen {
	case 2:
		spacing = "         "
	case 1:
		spacing = "        "
	case 0:
		spacing = "       "
	case -1:
		spacing = "      "
	case -2:
		spacing = "     "
	}
	return spacing
}
