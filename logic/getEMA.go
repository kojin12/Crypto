package logic

import (
	"strconv"
)

func GetEMA(candles [][]string, count int) float64 {
	var closes []float64

	for _, c := range candles {
		floatClose, err := strconv.ParseFloat(c[4], 64)

		if err != nil {
			panic(err)
		}

		closes = append(closes, floatClose)
	}

	k := 2.0 / float64(count+1)

	var sum float64

	prevCloses := closes[:count]
	for _, i := range prevCloses {
		sum += i
	}

	emaPrev := sum / float64(count)

	for _, price := range closes[count:] {
		emaPrev = price*float64(k) + emaPrev*float64(1-k)
	}

	return emaPrev
}

func GetEMAFromCloses(closes []float64, count int) float64 {
	k := 2.0 / float64(count+1)

	sum := 0.0
	for _, price := range closes[:count] {
		sum += price
	}
	emaPrev := sum / float64(count)

	for _, price := range closes[count:] {
		emaPrev = price*k + emaPrev*(1-k)
	}

	return emaPrev
}
