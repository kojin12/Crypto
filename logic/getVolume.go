package logic

import (
	"errors"
	"strconv"
)

func GetVolume(candles [][]string, period int) (string, error) {
	if period <= 0 {
		return "", errors.New("period must be greater than 0")
	}

	if len(candles) < period {
		return "", errors.New("not enough candles to calculate volume trend")
	}

	lastCandles := candles[len(candles)-period:]
	var sum float64

	for i, c := range lastCandles {
		if len(c) <= 5 {
			return "", errors.New("invalid candle format at index " + strconv.Itoa(i))
		}
		vol, err := strconv.ParseFloat(c[5], 64)
		if err != nil {
			return "", err
		}
		sum += vol
	}

	avg := sum / float64(period)

	lastCandleVolStr := candles[len(candles)-1][5]
	lastCandleVol, err := strconv.ParseFloat(lastCandleVolStr, 64)
	if err != nil {
		return "", err
	}

	if lastCandleVol > avg*3 {
		return "Strong Up", nil
	}
	if lastCandleVol > avg*1.5 {
		return "Volume Up", nil
	}
	return "Volume Down", nil
}
