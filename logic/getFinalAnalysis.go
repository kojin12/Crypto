package logic

import "fmt"

type FinalResult struct {
	Ema50      float64
	Ema200     float64
	Trend15    string
	Trend60    string
	Rsi        float64
	RsiFilter  string
	MACD       map[string]float64
	Atr        float64
	Volume     string
	Patterns   map[string]bool
	Supports   []float64
	Resistance []float64
	ADX        float64
	ATRPercent float64
}

func Final(candles [][]string, candles60 [][]string) (FinalResult, error) {
	var res FinalResult

	ema50, err := GetEMA(candles, 50)
	if err != nil {
		return res, fmt.Errorf("ema50: %w", err)
	}
	res.Ema50 = ema50

	ema200, err := GetEMA(candles, 200)
	if err != nil {
		return res, fmt.Errorf("ema200: %w", err)
	}
	res.Ema200 = ema200

	res.Trend15 = GetTrend15(res.Ema50, res.Ema200)

	reg := RegCandles(candles)

	adx, err := GetADX(reg, 14)
	if err != nil {
		return res, fmt.Errorf("adx: %w", err)
	}
	res.ADX = adx

	atrPercent, err := GetATRPercent(candles, 14)
	if err != nil {
		return res, fmt.Errorf("atr percent: %w", err)
	}
	res.ATRPercent = atrPercent

	ema50_60, err := GetEMA(candles60, 50)
	if err != nil {
		return res, fmt.Errorf("ema50 60: %w", err)
	}
	ema200_60, err := GetEMA(candles60, 200)
	if err != nil {
		return res, fmt.Errorf("ema200 60: %w", err)
	}
	res.Trend60 = GetTrend60(ema50_60, ema200_60)

	rsi, err := GetRSI(candles)
	if err != nil {
		return res, fmt.Errorf("rsi: %w", err)
	}
	res.Rsi = rsi
	if res.Rsi > 70 {
		res.RsiFilter = "Dont Buy"
	} else if res.Rsi < 30 {
		res.RsiFilter = "Dont Sell"
	}

	macd, err := GetMACD(candles)
	if err != nil {
		return res, fmt.Errorf("macd: %w", err)
	}
	res.MACD = macd

	atr, err := GetATR(candles, 14)
	if err != nil {
		return res, fmt.Errorf("atr: %w", err)
	}
	res.Atr = atr

	volume, err := GetVolume(candles, 20)
	if err != nil {
		return res, fmt.Errorf("volume: %w", err)
	}
	res.Volume = volume

	patterns := GetPatterns(candles)
	res.Patterns = patterns

	supports, resistances, err := GetRS(candles)
	if err != nil {
		return res, fmt.Errorf("rs: %w", err)
	}
	res.Supports = supports
	res.Resistance = resistances

	return res, nil
}
