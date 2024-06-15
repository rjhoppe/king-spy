package c2t

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type TickersOutput struct {
	ticker1Hist   float64
	ticker1Latest float64
	ticker2Hist   float64
	ticker2Latest float64
	ticker1Val    string
	ticker2Val    string
	timeVal       string
}

func formatOutputC2T(to TickersOutput) {
	ticker1Perf := ((to.ticker1Latest - to.ticker1Hist) / to.ticker1Hist) * 100
	if ticker1Perf > 0 {
		ticker1Positive = "+"
	} else {
		ticker1Positive = ""
	}

	ticker2Perf := ((to.ticker2Latest - to.ticker2Hist) / to.ticker2Hist) * 100
	if ticker2Perf > 0 {
		ticker2Positive = "+"
	} else {
		ticker2Positive = ""
	}

	deltaPerf := ticker1Perf - ticker2Perf
	if deltaPerf > 0 {
		deltaPositive = "+"
	} else {
		deltaPositive = ""
	}

	ticker1Spaces := tickerSpacing(to.ticker1Val)
	ticker2Spaces := tickerSpacing(to.ticker2Val)

	fmt.Println("")
	fmt.Println(color.YellowString(strings.ToUpper(to.ticker1Val) + " VS " + strings.ToUpper(to.ticker2Val) + " " + to.timeVal + " PERFORMANCE"))
	if ticker1Positive == "+" {
		ticker1C := color.New(color.FgGreen)
		fmt.Println("==================================================================================")
		fmt.Printf("%v: %v", color.YellowString(strings.ToUpper(to.ticker1Val)), ticker1Spaces)
		ticker1C.Printf("%v%.2f%% \n", ticker1Positive, ticker1Perf)
	} else {
		ticker1C := color.New(color.FgRed)
		fmt.Println("==================================================================================")
		fmt.Printf("%v: %v", color.YellowString(strings.ToUpper(to.ticker1Val)), ticker1Spaces)
		ticker1C.Printf("%v%.2f%% \n", ticker1Positive, ticker1Perf)
	}

	if ticker2Positive == "+" {
		tickerC := color.New(color.FgGreen)
		fmt.Printf("%v: %v", color.YellowString(strings.ToUpper(to.ticker2Val)), ticker2Spaces)
		tickerC.Printf("%v%.2f%% \n", ticker2Positive, ticker2Perf)
	} else {
		tickerC := color.New(color.FgRed)
		fmt.Printf("%v: %v", color.YellowString(strings.ToUpper(to.ticker2Val)), ticker2Spaces)
		tickerC.Printf("%v%.2f%% \n", ticker2Positive, ticker2Perf)
	}

	deltaSpacing := ""
	spacingCalc := (len(to.ticker1Val) + len(to.ticker2Val)) - 7
	for spacingCalc < 0 {
		deltaSpacing += " "
		spacingCalc += 1
	}

	if deltaPositive == "+" {
		deltaC := color.New(color.FgGreen)
		fmt.Printf("%v %v: %v", color.YellowString(strings.ToUpper(to.ticker1Val)+" VS"), color.YellowString(strings.ToUpper(to.ticker2Val)), deltaSpacing)
		deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
	} else {
		deltaC := color.New(color.FgRed)
		fmt.Printf("%v %v: %v", color.YellowString(strings.ToUpper(to.ticker1Val)+" VS"), color.YellowString(strings.ToUpper(to.ticker2Val)), deltaSpacing)
		deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
	}
	fmt.Println("==================================================================================")
}

func tickerSpacing(tickerVal string) string {
	lenOfTicker1 := len(tickerVal)
	spacing := 11
	spacing -= lenOfTicker1

	totalSpaces := " "
	for i := 1; i < spacing; i++ {
		totalSpaces += " "
	}

	return totalSpaces
}
