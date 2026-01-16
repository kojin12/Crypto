package logic

func GetScore(data FinalResult, price float64) int {
	score := 0

	if data.Trend15 == "Up" {
		score += 1
	} else {
		score -= 1
	}

	if data.Trend60 == "Up" {
		score += 2
	} else {
		score -= 2
	}

	if data.Trend15 == data.Trend60 {
		score += 1
	} else {
		score -= 1
	}

	rsi := data.Rsi

	if rsi >= 45 && rsi <= 55 {
		score += 1
	} else if rsi > 55 && rsi <= 65 {
		score += 1
	} else if rsi >= 35 && rsi < 45 {
		score -= 1
	} else if rsi > 70 {
		score -= 2
	} else if rsi < 30 {
		score -= 2
	}

	macd := data.MACD

	if macd["macd_line"] > macd["signal_line"] {
		score += 1
		if macd["histogram"] > 0 {
			score += 1
		}
	} else {
		score -= 1
		if macd["histogram"] < 0 {
			score -= 1
		}
	}

	if data.Volume == "Volume up" || data.Volume == "Whale activity" {
		score += 1
	} else if data.Volume == "Volume down" {
		score -= 1
	}

	ema50 := data.Ema50
	ema200 := data.Ema200

	if price > ema50 {
		score += 1
	} else {
		score -= 1
	}

	if ema50 > ema200 {
		score += 1
	} else {
		score -= 1
	}

	patterns := data.Patterns

	bullish_patterns := []string{
		"hammer", "dragonfly_doji", "builish_engulfing", "white_solders",
	}

	bearish_patterns := []string{
		"shooting_star", "gravestone_doji", "bearish_engulfing", "black_solders",
	}

	for _, p := range bullish_patterns {
		if patterns[p] {
			score += 1
		}
	}

	for _, p := range bearish_patterns {
		if patterns[p] {
			score -= 1
		}
	}

	if patterns["doji"] {
		score += 0
	}

	return score
}
