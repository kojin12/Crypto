package alert

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type SpotRes struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type LinearRes struct {
	Success bool    `json:"success"`
	Data    DataRes `json:"data"`
}

type DataRes struct {
	Symbol    string  `json:"symbol"`
	FairPrice float64 `json:"fairPrice"`
}

type ChangeRes struct {
	Diff   float64
	Profit float64
	Result bool
	Price  PriceSpotAndLinear
}

type PriceSpotAndLinear struct {
	SpotPrice   float64
	LinearPrice float64
}

func GetSpotPrice(pair string) float64 {
	baseURL := "https://api.mexc.com/api/v3/ticker/price"

	urlResult, err := url.Parse(baseURL)

	if err != nil {
		panic(err)
	}

	params := url.Values{}
	params.Add("symbol", pair)

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

	var SpotResult SpotRes

	err = json.Unmarshal(body, &SpotResult)

	floatPrice, err := strconv.ParseFloat(SpotResult.Price, 64)
	if err != nil {
		panic(err)
	}

	return floatPrice
}

func GetLinearPrice(pair string) float64 {

	baseURL := "https://contract.mexc.com/api/v1/contract/fair_price/" + pair

	resp, err := http.Get(baseURL)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var LinearData LinearRes

	err = json.Unmarshal(body, &LinearData)

	return LinearData.Data.FairPrice
}

func CalcChangeAndProfit(SpotPrice float64, LinearPrice float64, deposit float64) ChangeRes {
	spotFee := 0.004
	futFee := 0.002

	priceDiff := LinearPrice - SpotPrice
	feeCost := SpotPrice*spotFee + LinearPrice*futFee

	var res ChangeRes

	if priceDiff > feeCost {
		oneProfit := priceDiff - feeCost

		takeProfit := deposit / SpotPrice

		finalProfit := takeProfit * oneProfit

		res.Result = true
		res.Diff = priceDiff
		res.Profit = finalProfit
		res.Price.LinearPrice = LinearPrice
		res.Price.SpotPrice = SpotPrice

		return res
	}

	res.Result = false
	res.Diff = priceDiff
	res.Profit = 0.0
	res.Price.LinearPrice = LinearPrice
	res.Price.SpotPrice = SpotPrice

	return res
}
