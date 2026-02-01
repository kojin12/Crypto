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
	ADX        float64
	ATRPercent float64
}

func Final(candles [][]string, candles60 [][]string) FinalResult {
	var Res FinalResult
	var wg sync.WaitGroup
	var mu sync.Mutex

	Ema50 := GetEMA(candles, 50)
	Ema200 := GetEMA(candles, 200)
	Res.Ema50 = Ema50
	Res.Ema200 = Ema200

	var Reg []OHLC

	wg.Add(11)

	go func() {
		defer wg.Done()
		Trend15 := GetTrend15(Ema50, Ema200)
		mu.Lock()
		Res.Trend15 = Trend15
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		reg := RegCandles(candles)
		mu.Lock()
		Reg = reg
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		ADX := GetADX(Reg, 14)
		Res.ADX = ADX
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		atrPerc := GetATRPercent(candles, 14)
		mu.Lock()
		Res.ATRPercent = atrPerc
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		ema50_60 := GetEMA(candles60, 50)
		ema200_60 := GetEMA(candles60, 200)
		Trend60 := GetTrend60(ema50_60, ema200_60)
		mu.Lock()
		Res.Trend60 = Trend60
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		rsi := GetRSI(candles)
		mu.Lock()
		Res.Rsi = rsi
		if rsi > 70 {
			Res.RsiFilter = "Dont Buy"
		} else if rsi < 30 {
			Res.RsiFilter = "Dont Sell"
		}
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		MACD := GetMACD(candles)
		mu.Lock()
		Res.MACD = MACD
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		Atr := GetATR(candles, 14)
		mu.Lock()
		Res.Atr = Atr
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		Volume := GetVolume(candles, 20)
		mu.Lock()
		Res.Volume = Volume
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		Patterns := GetPatterns(candles)
		mu.Lock()
		Res.Patterns = Patterns
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		support, resistance := GetRS(candles)
		mu.Lock()
		Res.Supports = support
		Res.Resistance = resistance
		mu.Unlock()
	}()

	wg.Wait()

	return Res
}
