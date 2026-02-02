package config

type PairConfig struct {
	ByBit     string
	Mexc      string
	CoinGecko string
}

var CoinConfig = map[string]PairConfig{
	"btc": {
		ByBit:     "BTCUSDT",
		Mexc:      "BTCUSDT",
		CoinGecko: "bitcoin",
	},
	"eth": {
		ByBit:     "ETHUSDT",
		Mexc:      "ETHUSDT",
		CoinGecko: "ethereum",
	},
	"ton": {
		ByBit:     "TONUSDT",
		Mexc:      "TONUSDT",
		CoinGecko: "the-open-network",
	},
	"sol": {
		ByBit:     "SOLUSDT",
		Mexc:      "SOLUSDT",
		CoinGecko: "solana",
	},
}
