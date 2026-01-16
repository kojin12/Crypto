package alert

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type LickData struct {
	Symbol       string `json:"symbol"`
	LastPrice    string `json:"lastPrice"`
	Volume24h    string `json:"volume24h"`
	Turnover24h  string `json:"turnover24h"`
	Price24hPcnt string `json:"price24hPcnt"`
}

func getDataLick(pair string) float64 {
	baseURL := "https://api.bybit.com/v5/market/tickers"

	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}

	params := url.Values{}

	params.Add("category", "linear")
	params.Add("symbol", pair)

	u.RawQuery = params.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var Resp LickData

	if err := json.Unmarshal(body, &Resp); err != nil {
		panic(err)
	}

	lick, _ := strconv.ParseFloat(Resp.Turnover24h, 64)
	return lick
}
