package alert

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type ByBitPairResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Category string       `json:"category"`
		List     []Instrument `json:"list"`
	} `json:"result"`
}

type Instrument struct {
	Symbol    string `json:"symbol"`
	Status    string `json:"status"`
	BaseCoin  string `json:"baseCoin"`
	QuoteCoin string `json:"quoteCoin"`
}

func getPairsByStatus(status string) []Instrument {
	baseURL := "https://api.bybit.com/v5/market/instruments-info"
	u, _ := url.Parse(baseURL)
	params := url.Values{}
	params.Add("category", "linear")
	params.Add("status", status)
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

	var response ByBitPairResponse
	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}

	if response.RetCode != 0 {
		panic(response.RetMsg)
	}

	return response.Result.List
}

func GetCoinList() []Instrument {
	list := []Instrument{}

	pre := getPairsByStatus("PreLaunch")
	list = append(list, pre...)

	tr := getPairsByStatus("Trading")
	list = append(list, tr...)

	return list
}
