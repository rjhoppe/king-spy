package random

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type RandomOutput struct {
	spyHist      float64
	spyLatest    float64
	tickerHist   float64
	tickerLatest float64
	tickerVal    string
	tickerName   string
	timeVal      string
}

func formatOutputRandom(r RandomOutput) {
	spyPerf := ((r.spyLatest - r.spyHist) / r.spyHist) * 100
	if spyPerf > 0 {
		spyPositive = "+"
	} else {
		spyPositive = ""
	}

	tickerPerf := ((r.tickerLatest - r.tickerHist) / r.tickerHist) * 100
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
		fmt.Println("")
		fmt.Println("==================================================================================")
		spyValC := color.New(color.FgGreen)
		spyTextC := color.YellowString("SPY")
		fmt.Printf("%v: %v performance: ", spyTextC, r.timeVal)
		spyValC.Printf("%v%.2f%% \n", spyPositive, spyPerf)
	} else {
		fmt.Println("")
		fmt.Println("==================================================================================")
		spyValC := color.New(color.FgRed)
		spyTextC := color.YellowString("SPY")
		fmt.Printf("%v: %v performance: ", spyTextC, r.timeVal)
		spyValC.Printf("%v%.2f%% \n", spyPositive, spyPerf)
	}

	if tickerPositive == "+" {
		tickerValC := color.New(color.FgGreen)
		fmt.Printf("%v %v: %v performance: ", color.YellowString(r.tickerName), color.YellowString("("+strings.ToUpper(r.tickerVal)+")"), r.timeVal)
		tickerValC.Printf("%v%.2f%% \n", tickerPositive, tickerPerf)
	} else {
		tickerValC := color.New(color.FgRed)
		fmt.Printf("%v %v: %v performance: ", color.YellowString(r.tickerName), color.YellowString("("+strings.ToUpper(r.tickerVal)+")"), r.timeVal)
		tickerValC.Printf("%v%.2f%% \n", tickerPositive, tickerPerf)
	}

	if deltaPositive == "+" {
		deltaC := color.New(color.FgGreen)
		fmt.Printf("%v %v: %v performance vs SPY: ", color.YellowString(r.tickerName), color.YellowString("("+strings.ToUpper(r.tickerVal)+")"), r.timeVal)
		deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
		fmt.Println("==================================================================================")
	} else {
		deltaC := color.New(color.FgRed)
		fmt.Printf("%v %v: %v performance vs SPY: ", color.YellowString(r.tickerName), color.YellowString("("+strings.ToUpper(r.tickerVal)+")"), r.timeVal)
		deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
		fmt.Println("==================================================================================")
	}
}
