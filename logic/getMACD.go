package logic

import (
	"strconv"
)

func GetMACD(candles [][]string) map[string]float64 {
	var closes []float64
	var signal float64
	for _, c := range candles {
		floatClose, err := strconv.ParseFloat(c[4], 64)

		if err != nil {
			panic(err)
		}

		closes = append(closes, floatClose)
	}

	ema12 := GetEMA(candles, 12)
	ema26 := GetEMA(candles, 26)
	macdLine := ema12 - ema26
	var macdValues []float64

	for i := 0; i < len(closes); i++ {
		if i < 26 {
			continue
		}
		ema12I := GetEMAFromCloses(closes[:i], 12)
		ema26I := GetEMAFromCloses(closes[:i], 26)
		macdValues = append(macdValues, ema12I-ema26I)
	}
	if len(macdValues) < 9 {
		signal = 0
	} else {
		var sum float64
		for _, v := range macdValues[len(macdValues)-9:] {
			sum += v
		}
		signal = sum / 9
	}
	histogram := macdLine - signal
	result := make(map[string]float64)

	result["macdLine"] = macdLine
	result["signalLine"] = signal
	result["histogram"] = histogram

	return result
}
