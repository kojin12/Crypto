package handlers

import (
	"encoding/json"
	"log"
	"main/logic"
	getdatamexc "main/logic/getDataMexc"
	"net/http"
)

type DataResponse struct {
	Data  interface{} `json:"data"`
	Score int         `json:"score"`
}

func GetDataHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	q := r.URL.Query()
	pair := q.Get("pair")
	timeFrame := q.Get("timeframe")
	market := q.Get("market")
	log.Printf("Received parameters: pair=%s, timeframe=%s, market=%s", pair, timeFrame, market)
	limit := "200"

	if pair == "" {
		http.Error(w, "pair is required", http.StatusBadRequest)
		return
	}
	if timeFrame == "" {
		http.Error(w, "timeframe is required", http.StatusBadRequest)
		return
	}
	if market == "" {
		http.Error(w, "market is required", http.StatusBadRequest)
		return
	}

	var response logic.FinalResult
	var score int

	if market == "mexc" {
		ohlc, price := getdatamexc.GetMexcOHLC(pair, timeFrame, limit)
		ohlc60, _ := getdatamexc.GetMexcOHLC(pair, "60m", limit)

		response = logic.Final(ohlc, ohlc60)
		score = logic.GetScore(response, price)
	}

	result := DataResponse{
		Data:  response,
		Score: score,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
