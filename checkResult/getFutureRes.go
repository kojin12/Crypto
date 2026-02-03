package checkresult

import "main/logic"

type FutureResult struct {
	Data        logic.FinalResult
	PriceResult float64
	Result      bool
	Direction   string
	Indicators  DataResult
}

func getFutureResult(prevData logic.FinalResult, startPrice float64, score int, data FutureResult, futurePrice float64) FutureResult {
	var result FutureResult
	result.Data = prevData
	if score > 3 {
		result.Direction = "Long"
	} else if score < -3 {
		result.Direction = "Short"
	}

	if result.Direction == "Short" && futurePrice < startPrice {

	}
}
