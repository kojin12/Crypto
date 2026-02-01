package logic

import "strconv"

func GetATRPercent(candles [][]string, period int) float64 {
	atr := GetATR(candles, period)
	lastClose, _ := strconv.ParseFloat(candles[len(candles)-1][4], 64)
	if lastClose == 0 {
		return 0
	}
	return (atr / lastClose) * 100
}
