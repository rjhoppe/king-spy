package utils

import (
	"strconv"
	"time"
)

type TimeAssignVals struct {
	TimeVal string
	Ticker  string
	Cmd     string
	UrlType string
}

type UrlAssignVals struct {
	StartTime string
	EndTime   string
	Timeframe string
	Iterator  int
}

func AssignTime(t TimeAssignVals) (u UrlAssignVals) {
	var (
		startTime string
		endTime   string
		timeframe string
		iterator  int
	)

	curTime := time.Now()
	switch t.TimeVal {
	case "1M":
		pastTimeVal := curTime.AddDate(0, -1, 0)
		startTime = pastTimeVal.Format(time.RFC3339)
		if t.Cmd == "high" || t.Cmd == "low" {
			timeframe = "1D"
			iterator = 10
			endTime = curTime.Format(time.RFC3339)
		} else {
			timeframe = ""
			iterator = 0
			endTimeVal := pastTimeVal.Add(72 * time.Hour)
			endTime = endTimeVal.Format(time.RFC3339)
		}
	case "3M":
		pastTimeVal := curTime.AddDate(0, -3, 0)
		startTime = pastTimeVal.Format(time.RFC3339)
		endTime = curTime.Format(time.RFC3339)
		timeframe = "1W"
		iterator = 10
	case "6M":
		pastTimeVal := curTime.AddDate(0, -6, 0)
		startTime = pastTimeVal.Format(time.RFC3339)
		if t.Cmd == "high" || t.Cmd == "low" {
			timeframe = "1M"
			iterator = 4
			endTime = curTime.Format(time.RFC3339)
		} else {
			timeframe = ""
			iterator = 0
			endTimeVal := pastTimeVal.Add(72 * time.Hour)
			endTime = endTimeVal.Format(time.RFC3339)
		}
	case "1Y":
		pastTimeVal := curTime.AddDate(-1, 0, 0)
		startTime = pastTimeVal.Format(time.RFC3339)
		if (t.Cmd == "high") || (t.Cmd == "low") {
			endTime = curTime.Format(time.RFC3339)
			timeframe = "1M"
			iterator = 11
		} else {
			timeframe = ""
			iterator = 0
			endTimeVal := pastTimeVal.Add(72 * time.Hour)
			endTime = endTimeVal.Format(time.RFC3339)
		}
	case "3Y":
		pastTimeVal := curTime.AddDate(-3, 0, 0)
		startTime = pastTimeVal.Format(time.RFC3339)
		endTimeVal := pastTimeVal.Add(72 * time.Hour)
		endTime = endTimeVal.Format(time.RFC3339)
	case "YTD":
		curYear, _, _ := time.Now().Date()
		curYearString := strconv.Itoa(curYear)
		startTime = curYearString + "-01-01T00:00:00-00:00"
		endTime = curTime.Format(time.RFC3339)
	}

	u = UrlAssignVals{
		StartTime: startTime,
		EndTime:   endTime,
		Timeframe: timeframe,
		Iterator:  iterator,
	}

	return u
}
