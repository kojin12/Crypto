package logic

import (
	"math"
	"strconv"
)

func GetATR(candles [][]string, period int) float64 {
	var count float64
	lastCandles := candles[len(candles)-period:]

	for j, _ := range lastCandles {
		if j == 0 {
			floatCandleHigh, _ := strconv.ParseFloat(lastCandles[j][2], 64)
			floatCandleLow, _ := strconv.ParseFloat(lastCandles[j][3], 64)

			tr := floatCandleHigh - floatCandleLow
			count += tr
		} else {
			floatCandleHigh, _ := strconv.ParseFloat(lastCandles[j][2], 64)
			floatCandleLow, _ := strconv.ParseFloat(lastCandles[j][3], 64)
			floatCandleClose, _ := strconv.ParseFloat(lastCandles[j-1][4], 64)
			data := floatCandleHigh - floatCandleLow
			abs1 := math.Abs(floatCandleHigh - floatCandleClose)
			abs2 := math.Abs(floatCandleLow - floatCandleClose)

			base := math.Max(data, abs1)
			tr := math.Max(base, abs2)
			count += tr
		}
	}

	return count / float64(period)
}
