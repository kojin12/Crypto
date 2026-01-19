package getanalispair

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type BybitDataTicket struct {
	RetCcode string       `json:"ret_code"`
	Msg      string       `json:"ret_msg"`
	Result   ResultTicket `json:"result"`
}

type ResultTicket struct {
	Symbol       string `json:"symbol"`
	LastPrice    string `json:"last_price"`
	PrevPrice24h string `json:"prev_price_24h"`
	Price24hPcnt string `json:"price_24h_pcnt"`
}

func GetPriceChange(pair string) ResultTicket {
	baseURL := "https://api.bybit.com/v2/public/tickers"

	parseURL, err := url.Parse(baseURL)

	if err != nil {
		panic(err)
	}

	params := url.Values{}

	params.Add("symbol", pair)

	parseURL.RawQuery = params.Encode()

	resp, err := http.Get(parseURL.String())

	if err != nil {
		panic(err)
	}

	var data BybitDataTicket

	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	err = json.Unmarshal(body, &data)

	if err != nil {
		panic(err)
	}

	return data.Result
}
