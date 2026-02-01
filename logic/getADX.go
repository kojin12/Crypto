package logic

import (
	"errors"
	"math"
)

func GetADX(candles []OHLC, period int) (float64, error) {
	if period <= 0 {
		return 0, errors.New("period must be greater than 0")
	}

	if len(candles) < period+1 {
		return 0, errors.New("not enough candlesticks to calculate ADX")
	}

	trs := make([]float64, len(candles))
	plusDMs := make([]float64, len(candles))
	minusDMs := make([]float64, len(candles))

	for i := 1; i < len(candles); i++ {
		curr := candles[i]
		prev := candles[i-1]

		tr := math.Max(
			curr.High-curr.Low,
			math.Max(
				math.Abs(curr.High-prev.Close),
				math.Abs(curr.Low-prev.Close),
			),
		)
		trs[i] = tr

		upMove := curr.High - prev.High
		downMove := prev.Low - curr.Low

		if upMove > downMove && upMove > 0 {
			plusDMs[i] = upMove
		}

		if downMove > upMove && downMove > 0 {
			minusDMs[i] = downMove
		}
	}

	var smTR, smPlus, smMinus float64
	for i := 1; i <= period; i++ {
		smTR += trs[i]
		smPlus += plusDMs[i]
		smMinus += minusDMs[i]
	}

	for i := period + 1; i < len(candles); i++ {
		smTR = smTR - (smTR / float64(period)) + trs[i]
		smPlus = smPlus - (smPlus / float64(period)) + plusDMs[i]
		smMinus = smMinus - (smMinus / float64(period)) + minusDMs[i]
	}

	if smTR == 0 {
		return 0, errors.New("smoothed true range is zero, cannot calculate DI")
	}

	diPlus := (smPlus / smTR) * 100
	diMinus := (smMinus / smTR) * 100

	denominator := diPlus + diMinus
	if denominator == 0 {
		return 0, errors.New("sum of DI+ and DI- is zero, cannot calculate ADX")
	}

	adx := (math.Abs(diPlus-diMinus) / denominator) * 100
	return adx, nil
}
