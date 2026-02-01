package logic

import (
	"errors"
	"strconv"
)

func GetATRPercent(candles [][]string, period int) (float64, error) {
	atr, err := GetATR(candles, period)
	if err != nil {
		return 0, err
	}

	if len(candles) == 0 {
		return 0, errors.New("no candlesticks provided")
	}

	lastCloseStr := candles[len(candles)-1][4]
	lastClose, err := strconv.ParseFloat(lastCloseStr, 64)
	if err != nil {
		return 0, err
	}

	if lastClose == 0 {
		return 0, errors.New("last close price is zero, cannot calculate percent")
	}

	return (atr / lastClose) * 100, nil
}
