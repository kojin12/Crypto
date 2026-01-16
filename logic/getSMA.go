package logic

import "strconv"

func GetSMA(candles [][]string, count int) float64 {
	total := 0
	last := candles[len(candles)-count:]

	for _, c := range last {
		close, err := strconv.ParseFloat(c[4], 64)
		if err != nil {
			panic(err)
		}

		total += int(close)
	}

	return float64(total) / float64(count)
}
