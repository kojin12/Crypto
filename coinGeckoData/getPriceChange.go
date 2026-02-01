package coingeckodata

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DataPriceChange struct {
	Price       float64 `json:"usd"`
	PriceChange float64 `json:"usd_24h_change"`
}

func GetPriceChange(pair string) (float64, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd&include_24hr_change=true", pair)

	resp, err := http.Get(url)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	var result map[string]DataPriceChange
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0.0, err
	}

	data, ok := result[pair]
	if !ok {
		return 0.0, fmt.Errorf("no data result for, %s" + pair)
	}
	return data.PriceChange, nil
}
