package logic

import (
	"errors"
	"strconv"
)

func GetSMA(candles [][]string, count int) (float64, error) {
	if count <= 0 {
		return 0, errors.New("period must be greater than 0")
	}

	if len(candles) < count {
		return 0, errors.New("not enough candles to calculate SMA")
	}

	var total float64
	last := candles[len(candles)-count:]

	for i, c := range last {
		if len(c) < 5 {
			return 0, errors.New("invalid candle format at index " + strconv.Itoa(i))
		}

		closeStr := c[4]
		closeVal, err := strconv.ParseFloat(closeStr, 64)
		if err != nil {
			return 0, err
		}

		total += closeVal
	}

	return total / float64(count), nil
}
