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

	if spyPositive == "+" {
		spyC := color.New(color.FgGreen)
		spyTextC := color.YellowString("SPY")
		// tickerTextC := color.YellowString("VS " + strings.ToUpper(o.tickerVal))
		fmt.Println("")
		// fmt.Printf("%v %v \n", spyTextC, tickerTextC)
		fmt.Println("==================================================================================")
		fmt.Printf("%v %v performance: ", spyTextC, o.timeVal)
		spyC.Printf("%v%.2f%% \n", spyPositive, spyPerf)
	} else {
		spyC := color.New(color.FgRed)
		spyTextC := color.YellowString("SPY")
		// tickerTextC := color.YellowString("VS " + strings.ToUpper(o.tickerVal))
		fmt.Println("")
		// fmt.Printf("%v %v \n", spyTextC, tickerTextC)
		fmt.Println("==================================================================================")
		fmt.Printf("%v %v performance: ", spyTextC, o.timeVal)
		spyC.Printf("%v%.2f%% \n", spyPositive, spyPerf)
	}

	if tickerPositive == "+" {
		tickerC := color.New(color.FgGreen)
		fmt.Printf("%v %v performance: ", color.YellowString(strings.ToUpper(o.tickerVal)), o.timeVal)
		tickerC.Printf("%v%.2f%% \n", tickerPositive, tickerPerf)
	} else {
		tickerC := color.New(color.FgRed)
		fmt.Printf("%v %v performance: ", color.YellowString(strings.ToUpper(o.tickerVal)), o.timeVal)
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
