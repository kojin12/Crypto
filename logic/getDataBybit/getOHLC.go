package getdatabybit

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type OHLCResponse struct {
	RetCode int        `json:"retCode"`
	RetMsg  string     `json:"retMsg"`
	Result  OHLCResult `json:"result"`
}

type OHLCResult struct {
	Symbol   string     `json:"symbol"`
	Category string     `json:"category"`
	List     [][]string `json:"list"`
}

func GetOHLC(timeFrame string, pair string, limit string) ([][]string, float64) {
	baseURL := "https://api.bybit.com/v5/market/kline"

	urlResult, err := url.Parse(baseURL)

	if err != nil {
		panic(err)
	}

	params := url.Values{}
	params.Add("category", "linear")
	params.Add("symbol", pair)
	params.Add("interval", timeFrame)
	params.Add("limit", limit)

	urlResult.RawQuery = params.Encode()

	resp, err := http.Get(urlResult.String())

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var BybitData OHLCResponse

	err = json.Unmarshal(body, &BybitData)

	if err != nil {
		panic(err)
	}

	reversedData := BybitData.Result.List
	var validateData [][]string

	for i := len(reversedData) - 1; i >= 0; i-- {
		validateData = append(validateData, reversedData[i])
	}

	price, err := strconv.ParseFloat(validateData[len(validateData)-1][4], 64)
	if err != nil {
		panic(err)
	}

	return validateData, price
}
