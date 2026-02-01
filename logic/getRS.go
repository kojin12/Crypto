package logic

import (
	"errors"
	"strconv"
)

func GetRS(candles [][]string) ([]float64, []float64, error) {
	if len(candles) < 3 {
		return nil, nil, errors.New("not enough candles to calculate support/resistance")
	}

	var supports []float64
	var resistances []float64

	for i := 1; i < len(candles)-1; i++ {
		if len(candles[i]) < 5 || len(candles[i-1]) < 5 || len(candles[i+1]) < 5 {
			return nil, nil, errors.New("invalid candle format")
		}

		floatHigh, err := strconv.ParseFloat(candles[i][2], 64)
		if err != nil {
			return nil, nil, err
		}

		floatLow, err := strconv.ParseFloat(candles[i][3], 64)
		if err != nil {
			return nil, nil, err
		}

		floatPrevHigh, err := strconv.ParseFloat(candles[i-1][2], 64)
		if err != nil {
			return nil, nil, err
		}

		floatNextHigh, err := strconv.ParseFloat(candles[i+1][2], 64)
		if err != nil {
			return nil, nil, err
		}

		floatPrevLow, err := strconv.ParseFloat(candles[i-1][3], 64)
		if err != nil {
			return nil, nil, err
		}

		floatNextLow, err := strconv.ParseFloat(candles[i+1][3], 64)
		if err != nil {
			return nil, nil, err
		}

		if floatHigh > floatPrevHigh && floatHigh > floatNextHigh {
			resistances = append(resistances, floatHigh)
		}

		if floatLow < floatPrevLow && floatLow < floatNextLow {
			supports = append(supports, floatLow)
		}
	}

	return supports, resistances, nil
}
