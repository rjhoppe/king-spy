package wsb

func formatOutput(ticker string) (spacing string) {
	tickerLen := len(ticker)
	switch tickerLen {
	case 1:
		spacing = "     "
	case 2:
		spacing = "    "
	case 3:
		spacing = "   "
	case 4:
		spacing = "  "
	case 5:
		spacing = " "
	}
	return spacing
}
