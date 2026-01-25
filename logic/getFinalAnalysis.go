package logic

import "sync"

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
}

func Final(candles [][]string, candles60 [][]string) FinalResult {
	var Res FinalResult
	var wg sync.WaitGroup

	
	Ema50 := GetEMA(candles, 50)
	Ema200 := GetEMA(candles, 200)
	Res.Ema50 = Ema50
	Res.Ema200 = Ema200

	wg.Add(9)

	go func() {
		defer wg.Done()
		Res.Trend15 = getTrend15(Ema50, Ema200)
	}()
	go func() {
		defer wg.Done()
		ema50_60 := GetEMA(candles60, 50)
		ema200_60 := GetEMA(candles60, 200)
		Res.Trend60 = getTrend60(ema50_60, ema200_60)
	}()

	go func() {
		defer wg.Done()
		Res.Rsi = GetRSI(candles)
	}()

	go func() {
		defer wg.Done()
		Res.MACD = getMACD(candles)
	}()
	go func() {
		defer wg.Done()
		Res.Atr = GetATR(candles, 14)
	}()
	go func() {
		defer wg.Done()
		Res.Volume = GetVolume(candles, 20)
	}()
	go func() {
		defer wg.Done()
		Res.Patterns = getPatterns(candles)
	}()
	go func() {
		defer wg.Done()
		support, resistance := getRS(candles)
		Res.Supports = support
		Res.Resistance = resistance
	}()

	go func() {
		defer wg.Done()
		rsi := GetRSI(candles)
		if rsi > 70 {
			Res.RsiFilter = "Dont Buy"
		} else if rsi < 30 {
			Res.RsiFilter = "Dont Sell"
		}
	}()

	wg.Wait()

	return Res
}
