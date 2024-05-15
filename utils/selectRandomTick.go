package utils

import (
	"math/rand"
)

func SelectRandomTicker() (string, string) {
	// make this a dict with format tickerName: tickerSymbol or vice versa
	tickerDict := map[string]string{
		"AAPL":  "Apple Inc.",
		"GOOGL": "Alphabet Inc.",
		"AMZN":  "Amazon.com Inc",
		"META":  "Meta Platforms Inc",
		"NVDA":  "NVIDIA Corp",
		"TSLA":  "Tesla, Inc",
		"MSFT":  "Microsoft Corp",
		"NFLX":  "Netflix Inc",
		"CRM":   "Salesforce Inc",
		"CAT":   "Caterpillar Inc.",
		"V":     "Visa Inc",
		"MA":    "Mastercard Inc",
		"AXP":   "American Express Company",
		"KO":    "Coca-Cola Co",
		"MCD":   "McDonald's Corp",
		"CMG":   "Chipotle Mexican Grill, Inc.",
		"DIS":   "Walt Disney Co",
		"VZ":    "Verizon Communications Inc.",
		"CVX":   "Chevron Corp",
		"XOM":   "Exxon Mobil Corp",
		"OXY":   "Occidental Petroleum Corp",
		"PWR":   "Quanta Services Inc",
		"CSCO":  "Cisco Systems Inc",
		"ORCL":  "Oracle Corp",
		"INTU":  "Intuit Inc",
		"ADBE":  "Adobe Inc",
		"DELL":  "Dell Technologies Inc",
		"AVGO":  "Broadcom Inc",
		"AMD":   "Advanced Micro Device, Inc.",
		"SMCI":  "Super Micro Computer Inc",
		"KLAC":  "KLA Corp",
		"TSM":   "Taiwan Semiconductor Mfg. Co. Ltd.",
		"MRVL":  "Marvell Technology Inc",
		"MU":    "Micron Technology Inc",
		"ARM":   "Arm Holdings PLC - ADR",
		"SNPS":  "Synopsys Inc",
		"AMAT":  "Applied Materials, Inc.",
		"QQQ":   "Invesco QQQ Trust Series 1",
		"TQQQ":  "ProShares UltraPro QQQ",
		"ASML":  "ASML Holding NV",
		"WM":    "Waste Management, Inc.",
		"GE":    "General Electrics Co",
		"GS":    "Goldman Sachs Group Inc",
		"BLK":   "BlackRock Inc",
		"JPM":   "JPMorgan Chase & Co",
		"MS":    "Morgan Stanley",
		"APO":   "Apollo Global Management Ord Shs",
		"ACN":   "Accenture Plc",
		"PGR":   "Progressive Corp",
		"ALL":   "Allstate Corp",
		"BKNG":  "Booking Holdings Inc",
		"EXPE":  "Expedia Group Inc",
		"ABNB":  "Airbnb Inc",
		"UAL":   "United Airlines Holdings Inc",
		"DAL":   "Delta Air Lines, Inc.",
		"LLY":   "Eli Lilly And Co",
		"UNH":   "UnitedHealth Group Inc",
		"VRTX":  "Vertex Pharmaceuticals Incorporated",
		"MMM":   "3M Co",
		"ABT":   "Abbot Laboratories",
		// "ALGN",
		// "ISRG",
		// "SHOP",
		// "TEAM",
		// "PANW",
	}

	i := 0
	tickerList := make([]string, len(tickerDict))
	for k := range tickerDict {
		tickerList[i] = k
		i++
	}

	randomIndex := rand.Intn(len(tickerList))
	randTickVal := tickerList[randomIndex]
	randTickName := tickerDict[randTickVal]
	return randTickVal, randTickName
}
