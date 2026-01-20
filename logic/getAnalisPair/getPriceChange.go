package getanalispair

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type BybitDataTicket struct {
	RetCode int             `json:"retCode"`
	Msg     string          `json:"retMsg"`
	Result  ResultContainer `json:"result"`
}

type ResultContainer struct {
	Category string         `json:"category"`
	List     []ResultTicket `json:"list"`
}

type ResultTicket struct {
	Symbol       string `json:"symbol"`
	LastPrice    string `json:"lastPrice"`
	PrevPrice24h string `json:"prevPrice24h"`
	Price24hPcnt string `json:"price24hPcnt"`
}

func GetPriceChange(pair string) ResultTicket {
	baseURL := "https://api.bybit.com/v5/market/tickers"

	parseURL, err := url.Parse(baseURL)

	if err != nil {
		panic(err)
	}

	params := url.Values{}

	params.Add("category", "spot")
	params.Add("symbol", pair)

	parseURL.RawQuery = params.Encode()

	resp, err := http.Get(parseURL.String())

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var data BybitDataTicket

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("API Response:", string(body))

	err = json.Unmarshal(body, &data)

	if err != nil {
		panic(err)
	}

	if len(data.Result.List) == 0 {
		panic("no results from API")
	}

	return data.Result.List[0]
}
