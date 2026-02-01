package logic

import "math"

func GetADX(candles []OHLC, period int) float64 {
	if len(candles) < period+1 {
		return 0
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
		return 0
	}

	diPlus := (smPlus / smTR) * 100
	diMinus := (smMinus / smTR) * 100

	if diPlus+diMinus == 0 {
		return 0
	}

	return (math.Abs(diPlus-diMinus) / (diPlus + diMinus)) * 100
}
