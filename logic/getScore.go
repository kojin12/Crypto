package logic

func GetScore(data FinalResult, price float64) int {
	weights := map[string]float64{
		"Trend15":        1,
		"Trend60":        2,
		"TrendMatch":     1.5,
		"RSI":            1,
		"MACD":           2,
		"Volume":         1,
		"EMA":            1,
		"PatternBullish": 1,
		"PatternBearish": 1,
	}

	score := 0.0

	if data.Trend15 == "Up" {
		score += weights["Trend15"]
	} else {
		score -= weights["Trend15"]
	}

	if data.Trend60 == "Up" {
		score += weights["Trend60"]
	} else {
		score -= weights["Trend60"]
	}

	if data.Trend15 == data.Trend60 {
		score += weights["TrendMatch"]
	} else {
		score -= weights["TrendMatch"]
	}

	rsi := data.Rsi
	if rsi >= 45 && rsi <= 55 {
		score += weights["RSI"]
	} else if rsi > 55 && rsi <= 65 {
		score += weights["RSI"]
	} else if rsi >= 35 && rsi < 45 {
		score -= weights["RSI"]
	} else if rsi > 70 {
		score -= weights["RSI"] * 2
	} else if rsi < 30 {
		score -= weights["RSI"] * 2
	}

	macd := data.MACD
	if macd["macd_line"] > macd["signal_line"] {
		score += weights["MACD"]
		if macd["histogram"] > 0 {
			score += weights["MACD"]
		}
	} else {
		score -= weights["MACD"]
		if macd["histogram"] < 0 {
			score -= weights["MACD"]
		}
	}

	if data.Volume == "Volume up" || data.Volume == "Whale activity" {
		score += weights["Volume"]
	} else if data.Volume == "Volume down" {
		score -= weights["Volume"]
	}

	ema50 := data.Ema50
	ema200 := data.Ema200
	if price > ema50 {
		score += weights["EMA"]
	} else {
		score -= weights["EMA"]
	}
	if ema50 > ema200 {
		score += weights["EMA"]
	} else {
		score -= weights["EMA"]
	}

	patterns := data.Patterns
	bullish_patterns := []string{"hammer", "dragonfly_doji", "builish_engulfing", "white_solders"}
	bearish_patterns := []string{"shooting_star", "gravestone_doji", "bearish_engulfing", "black_solders"}

	for _, p := range bullish_patterns {
		if patterns[p] {
			score += weights["PatternBullish"]
		}
	}
	for _, p := range bearish_patterns {
		if patterns[p] {
			score -= weights["PatternBearish"]
		}
	}
	if patterns["doji"] {
		score += 0
	}

	return int(score)
}
