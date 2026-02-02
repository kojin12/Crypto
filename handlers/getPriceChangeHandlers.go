package handlers

import (
	"encoding/json"
	coingeckodata "main/coinGeckoData"
	"main/config"
	"net/http"
)

type ResultData struct {
	PriceChange float64
}

func GetPriceChangeHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	q := r.URL.Query()

	pair := q.Get("pair")
	if pair == "" {
		http.Error(w, "pair is required", http.StatusBadRequest)
		return
	}

	pairSymbol, ok := config.CoinConfig[pair]
	if !ok {
		http.Error(w, "unknown pair", http.StatusBadRequest)
	}

	var ResultDataVar ResultData
	result, err := coingeckodata.GetPriceChange(pairSymbol.CoinGecko)
	if err != nil {
	}

	ResultDataVar.PriceChange = result
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResultDataVar)
}
