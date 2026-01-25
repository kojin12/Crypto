package logic

import "strconv"

func GetVolume(candles [][]string, period int) string {
	lastCandles := candles[len(candles)-period:]
	var sum float64

	for _, i := range lastCandles {
		floatI, _ := strconv.ParseFloat(i[6], 64)
		sum += floatI
	}

	avg := sum / float64(period)
	lastCandle, _ := strconv.ParseFloat(candles[len(candles)-1][6], 64)
	if lastCandle > avg*3 {
		return "Strong Up"
	}
	if lastCandle > avg*1.5 {
		return "Volume up"
	} else {
		return "Volume down"
	}

}
