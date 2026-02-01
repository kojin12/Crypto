package logic

import (
	"errors"
	"strconv"
)

func GetEMA(candles [][]string, count int) (float64, error) {
	if count <= 0 {
		return 0, errors.New("period must be > 0")
	}

	if len(candles) < count {
		return 0, errors.New("not enough candles to calculate EMA")
	}

	var closes []float64
	for _, c := range candles {
		if len(c) <= 4 {
			return 0, errors.New("invalid candle format")
		}

		price, err := strconv.ParseFloat(c[4], 64)
		if err != nil {
			return 0, err
		}
		closes = append(closes, price)
	}

	k := 2.0 / float64(count+1)

	sum := 0.0
	for _, price := range closes[:count] {
		sum += price
	}
	ema := sum / float64(count)

	for _, price := range closes[count:] {
		ema = price*k + ema*(1-k)
	}

	return ema, nil
}

func GetEMAFromCloses(closes []float64, count int) (float64, error) {
	if count <= 0 {
		return 0, errors.New("period must be > 0")
	}

	if len(closes) < count {
		return 0, errors.New("not enough values to calculate EMA")
	}

	k := 2.0 / float64(count+1)

	sum := 0.0
	for _, price := range closes[:count] {
		sum += price
	}
	ema := sum / float64(count)

	for _, price := range closes[count:] {
		ema = price*k + ema*(1-k)
	}

	return ema, nil
}
