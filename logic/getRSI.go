package logic

import (
	"math"
	"strconv"
)

func GetRSI(candles [][]string) float64 {
	period := 14
	var closes []float64

	for _, c := range candles {
		floatClose, err := strconv.ParseFloat(c[4], 64)

		if err != nil {
			panic(err)
		}

		closes = append(closes, floatClose)
	}

	var gains []float64
	var losses []float64

	for i := 1; i < len(closes); i++ {
		diff := closes[i] - closes[i-1]
		if diff > 0 {
			gains = append(gains, diff)
			losses = append(losses, 0.0)
		} else {
			gains = append(gains, 0.0)
			losses = append(losses, math.Abs(diff))
		}
	}

	var (
		sumGains  float64
		sumLosses float64
	)

	for _, g := range gains[len(gains)-period:] {
		sumGains += g
	}
	for _, l := range losses[len(losses)-period:] {
		sumLosses += l
	}

	avgGain := sumGains / float64(period)
	avgLosses := sumLosses / float64(period)
	if avgLosses == 0 {
		return 100.0
	}
	rs := avgGain / avgLosses
	return 100 - (100 / (1 + rs))

}
