package logic

import (
	"errors"
	"strconv"
)

func GetMACD(candles [][]string) (map[string]float64, error) {
	if len(candles) == 0 {
		return nil, errors.New("no candles provided")
	}

	var closes []float64
	for _, c := range candles {
		if len(c) < 5 {
			return nil, errors.New("invalid candle format")
		}
		price, err := strconv.ParseFloat(c[4], 64)
		if err != nil {
			return nil, err
		}
		closes = append(closes, price)
	}

	ema12, err := GetEMA(candles, 12)
	if err != nil {
		return nil, err
	}

	ema26, err := GetEMA(candles, 26)
	if err != nil {
		return nil, err
	}

	macdLine := ema12 - ema26

	macdValues := make([]float64, 0, len(closes)-25)
	for i := 26; i <= len(closes); i++ {
		if i < 26 {
			continue
		}
		ema12I, err := GetEMAFromCloses(closes[:i], 12)
		if err != nil {
			return nil, err
		}
		ema26I, err := GetEMAFromCloses(closes[:i], 26)
		if err != nil {
			return nil, err
		}
		macdValues = append(macdValues, ema12I-ema26I)
	}

	if len(macdValues) < 9 {
		return nil, errors.New("not enough MACD values to calculate signal line")
	}

	signal, err := GetEMAFromCloses(macdValues[len(macdValues)-9:], 9)
	if err != nil {
		return nil, err
	}

	histogram := macdLine - signal

	result := map[string]float64{
		"macdLine":   macdLine,
		"signalLine": signal,
		"histogram":  histogram,
	}

	return result, nil
}
