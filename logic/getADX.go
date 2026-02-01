package logic

import "math"

func GetADX(candles []OHLC, period int) float64 {
	trs := make([]float64, len(candles))
	plusDMs := make([]float64, len(candles))
	minusDMs := make([]float64, len(candles))

	for i := 1; i < len(candles); i++ {
		curr := candles[i]
		prev := candles[i-1]

		tr := math.Max(curr.High-curr.Low, math.Max(math.Abs(curr.High-prev.Close), math.Abs(curr.Low-prev.Close)))
		trs[i] = tr

		upMove := curr.High - prev.High
		downMove := prev.Low - curr.Low

		if upMove > downMove && upMove > 0 {
			plusDMs[i] = upMove
		} else {
			plusDMs[i] = 0
		}

		if downMove > upMove && downMove > 0 {
			minusDMs[i] = downMove
		} else {
			minusDMs[i] = 0
		}
	}

	smTR := 0.0
	smPlus := 0.0
	smMinus := 0.0
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

	diPlus := (smPlus / smTR) * 100
	diMinus := (smMinus / smTR) * 100
	dx := (math.Abs(diPlus-diMinus) / (diPlus + diMinus)) * 100
	adx := dx

	return adx
}
