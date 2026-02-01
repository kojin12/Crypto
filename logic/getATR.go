package logic

import (
	"errors"
	"math"
	"strconv"
)

func GetATR(candles [][]string, period int) (float64, error) {
	if period <= 0 {
		return 0, errors.New("period must be greater than 0")
	}

	if len(candles) < period {
		return 0, errors.New("not enough candlesticks to calculate ATR")
	}

	lastCandles := candles[len(candles)-period:]
	var sumTR float64

	for i := range lastCandles {
		high, err := strconv.ParseFloat(lastCandles[i][2], 64)
		if err != nil {
			return 0, err
		}

		low, err := strconv.ParseFloat(lastCandles[i][3], 64)
		if err != nil {
			return 0, err
		}

		if i == 0 {
			tr := high - low
			sumTR += tr
			continue
		}

		prevClose, err := strconv.ParseFloat(lastCandles[i-1][4], 64)
		if err != nil {
			return 0, err
		}

		rangeHL := high - low
		absHigh := math.Abs(high - prevClose)
		absLow := math.Abs(low - prevClose)

		tr := math.Max(rangeHL, math.Max(absHigh, absLow))
		sumTR += tr
	}

	return sumTR / float64(period), nil
}
