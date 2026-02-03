package logic

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

func Final(candles [][]string, candles60 [][]string) FinalResult {
	var res FinalResult

	ema50 := GetEMA(candles, 50)
	res.Ema50 = ema50

	ema200 := GetEMA(candles, 200)

	res.Ema200 = ema200

	res.Trend15 = GetTrend15(res.Ema50, res.Ema200)

	reg := RegCandles(candles)

	adx := GetADX(reg, 14)
	res.ADX = adx

	atrPercent := GetATRPercent(candles, 14)

	res.ATRPercent = atrPercent

	ema50_60 := GetEMA(candles60, 50)

	ema200_60 := GetEMA(candles60, 200)

	res.Trend60 = GetTrend60(ema50_60, ema200_60)

	rsi := GetRSI(candles)
	res.Rsi = rsi
	if res.Rsi > 70 {
		res.RsiFilter = "Dont Buy"
	} else if res.Rsi < 30 {
		res.RsiFilter = "Dont Sell"
	}
	res.RsiFilter = "Neutral"
	macd := GetMACD(candles)

	res.MACD = macd

	atr := GetATR(candles, 14)
	res.Atr = atr

	volume := GetVolume(candles, 20)

	res.Volume = volume

	patterns := GetPatterns(candles)
	res.Patterns = patterns

	supports, resistances := GetRS(candles)
	res.Supports = supports
	res.Resistance = resistances

	return res
}
